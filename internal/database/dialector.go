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
	"github.com/pkeorley/url-shortener/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Dialector represents a database dialect.
type Dialector string

const (
	// DialectorPostgres represents the Postgres dialect.
	DialectorPostgres Dialector = "postgres"
)

// getDialector returns a gorm.Dialector based on the provided Dialector.
func getDialector(d Dialector) gorm.Dialector {
	switch d {
	default:
		log.Fatal(ErrDialectorNotFound)
		return nil
	case DialectorPostgres:
		var pgCfg = config.New().GetPostgres()
		return postgres.Open(pgCfg.GetDSN())
	}
}
