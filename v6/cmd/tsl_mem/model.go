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
	"fmt"

	"github.com/yaacov/tree-search-language/v6/cmd/model"
	"github.com/yaacov/tree-search-language/v6/pkg/walkers/semantics"
)

// Book represent one book in our in-memmory data base.
type Book map[string]interface{}

// Books are the demo list of books.
var Books = []Book{}

// columnNamesMap mapps between user namespace and the document field names.
var columnNamesMap = map[string]string{
	"title":       "title",
	"author":      "author",
	"spec.pages":  "spec.pages",
	"spec.rating": "spec.rating",
	"on_loan":     "onloan",
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

func evalFactory(book Book) semantics.EvalFunc {
	return func(k string) (interface{}, bool) {
		// First try to get mapped column name
		mappedKey, err := checkColumnName(k)
		if err == nil {
			if v, ok := book[mappedKey]; ok {
				return v, true
			}
		}
		return nil, false
	}
}

func prepareCollection() (err error) {
	// Insert new books into the table.
	for _, b := range model.Books {
		// Create a new book.
		newBook := Book{
			"title":  b.(model.Book).Title,
			"author": b.(model.Book).Author,
			"onloan": b.(model.Book).OnLoan,
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
