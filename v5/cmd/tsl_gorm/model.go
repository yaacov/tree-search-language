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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/yaacov/tree-search-language/v5/cmd/model"
)

// Book represent a book model.
type Book struct {
	gorm.Model
	Title  string
	Author string
	Pages  uint
	Rating uint
	OnLoan bool
}

func connect(ctx context.Context, url string) (tx *gorm.DB, err error) {
	db, err := gorm.Open("sqlite3", url)
	check(err)

	// Migrate the schema.
	db.AutoMigrate(&Book{})

	// Begin transaction.
	tx = db.Begin()

	return
}

func prepareCollection(tx *gorm.DB) (err error) {
	// Delete all books in the DB.
	tx.Where("title is not null").Delete(Book{})

	// Insert new books into DB.
	for _, b := range model.Books {
		tx.Create(&Book{
			Title:  b.(model.Book).Title,
			Author: b.(model.Book).Author,
			Pages:  b.(model.Book).Spec.Pages,
			Rating: b.(model.Book).Spec.Rating,
			OnLoan: b.(model.Book).OnLoan,
		})

		if tx.Error != nil {
			return
		}
	}

	return
}
