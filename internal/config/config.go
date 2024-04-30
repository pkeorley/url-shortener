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

package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

// LoadDotenvValues loads environment variables from a .env file using the godotenv package.
func LoadDotenvValues(filenames ...string) {
	if err := godotenv.Load(filenames...); err != nil {
		log.Fatal(err)
	}
}

// Config represents the application configuration.
type Config struct{}

// New returns a new Config.
func New() *Config {
	LoadDotenvValues()
	return &Config{}
}

// String returns a string representation of the Config.
func (Config) String() string {
	return fmt.Sprintf("Config{}")
}

// GetPostgres returns a new instance of Postgres configured based on the provided Config.
func (Config) GetPostgres() *postgres {
	return newPostgres()
}

// GetFiber returns a new instance of Fiber configured based on the provided Config.
func (Config) GetFiber() *Fiber {
	return newFiber()
}
