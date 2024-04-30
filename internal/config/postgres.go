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
	"github.com/caarlos0/env/v11"
	"log"
)

// Postgres represents the Postgres configuration.
type Postgres struct {
	Username string `env:"POSTGRES_USERNAME,required"`
	Host     string `env:"POSTGRES_HOST,required"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	Password string `env:"POSTGRES_PASSWORD,required"`
	DBName   string `env:"DB_NAME,required"`
}

// String returns a string representation of the postgres configuration.
func (p Postgres) String() any {
	return fmt.Sprintf("postgres{Username:%v, Host:%v, Port:%v, Password:%v, DBName:%v}", p.Username, p.Host, p.Port, p.Password, p.DBName)
}

// GetDSN returns the DSN for the postgres database.
func (p Postgres) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", p.Host, p.Port, p.Username, p.DBName, p.Password)
}

// newPostgres returns a new Postgres configuration.
func newPostgres() *Postgres {
	pg, err := env.ParseAs[Postgres]()
	if err != nil {
		log.Fatal(err)
	}
	return &pg
}
