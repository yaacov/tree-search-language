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
	"flag"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/yaacov/tsl/pkg/tsl"
)

func main() {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var filter *bson.Element
	var s string
	var b []byte

	// Setup the input.
	inputPtr := flag.String("i", "title is not null", "the tsl string to parse (e.g. \"author = 'Jane'\")")
	preparePtr := flag.Bool("p", false, "prepare a book collection for queries")
	dbNamePtr := flag.String("d", "tsl", "db name to connect to")
	collectionNamePtr := flag.String("c", "books", "collection name to query on")
	urlPtr := flag.String("u", "mongodb://localhost:27017", "url for mongo server")
	flag.Parse()

	// Set context.
	ctx := context.Background()

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Try to connect to mongo server.
	client, err = connect(ctx, *urlPtr)
	if err != nil {
		log.Fatal(err)
	}

	// Set up a collection.
	collection = client.Database(*dbNamePtr).Collection(*collectionNamePtr)

	// Create a clean books collection.
	if *preparePtr {
		err = prepareCollection(ctx, collection)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Prepare a bson filter.
	filter, err = tsl.BSONWalk(tree)
	if err != nil {
		log.Fatal(err)
	}

	// Run query.
	cur, err := collection.Find(ctx, bson.NewDocument(filter))
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	// Loop on query elements.
	for cur.Next(ctx) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}

		// Print out books as json.
		b, _ = elem.MarshalBSON()
		s, _ = bson.ToExtJSON(true, b)
		fmt.Printf("%v\n", s)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
