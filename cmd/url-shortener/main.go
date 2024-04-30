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

package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkeorley/url-shortener/internal/config"
	"github.com/pkeorley/url-shortener/internal/database"
	"log"
)

var (
	db database.Database
)

func main() {
	cfg := config.New()

	db = database.New(database.DialectorPostgres)
	if err := db.DB.AutoMigrate(
		&database.ApiKey{},
		&database.ShortLink{},
	); err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               "url-shortener",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"app_name": "url-shortener",
			"api_url":  fmt.Sprintf("%s/api/v1", c.BaseURL()),
		})
	})

	app.Get("/i/:shortname", func(c *fiber.Ctx) error {
		shortname := c.Params("shortname", "")

		shortLink := db.GetShortLinkByShortname(shortname)
		if shortLink == nil {
			return database.ErrShortLinkNotExists
		}

		return c.Redirect(shortLink.URL)
	})

	app.Post("/api/v1/shortlinks/create", func(c *fiber.Ctx) error {
		// This struct is meant to parse the request body
		type Body struct {
			ApiKey    string `json:"api_key"`
			Shortname string `json:"shortname"`
			URL       string `json:"url"`
		}

		// This function is meant to check if the request body contains all required fields
		isValidBody := func(b Body) bool {
			return len(b.Shortname) > 0 && len(b.URL) > 0 && len(b.ApiKey) > 0
		}

		// This function is meant to find the expected field in the request body
		findExceptedValue := func(b Body) string {
			if len(b.Shortname) == 0 {
				return "shortname"
			} else if len(b.URL) == 0 {
				return "url"
			} else if len(b.ApiKey) == 0 {
				return "api_key"
			}
			return "" // All fields are non-empty
		}

		var body Body
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		if !isValidBody(body) {
			return fmt.Errorf("%q is required field in the request body", findExceptedValue(body))
		}

		// The short link creation logic

		apiKey, err := uuid.Parse(body.ApiKey)
		if err != nil {
			return err
		}

		shortLink, err := db.CreateShortLink(apiKey, body.Shortname, body.URL)
		if err != nil {
			return err
		}

		return c.JSON(shortLink)
	})

	log.Fatal(app.Listen(":" + cfg.GetPostgres().Port))
}
