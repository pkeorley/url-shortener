// MIT License
// Copyright (c) 2024 pkeorley
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice (including the next paragraph) shall be included in all copies
// or substantial portions of the Software.

package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"slices"
)

var (
	ErrDialectorNotFound  = fmt.Errorf("url-shortener/internal/database: dialector not found")
	ErrKeyNotFound        = errors.New("url-shortener/internal/database: key not found")
	ErrApiKeyNotExists    = errors.New("url-shortener/internal/database: api key doesn't exist")
	ErrShortLinkExists    = errors.New("url-shortener/internal/database: short link already exists")
	ErrShortLinkNotExists = errors.New("url-shortener/internal/database: short link doesn't exist")
)

// Database represents a database instance.
type Database struct {
	DB *gorm.DB
}

// New initializes a new Database instance based on the provided Dialector.
//
// It takes a Dialector as a parameter and returns a Database.
func New(d Dialector) Database {
	db, err := gorm.Open(getDialector(d), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return Database{DB: db}
}

// GetShortLinkByShortname retrieves a ShortLink from the Database based on the given shortname.
func (db Database) GetShortLinkByShortname(shortname string) *ShortLink {
	by, _ := db.GetShortLinkBy("shortname", shortname)
	return by
}

// GetShortLinkByURL retrieves a short link by the given URL.
func (db Database) GetShortLinkByURL(url string) *ShortLink {
	by, _ := db.GetShortLinkBy("url", url)
	return by
}

// GetShortLinkByApiKey retrieves a ShortLink by the provided API key.
func (db Database) GetShortLinkByApiKey(apiKey uuid.UUID) *ShortLink {
	by, _ := db.GetShortLinkBy("api_key", apiKey.String())
	return by
}

// GetShortLinkBy retrieves a ShortLink from the Database based on the specified key and value.
func (db Database) GetShortLinkBy(key string, value interface{}) (*ShortLink, error) {
	var shortLink = ShortLink{}

	if !slices.Contains([]string{"shortname", "url", "api_key"}, key) {
		return nil, ErrKeyNotFound
	}

	result := db.DB.Model(&ShortLink{}).Where(fmt.Sprintf("%s = ?", key), value).First(&shortLink)
	if result.RowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &shortLink, nil
}

// GetApiKeyByApiKey returns an API key based on the provided API key.
func (db Database) GetApiKeyByApiKey(apiKey uuid.UUID) *ApiKey {
	by, _ := db.GetApiKeyBy("api_key", apiKey.String())
	return by
}

// GetApiKeyBy retrieves an API key by a specified key and value.
func (db Database) GetApiKeyBy(key string, value interface{}) (*ApiKey, error) {
	var apiKey = ApiKey{}

	// TODO: Qué más?!?
	if !slices.Contains([]string{"api_key"}, key) {
		return nil, ErrKeyNotFound
	}

	// TODO: Preload?!?
	result := db.DB.Model(&ApiKey{}).Preload("ShortLinks").Where(fmt.Sprintf("%s = ?", key), value).First(&apiKey)

	if result.RowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &apiKey, nil
}

// CreateShortLink creates a short link in the database.
func (db Database) CreateShortLink(apiKey uuid.UUID, shortname string, url string) (*ShortLink, error) {
	// Does the API key exist?
	if db.GetApiKeyByApiKey(apiKey) == nil {
		return nil, ErrApiKeyNotExists
	}

	// Validation of the shortname
	isValidShortname := func() bool {
		return len(shortname) < 32 && len(shortname) > 0
	}
	if !isValidShortname() {
		return nil, errors.New("url-shortener/internal/database: the shortname is need to be between 1 and 32 characters")
	}

	/// Validation of the URL
	isValidURL := func() bool {
		// TODO: I should be use regex to validate url
		return true
	}
	if !isValidURL() {
		return nil, errors.New("url-shortener/internal/database: the url is not valid")
	}

	// If the shortname already exists, return an error
	if db.GetShortLinkByShortname(shortname) != nil {
		return nil, ErrShortLinkExists
	}

	var shortLink = ShortLink{
		ApiKey:    apiKey,
		Shortname: shortname,
		URL:       url,
	}

	// Create short link and return it to the shortLink variable
	db.DB.Model(&ShortLink{}).
		Create(&shortLink).
		First(&shortLink)

	// Update credits
	db.DB.Model(&ApiKey{}).Where("api_key = ?", apiKey).
		Update("utilized_credits", gorm.Expr("utilized_credits + 1")).
		Update("available_credits", gorm.Expr("available_credits - 1"))

	return &shortLink, nil
}

// CreateApiKey creates a new API key in the database.
func (db Database) CreateApiKey(key ...uuid.UUID) (*ApiKey, error) {
	if len(key) == 0 {
		key = append(key, uuid.New())
	}

	var apiKey = ApiKey{ApiKey: key[0]}

	// Create API key and return it to the apiKey variable
	result := db.DB.Model(&ApiKey{}).Create(&apiKey).First(&apiKey)
	if result.Error != nil {
		return nil, result.Error
	}

	return &apiKey, nil
}
