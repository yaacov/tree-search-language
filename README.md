# tsl

Tree Search Language (TSL) is a wonderful search langauge, implementing query based search engines was never that easy.

TSL include example method `tsl.Walk` for the generation of SQL queries, use the `tsl.Walk` method as a template for 
using the TSL tree to generate filter and search methods over any data source.

[![Go Report Card](https://goreportcard.com/badge/github.com/yaacov/tsl)](https://goreportcard.com/report/github.com/yaacov/tsl)
[![Build Status](https://travis-ci.org/yaacov/tsl.svg?branch=master)](https://travis-ci.org/yaacov/tsl)
[![GoDoc](https://godoc.org/github.com/yaacov/tsl/pkg/tsl?status.svg)](https://godoc.org/github.com/yaacov/tsl/pkg/tsl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

![TSL](/img/search.png?raw=true "TSL Logo")

[ awesome image by [gophers...](https://github.com/egonelbre/gophers) ]

The TSL language grammar is similar to SQL syntax, for example:
```
name like '%joe%' and (city = 'paris' or city = 'milan')
```

###### ParseTSL

The TSL package include the [ParseTSL](https://godoc.org/github.com/yaacov/tsl/pkg/tsl#ParseTSL) method for parsing TSL into a search tree:
```
tree, err := tsl.ParseTSL("name in ('joe', 'jane') and age not between 0 and 10")
```

###### Walk

The TSL package include a helper [Walk](https://godoc.org/github.com/yaacov/tsl/pkg/tsl#Walk) method that adds serch context to Squirrel [SelectBuilder](https://godoc.org/github.com/Masterminds/squirrel#SelectBuilder):
```
sql, args, err := sq.Select("name, city, state").
    Where(tsl.Walk(tree)).
    From("users").
    ToSql()
```

TSL is generated with [Antlr4 tool](https://github.com/antlr/antlr4/), the antlr4 grammar file is [TSL.g4](/TSL.g4),
SQL generation is done using [Squirrel](https://github.com/Masterminds/squirrel) SQL builder.

#### Cli tools

**tls_parser** and **tsl_to_sql** are example cli tools showcasing the TSL language and TSL golang package.

###### tls_parser

``` bash
$ ./tsl_parser -h
Usage of ./tls_parser:
  -i string
    	the tsl string to parse (e.g. "animal = 'kitty'")
  -o string
    	output format [json/yaml/prettyjson] (default "json")
```


``` bash
$ ./tsl_parser -i "(name = 'joe' or name = 'jane') and city = 'rome'" -o yaml
func: $and
left:
  func: $eq
  left: city
  right: rome
right:
  func: $or
  left:
    func: $eq
    left: name
    right: jane
  right:
    func: $eq
    left: name
    right: joe
```

###### tsl_to_sql

``` bash
$ ./tsl_to_sql -h
Usage of ./tsl_to_sql:
  -i string
    	the tsl string to parse (e.g. "animal = 'kitty'")
  -o string
    	output format [mysql/pgsql] (default "mysql")
  -t string
    	the table name to use for SQL generation (default "<table-name>")

```

``` bash
$ ./tsl_to_sql -i "name != 'eli''s' or city like '%rome%' and state not between 'italy' and 'france'" -t users -o pgsql
sql:  SELECT * FROM users WHERE (name <> $1 OR (city LIKE $2 AND state NOT BETWEEN $3 AND $4))
args: [eli's %rome% italy france]

```

#### Language

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

###### import

``` go
import (
  ...
	"github.com/yaacov/tsl/pkg/tsl"
  ...
)
```

###### ParseTSL

``` go
  ...
  // Set a TSL input string
  input = "name='joe' or name='jane'"

  // Parse input string into a TSL tree
  tree, err := tsl.ParseTSL(input)
  ...
```

###### Walk

``` go
  import (
    ...
    sq "github.com/Masterminds/squirrel"
    "github.com/yaacov/tsl/pkg/tsl"
    ...
  )
  ...
  // Convert TSL tree into SQL string using squirrel sql builder.
  sql, args, err := sq.Select("name, city, state").
    Where(tsl.Walk(tree)).
    From("users").
    ToSql()
  ...
```
