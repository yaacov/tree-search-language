// Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main
package main

import (
	"flag"
	"fmt"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/yaacov/tsl/pkg/tsl"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
}

func main() {
	var s sq.SelectBuilder

	// Setup the input
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"animal = 'kitty'\")")
	tablePtr := flag.String("t", "<table-name>", "the table name to use for SQL generation")
	outputPtr := flag.String("o", "mysql", "output format [mysql/pgsql]")
	flag.Parse()

	// Parse input string into a TSL tree
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)

	// If parser has erros, we can not print the tree
	if err != nil {
		os.Exit(1)
	}

	switch *outputPtr {
	case "pgsql":
		// If we are using PostgreSQL style use $ instead of ?
		// for placeholders
		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		s = psql.Select("*")
	case "mysql":
		s = sq.Select("*")
	default:
		s = sq.Select("*")
	}

	filter, err := tsl.SquirrelWalk(tree)
	check(err)

	// If SquirrelWalk has erros, we can not print the tree
	if err != nil {
		os.Exit(1)
	}

	sql, args, err := s.
		Where(filter).
		From(*tablePtr).
		ToSql()

	check(err)
	fmt.Printf("sql:  %s\n", sql)
	fmt.Printf("args: %v\n", args)
}
