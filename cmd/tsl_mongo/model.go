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

	"github.com/mongodb/mongo-go-driver/mongo"
)

type bookSpecs struct {
	Pages  uint `bson:"pages,omitempty" json:"pages,omitempty"`
	Rating uint `bson:"rating,omitempty" json:"rating,omitempty"`
}

type book struct {
	Title  string    `bson:"title,omitempty" json:"title,omitempty"`
	Author string    `bson:"author,omitempty" json:"author,omitempty"`
	Spec   bookSpecs `bson:"spec,omitempty" json:"spec,omitempty"`
}

var books = []interface{}{
	book{Title: "Book", Author: "Joe", Spec: bookSpecs{Pages: 100, Rating: 4}},
	book{Title: "Other Book", Author: "Jane", Spec: bookSpecs{Pages: 200, Rating: 3}},
	book{Title: "Some Book", Author: "Jane", Spec: bookSpecs{Pages: 50, Rating: 5}},
	book{Title: "Some Other Book", Author: "Jane", Spec: bookSpecs{Pages: 50}},
	book{Title: "Good Book", Author: "Joe", Spec: bookSpecs{Pages: 150, Rating: 4}},
}

func connect(ctx context.Context, url string) (client *mongo.Client, err error) {
	client, err = mongo.NewClient(url)
	if err != nil {
		return
	}
	err = client.Connect(ctx)

	return
}

func prepareCollection(ctx context.Context, collection *mongo.Collection) (err error) {
	// Drop the collection.
	err = collection.Drop(ctx)
	if err != nil {
		return
	}

	// Insert new books into the collection.
	_, err = collection.InsertMany(ctx, books)

	return
}
