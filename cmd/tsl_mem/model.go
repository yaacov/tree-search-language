// Copyright 2019 Yaacov Zamir <kobi.zamir@gmail.com>
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

// Author: 2019 Nimrod Shneor <nimrodshn@gmail.com>

package main

import "github.com/yaacov/tsl/cmd/model"

// Book represent one book in our in memory data table.
type Book map[string]interface{}

// Books are the demo list of books.
var Books = []Book{}

func prepareCollection() (err error) {
	// Insert new books into the table.
	for _, b := range model.Books {
		Books = append(Books, Book{
			"title":       b.(model.Book).Title,
			"author":      b.(model.Book).Author,
			"spec.pages":  b.(model.Book).Spec.Pages,
			"spec.rating": b.(model.Book).Spec.Rating,
		})
	}

	return
}
