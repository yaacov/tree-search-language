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

package main

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/yaacov/tsl/cmd/model"
)

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
	_, err = collection.InsertMany(ctx, model.Books)

	return
}
