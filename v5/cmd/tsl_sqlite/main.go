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
	"encoding/json"
	"flag"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/yaacov/tree-search-language/v5/pkg/tsl"
	"github.com/yaacov/tree-search-language/v5/pkg/walkers/ident"
	walker "github.com/yaacov/tree-search-language/v5/pkg/walkers/sql"

	"github.com/yaacov/tree-search-language/v5/cmd/model"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// columnNamesMap mapps between user namespace and the SQL column names.
var columnNamesMap = map[string]string{
	"title":       "title",
	"author":      "author",
	"spec.pages":  "pages",
	"spec.rating": "rating",
	"on_loan":     "onloan",
}

// checkColumnName checks if a coulumn name is valid in user space replace it
// with the mapped column name and returns and error if not a valid name.
func checkColumnName(s string) (string, error) {
	// Chekc for column name in map.
	if v, ok := columnNamesMap[s]; ok {
		return v, nil
	}

	// If not found return string as is, and an error.
	return s, fmt.Errorf("column \"%s\" not found", s)
}

func main() {
	var bookID uint
	var rows *sql.Rows

	// Setup the input.
	inputPtr := flag.String("i", "title is not null", "the tsl string to parse (e.g. \"title = 'Book'\")")
	preparePtr := flag.Bool("p", false, "prepare a book collection for queries")
	filePtr := flag.String("f", "./sqlite.db", "the sqlite database file name")
	flag.Parse()

	// Set context.
	ctx := context.Background()

	// Try to connect to sqlite driver.
	tx, err := connect(ctx, *filePtr)
	check(err)
	defer tx.Commit()

	// Create a clean books collection.
	if *preparePtr {
		err = prepareCollection(ctx, tx)
		check(err)
	}

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)

	// Check and replace user identifiers with the SQL table column names.
	tree, err = ident.Walk(tree, checkColumnName)
	check(err)

	// Prepare SQL filter.
	filter, err := walker.Walk(tree)
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
		elem := model.Book{}
		err = rows.Scan(
			&bookID,
			&elem.Title,
			&elem.Author,
			&elem.Spec.Pages,
			&elem.Spec.Rating,
			&elem.OnLoan,
		)
		check(err)

		b, _ := json.Marshal(elem)
		fmt.Printf("%s\n", string(b))
	}

	// Check for errors and exit.
	err = rows.Err()
	check(err)
}
