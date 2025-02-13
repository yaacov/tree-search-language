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
	"encoding/json"
	"flag"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"gopkg.in/yaml.v3"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
	"github.com/yaacov/tree-search-language/v6/pkg/walkers/graphviz"
	"github.com/yaacov/tree-search-language/v6/pkg/walkers/sql"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var s []byte
	var st string

	// Setup the input.
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"title = 'kitty'\")")
	outputPtr := flag.String("o", "json", "output format [json/yaml/prettyjson/sql/dot]")
	tablePtr := flag.String("t", "table_name", "table name for SQL output")
	flag.Parse()

	// Sanity check.
	if *inputPtr == "" {
		log.Fatal("missing required flag -i (the tsl string to parse)")
	}

	// Parse input string into a TSL tree.
	tree, err := tsl.ParseTSL(*inputPtr)
	check(err)
	defer tree.Free()

	switch *outputPtr {
	case "json":
		s, err = json.Marshal(tree)
	case "yaml":
		s, err = yaml.Marshal(tree)
	case "dot":
		st, err = graphviz.Walk("", tree, "root")
		s = []byte(fmt.Sprintf("digraph {\n%s\n}\n", st))
	case "sql":
		var sqlStr string
		var args []interface{}
		var filter sq.Sqlizer

		// Use Squirrel to walk the tree, and create SQL filter.
		filter, err = sql.Walk(tree)
		check(err)

		// Add SQL template to bytes slice.
		sqlStr, args, err = sq.
			Select("*").
			Where(filter).
			From(*tablePtr).
			ToSql()
		check(err)

		s = []byte(fmt.Sprintf("sql:  %s\nargs: %v", sqlStr, args))
	default:
		log.Fatalf("unknown output format: %s", *outputPtr)
	}

	check(err)
	fmt.Println(string(s))
}
