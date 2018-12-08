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
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"

	"github.com/graphql-go/graphql"
	"github.com/yaacov/tsl/pkg/tsl"
	walker "github.com/yaacov/tsl/pkg/walkers/sql"

	"github.com/yaacov/tsl/cmd/model"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	preparePtr := flag.Bool("p", false, "prepare a book collection for queries")
	filePtr := flag.String("f", "./sqlite.db", "the sqlite database file name")
	flag.Parse()

	// Set context.
	ctx := context.Background()

	// Create a clean books collection.
	if *preparePtr {
		tx, err := connect(ctx, *filePtr)
		check(err)

		err = prepareCollection(ctx, tx)
		check(err)

		tx.Commit()
	}

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"books": &graphql.Field{
				Type:        graphql.NewList(bookType),
				Description: "list books",
				Args: graphql.FieldConfigArgument{
					"filter": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "use a TSL phrase to filter results",
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var bookID uint
					var rows *sql.Rows
					var books []model.Book

					filter := params.Args["filter"].(string)

					// Try to connect to mongo server.
					tx, err := connect(ctx, *filePtr)
					if err != nil {
						return nil, err
					}
					defer tx.Commit()

					// Parse input string into a TSL tree.
					tree, err := tsl.ParseTSL(filter)
					if err != nil {
						return nil, err
					}

					// Prepare SQL filter.
					sqFilter, err := walker.Walk(tree)
					if err != nil {
						return nil, err
					}

					// Query SQL table.
					rows, err = sq.
						Select("*").
						Where(sqFilter).
						From("books").
						RunWith(tx).
						QueryContext(ctx)
					if err != nil {
						return nil, err
					}

					for rows.Next() {
						elem := model.Book{}
						err = rows.Scan(
							&bookID,
							&elem.Title,
							&elem.Author,
							&elem.Spec.Pages,
							&elem.Spec.Rating,
						)
						if err != nil {
							return nil, err
						}

						books = append(books, elem)
					}

					// Check for errors and exit.
					err = rows.Err()
					if err != nil {
						return nil, err
					}

					return books, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	qraphqlHandler := func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})

		// Write output to response.
		json.NewEncoder(w).Encode(result)
	}

	http.HandleFunc("/graphql", qraphqlHandler)
	http.ListenAndServe(":8080", nil)
}
