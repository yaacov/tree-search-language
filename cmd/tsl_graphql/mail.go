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
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/graphql-go/graphql"

	"github.com/yaacov/tsl/pkg/tsl"
	"github.com/yaacov/tsl/pkg/walkers/ident"
	walker "github.com/yaacov/tsl/pkg/walkers/sql"

	"github.com/yaacov/tsl/cmd/model"
)

var preparePtr *bool
var filePtr *string

// columnNamesMap mapps between user namespace and the SQL column names.
var columnNamesMap = map[string]string{
	"title":       "title",
	"author":      "author",
	"spec.pages":  "pages",
	"spec.rating": "rating",
}

var bookSpecType = graphql.NewObject(graphql.ObjectConfig{
	Name: "BookSpecs",
	Fields: graphql.Fields{
		"pages": &graphql.Field{
			Type: graphql.Int,
		},
		"rating": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: graphql.String,
		},
		"spec": &graphql.Field{
			Type: bookSpecType,
		},
	},
})

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

func sqlQuery(ctx context.Context, filter string) (books []model.Book, err error) {
	var bookID uint
	var rows *sql.Rows

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

	// Check and replace user identifiers with the SQL table column names.
	tree, err = ident.Walk(tree, checkColumnName)
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

	return
}

func main() {
	// Setup the input.
	preparePtr = flag.Bool("p", false, "prepare a book collection for queries")
	filePtr = flag.String("f", "./sqlite.db", "the sqlite database file name")
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
					filter := params.Args["filter"].(string)
					return sqlQuery(ctx, filter)
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
