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

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
	"github.com/yaacov/tree-search-language/v6/pkg/walkers/ident"
	"github.com/yaacov/tree-search-language/v6/pkg/walkers/semantics"
	"gopkg.in/yaml.v3"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var s []byte
	books := []Book{}

	// Setup the input.
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"author = 'Joe'\")")
	outputPtr := flag.String("o", "json", "output format [json/yaml]")
	flag.Parse()

	// Sanity check.
	if *inputPtr == "" {
		log.Fatal("missing required flag -i (the tsl string to parse)")
	}

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)
	defer tree.Free()

	// Walk the TSL tree and replace identifiers.
	newTree, _, err := ident.Walk(tree, checkColumnName)
	check(err)
	defer newTree.Free()

	// Prepare the books in memory collection.
	err = prepareCollection()
	check(err)

	// Filter the books collection using our transformed TSL tree.
	for _, book := range Books {
		eval := evalFactory(book)

		matchingFilter, err := semantics.Walk(newTree, eval)
		check(err)

		// Convert interface{} to bool
		if match, ok := matchingFilter.(bool); ok && match {
			books = append(books, book)
		}
	}

	// Printout the filtered list.
	switch *outputPtr {
	case "json":
		s, err = json.Marshal(books)
	case "yaml":
		s, err = yaml.Marshal(books)
	default:
		err = fmt.Errorf("unsuported output format: %s", *outputPtr)
	}

	check(err)
	fmt.Printf("%s\n", s)
}
