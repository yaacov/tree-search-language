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
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/yaacov/tsl/pkg/tsl"
)

const dbName = "tsl"
const collectionName = "books"

var books = []interface{}{
	book{Title: "Book", Author: "Joe", Spec: bookSpecs{Pages: 100, Raiting: 4}},
	book{Title: "Other Book", Author: "Jane", Spec: bookSpecs{Pages: 200, Raiting: 3}},
	book{Title: "Some Other Book", Author: "Jane", Spec: bookSpecs{Pages: 50, Raiting: 5}},
	book{Title: "Some Other Book", Author: "Jane", Spec: bookSpecs{Pages: 50}},
}

func connect(ctx context.Context, url string) (client *mongo.Client, err error) {
	client, err = mongo.NewClient(url)
	if err != nil {
		return
	}
	err = client.Connect(ctx)

	return
}

func prepareCollection(ctx context.Context, c *mongo.Client) (collection *mongo.Collection, err error) {
	// Drop the collection
	collection = c.Database(dbName).Collection(collectionName)
	err = collection.Drop(ctx)
	if err != nil {
		return
	}

	// Insert new books into the collection
	collection = c.Database(dbName).Collection(collectionName)
	_, err = collection.InsertMany(ctx, books)

	return
}

func main() {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var filter interface{}
	var s string
	var b []byte

	// Setup the input
	inputPtr := flag.String("i", "title is not null", "the tsl string to parse (e.g. \"author = 'Jane'\")")
	urlPtr := flag.String("u", "mongodb://localhost:27017", "url for mongo server")
	flag.Parse()

	// Set context
	ctx := context.Background()

	// Parse input string into a TSL tree
	tree, err := tsl.ParseTSL(*inputPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Try to connect to mongo server
	client, err = connect(ctx, *urlPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Create a clean books collection
	collection, err = prepareCollection(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare filter
	filter = bson.NewDocument(bsonWalk(tree))

	// Run query
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	// Loop on query elements
	for cur.Next(ctx) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}

		// Print out books as json
		b, _ = elem.MarshalBSON()
		s, _ = bson.ToExtJSON(true, b)
		fmt.Printf("%v\n", s)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
