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

// Package main.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/yaacov/tree-search-language/pkg/tsl"
	"github.com/yaacov/tree-search-language/pkg/walkers/ident"
	"github.com/yaacov/tree-search-language/pkg/walkers/semantics"
	yaml "gopkg.in/yaml.v2"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// columnNamesMap mapps between user namespace and the document field names.
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
	var s []byte
	books := []Book{}

	// Setup the input.
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"author = 'Joe'\")")
	outputPtr := flag.String("o", "json", "output format [json/yaml/prettyjson]")
	flag.Parse()

	// Sanity check.
	if *inputPtr == "" {
		log.Fatal("missing required flag -i (the tsl string to parse)")
	}

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)

	// Check and replace user identifiers with the document field names.
	tree, err = ident.Walk(tree, checkColumnName)
	check(err)

	// Prepare the books in memeory collection.
	err = prepareCollection()
	check(err)

	// Filter the books collection using our stl tree.
	for _, book := range Books {
		matchingFilter, err := semantics.Walk(tree, evalFactory(book))
		check(err)
		if matchingFilter {
			books = append(books, book)
		}
	}

	// Printout the filtered list.
	switch *outputPtr {
	case "json":
		s, err = json.Marshal(books)
	case "yaml":
		s, err = yaml.Marshal(books)
	case "prettyjson":
		s, err = prettyjson.Marshal(books)
	default:
		err = fmt.Errorf("unsuported output format: %s", *outputPtr)
	}

	check(err)
	fmt.Printf("%s\n", s)
}
