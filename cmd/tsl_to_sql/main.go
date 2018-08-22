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

// Package main
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/yaacov/tsl/pkg/parser"
	"github.com/yaacov/tsl/pkg/tsl"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}
}

func main() {
	var err error
	var tree tsl.Node
	var sql string
	var args []interface{}

	// Setup the input
	inputPtr := flag.String("i", "", "the tsl string to parse (e.g. \"animal = 'kitty'\")")
	tablePtr := flag.String("t", "<table-name>", "the table name to use for SQL generation")
	outputPtr := flag.String("o", "mysql", "output format [mysql/pgsql]")
	flag.Parse()

	// Setup the ErrorListener
	errorListener := tsl.NewErrorListener()

	// Setup the input
	is := antlr.NewInputStream(*inputPtr)

	// Create the Lexer
	lexer := parser.NewTSLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Finally parse the expression (by walking the tree)
	var listener tsl.Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	err = errorListener.Err
	check(err)

	tree, err = listener.GetTree()
	check(err)

	// If listener has erros, we can not print the tree
	if err != nil {
		os.Exit(1)
	}

	switch *outputPtr {
	case "pgsql":
		sql, args, err = tsl.ToSelectBuilder(tree, true).
			From(*tablePtr).
			ToSql()
	case "mysql":
		sql, args, err = tsl.ToSelectBuilder(tree, false).
			From(*tablePtr).
			ToSql()
	default:
		sql, args, err = tsl.ToSelectBuilder(tree, false).
			From(*tablePtr).
			ToSql()
	}

	check(err)
	fmt.Printf("sql:  %s\n", sql)
	fmt.Printf("args: %v\n", args)
}
