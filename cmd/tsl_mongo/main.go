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
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/yaacov/tsl/pkg/tsl"
	"github.com/yaacov/tsl/pkg/walkers/ident"
	walker "github.com/yaacov/tsl/pkg/walkers/mongo"

	"github.com/yaacov/tsl/cmd/model"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// columnNamesMap mapps between user namespace and the MongoDB document fields.
var columnNamesMap = map[string]string{
	"title":       "title",
	"author":      "author",
	"spec.pages":  "spec.pages",
	"spec.rating": "spec.rating",
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
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var filter bson.D

	// Setup the input.
	inputPtr := flag.String("i", "title is not null", "the tsl string to parse (e.g. \"author = 'Jane'\")")
	preparePtr := flag.Bool("p", false, "prepare a book collection for queries")
	dbNamePtr := flag.String("d", "tsl", "db name to connect to")
	collectionNamePtr := flag.String("c", "books", "collection name to query on")
	urlPtr := flag.String("u", "mongodb://localhost:27017", "url for mongo server")
	flag.Parse()

	// Set context.
	ctx := context.Background()

	// Try to connect to mongo server.
	client, err = connect(ctx, *urlPtr)
	check(err)

	// Set up a collection.
	collection = client.Database(*dbNamePtr).Collection(*collectionNamePtr)

	// Create a clean books collection.
	if *preparePtr {
		err = prepareCollection(ctx, collection)
		check(err)
	}

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)

	// Check and replace user identifiers with MongoDB document field names.
	tree, err = ident.Walk(tree, checkColumnName)
	check(err)

	// Prepare a bson filter.
	filter, err = walker.Walk(tree)
	check(err)

	// Run query.
	cur, err := collection.Find(ctx, filter)
	check(err)
	defer cur.Close(ctx)

	// Loop on query elements.
	for cur.Next(ctx) {
		elem := model.Book{}
		err := cur.Decode(&elem)
		check(err)

		b, _ := json.Marshal(elem)
		fmt.Printf("%s\n", string(b))
	}

	// Check for errors and exit.
	err = cur.Err()
	check(err)
}
