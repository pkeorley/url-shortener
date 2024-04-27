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

type Fiber struct {
	Port string `env:"FIBER_PORT" envDefault:"3000"`
}

func (f Fiber) String() string {
	return fmt.Sprintf("Fiber{Port:%v}", f.Port)
}

func NewFiber() *Fiber {
	fiber, err := env.ParseAs[Fiber]()
	if err != nil {
		log.Fatal(err)
	}
	return &fiber
}
