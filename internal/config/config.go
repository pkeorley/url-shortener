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

import "fmt"

// Config represents the application configuration.
type Config struct {
	Postgres *Postgres
	Fiber    *Fiber
}

// String returns a string representation of the Config.
func (c Config) String() string {
	return fmt.Sprintf("Config{Postgres: %v}", c.Postgres.String())
}

// New returns a new Config.
func New() *Config {
	return &Config{
		Postgres: NewPostgres(),
		Fiber:    NewFiber(),
	}
}
