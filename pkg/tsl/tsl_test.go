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
	"encoding/json"
	"strings"
	"testing"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/yaacov/tsl/pkg/parser"
)

// Strip white spaces.
func removeWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

// parseTSL parse the TSL.
func parseTSL(input string) (n Node, err error) {
	// Setup the TSL ErrorListener
	errorListener := NewErrorListener()

	// Setup the input
	is := antlr.NewInputStream(input)

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
	var listener Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	err = errorListener.Err
	if err != nil {
		return
	}

	n, err = listener.GetTree()

	return
}

func TestListener(t *testing.T) {
	// Test valid string
	input := "a = 'hello'"

	// Test TSL parser
	n, err := parseTSL(input)
	if err != nil {
		t.Fail()
	}

	// Test json output
	expected := `
		{"func":"$eq","left":"a","right":"hello"}
	`
	expected = removeWhitespace(expected)
	s, _ := json.Marshal(n)
	if string(s) != expected {
		t.Fatalf("expected %s instead it was %s", expected, string(s))
		t.Fail()
	}
}

func TestListenerOpOrder(t *testing.T) {
	// Test valid string
	input := "a = 12.3e1 or b = 'world' and c = 'hello'"

	// Test TSL parser
	n, err := parseTSL(input)
	if err != nil {
		t.Fail()
	}

	// Test json output
	expected := `
		{"func":"$or","left":{"func":"$eq","left":"a","right":123},
		"right":{"func":"$and","left":{"func":"$eq","left":"b","right":"world"},
		"right":{"func":"$eq","left":"c","right":"hello"}}}
	`
	expected = removeWhitespace(expected)
	s, _ := json.Marshal(n)
	if string(s) != expected {
		t.Fatalf("expected %s instead it was %s", expected, string(s))
		t.Fail()
	}
}
