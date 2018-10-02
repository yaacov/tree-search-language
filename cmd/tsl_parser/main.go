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
	"encoding/json"
	"flag"
	"fmt"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/hokaccha/go-prettyjson"
	"gopkg.in/yaml.v2"

	"github.com/yaacov/tsl/pkg/tsl"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
}

func main() {
	var s []byte

	// Setup the input.
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"Title = 'kitty'\")")
	outputPtr := flag.String("o", "json", "output format [json/yaml/prettyjson/sql]")
	flag.Parse()

	// Sanity check.
	if *inputPtr == "" {
		err := fmt.Errorf("missing required flag -i (the tsl string to parse)")
		check(err)
	}

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)

	// If listener has errors, we can not print the tree.
	if err != nil {
		os.Exit(1)
	}

	switch *outputPtr {
	case "json":
		s, err = json.Marshal(tree)
	case "yaml":
		s, err = yaml.Marshal(tree)
	case "prettyjson":
		s, err = prettyjson.Marshal(tree)
	case "sql":
		var sql string
		var args []interface{}
		var filter sq.Sqlizer

		// Use Squirrel to walk the tree, and create SQL filter.
		filter, err = tsl.SquirrelWalk(tree)

		// Add SQL template to bytes slice.
		if err == nil {
			sql, args, err = sq.
				Select("*").
				Where(filter).
				From("table_name").
				ToSql()

			s = append(s, fmt.Sprintf("sql:  %s\n", sql)...)
			s = append(s, fmt.Sprintf("args: %v", args)...)
		}
	}

	check(err)
	fmt.Println(string(s))
}
