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
// Author: 2019 Yaacov Zamir <kobi.zamir@gmail.com>

package main

import (
	"github.com/yaacov/tsl/pkg/walkers/semantics"

	"github.com/yaacov/tsl/cmd/model"
)

// Books are the demo list of books.
var Books = []semantics.Doc{}

func prepareCollection() (err error) {
	// Insert new books into the table.
	for _, b := range model.Books {
		// Create a new book.
		newBook := semantics.Doc{
			"title":  b.(model.Book).Title,
			"author": b.(model.Book).Author,
		}

		// Add optional parameters.
		if b.(model.Book).Spec.Pages > 0 {
			newBook["spec.pages"] = b.(model.Book).Spec.Pages
		}
		if b.(model.Book).Spec.Rating > 0 {
			newBook["spec.rating"] = b.(model.Book).Spec.Rating
		}

		// Insert new book to the books arra.
		Books = append(Books, newBook)
	}

	return
}
