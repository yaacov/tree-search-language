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

package tsl

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/yaacov/tree-search-language/pkg/parser"
)

// ParseTSL parses the input string into TSL tree.
func ParseTSL(input string) (tree Node, err error) {
	// Setup the ErrorListener.
	errorListener := NewErrorListener()

	// Setup the input.
	is := antlr.NewInputStream(input)

	// Create the Lexer.
	lexer := parser.NewTSLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser.
	p := parser.NewTSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse the expression (by walking the tree).
	var listener Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// Check for errors.
	err = errorListener.Err
	if err != nil {
		return
	}

	// Get the parsed tree.
	tree, err = listener.GetTree()

	return
}
