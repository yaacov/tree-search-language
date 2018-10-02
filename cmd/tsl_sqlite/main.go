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

// Package main.
package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/yaacov/tsl/pkg/tsl"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var rows *sql.Rows

	// Setup the input.
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"Title = 'Book'\")")
	preparePtr := flag.Bool("p", false, "prepare a book collection for queries")
	filePtr := flag.String("f", "./sqlite.db", "the sqlite database file name")
	flag.Parse()

	// Sanity check.
	if *inputPtr == "" {
		err := fmt.Errorf("missing required flag -i (the tsl string to parse)")
		check(err)
	}

	// Set context.
	ctx := context.Background()

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)

	// Try to connect to mongo server.
	tx, err := connect(ctx, *filePtr)
	check(err)

	defer tx.Commit()

	// Create a clean books collection.
	if *preparePtr {
		err = prepareCollection(ctx, tx)
		check(err)
	}

	// Prepare SQL filter.
	filter, err := tsl.SquirrelWalk(tree)
	check(err)

	// Query SQL table.
	rows, err = sq.
		Select("*").
		Where(filter).
		From("books").
		RunWith(tx).
		QueryContext(ctx)
	check(err)

	for rows.Next() {
		b := book{}
		err = rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&b.Pages,
			&b.Rating,
		)
		check(err)

		fmt.Printf("%v\n", b)
	}

	// Check for errors and exit.
	err = rows.Err()
	check(err)
}
