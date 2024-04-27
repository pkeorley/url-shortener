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
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Model represents a base model.
type Model struct {
	CreatedAt time.Time      `gorm:"autoCreateTime:nano" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:nano" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// ApiKey represents an api_keys table in the database.
type ApiKey struct {
	Model
	ApiKey           uuid.UUID   `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"api_key"`
	AvailableCredits int         `gorm:"default:5" json:"available_credits"`
	UtilizedCredits  int         `gorm:"default:0" json:"utilized_credits"`
	ShortLinks       []ShortLink `gorm:"foreignKey:ApiKey" json:"short_links"`
}

// ShortLink represents a short_links table in the database.
type ShortLink struct {
	Model
	ApiKey    uuid.UUID `gorm:"type:uuid,not null" json:"api_key"`
	Shortname string    `gorm:"unique,not null" json:"shortname"`
	URL       string    `gorm:"not null" json:"url"`
}
