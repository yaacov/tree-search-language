# tsl
Tree Search Language (TSL) is a simple SQL like langauge

[![Go Report Card](https://goreportcard.com/badge/github.com/yaacov/tsl)](https://goreportcard.com/report/github.com/yaacov/tsl)
[![Build Status](https://travis-ci.org/yaacov/tsl.svg?branch=master)](https://travis-ci.org/yaacov/tsl)
[![GoDoc](https://godoc.org/github.com/yaacov/tsl/pkg/tsl?status.svg)](https://godoc.org/github.com/yaacov/tsl/pkg/tsl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


#### Cli tool
``` bash
$ ./tls_parser -h
Usage of ./tls_parser:
  -i string
    	the tsl string to parse (e.g. "animal = 'kitty'")
  -o string
    	output format [json/yaml/prettyjson] (default "json")
```

``` bash
$ ./tls_parser -i "(name = 'joe' or name = 'jane') and city = 'rome'" -o yaml
func: $and
left:
  func: $eq
  left: city
  right: '''rome'''
right:
  func: $or
  left:
    func: $eq
    left: name
    right: '''jane'''
  right:
    func: $eq
    left: name
    right: '''joe'''
```
#### Language

TSL is generated using Antlr4 tool, the grammer file used for generation is [TSL.g4](/TSL.g4).

###### Keywords
```
and or not is null like between in 
```
###### Operators
```
= <= >= != ~= ~! <>
```
###### Examples
```
name = 'Joe'
```
```
city in ('paris', 'rome', 'milan') or sate = 'spain'
```
```
(name = 'joe' or city = 'rome') and state = 'italy'
```

#### Code snippets

``` go
import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/yaacov/tsl/pkg/parser"
	"github.com/yaacov/tsl/pkg/tsl"
)
```
``` go
// Setup the input
is := antlr.NewInputStream(input)

// Create the Lexer
lexer := parser.NewTSLLexer(is)
stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

// Create the Parser
p := parser.NewTSLParser(stream)

// Finally parse the expression (by walking the tree)
var listener tsl.Listener
antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

// Get the parsed tree
n, _ = listener.GetTree()

/*
  input := "a = 'hello'"
  expected := `{"func":"$eq","right":"'hello'","left":"a"}`
*/
```

