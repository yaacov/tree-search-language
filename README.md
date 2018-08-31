# Tree Search Language (TSL)

Tree Search Language (TSL) is a wonderful search langauge, With similar grammar to SQL's
where part. implementing query based search engines was never that easy.

[![Go Report Card](https://goreportcard.com/badge/github.com/yaacov/tsl)](https://goreportcard.com/report/github.com/yaacov/tsl)
[![Build Status](https://travis-ci.org/yaacov/tsl.svg?branch=master)](https://travis-ci.org/yaacov/tsl)
[![GoDoc](https://godoc.org/github.com/yaacov/tsl/pkg/tsl?status.svg)](https://godoc.org/github.com/yaacov/tsl/pkg/tsl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

![TSL](/img/search.png?raw=true "TSL Logo")

[ awesome image by [gophers...](https://github.com/egonelbre/gophers) ]

The TSL language grammar is similar to SQL syntax, for example:
``` sql
name like '%joe%' and (city = 'paris' or city = 'milan')
```
``` sql
name in ('joe', 'jane') and grade not between 0 and 50
```

### ParseTSL

The TSL package include the [ParseTSL](https://godoc.org/github.com/yaacov/tsl/pkg/tsl#ParseTSL) method for parsing TSL into a search tree:
``` go
tree, err := tsl.ParseTSL("name in ('joe', 'jane') and grade not between 0 and 50")
```

After parsing the TSL tree will look like this:
``` bash
$and
|
+----------------------------+
|                            |
$in                          $nbetween
|                            |
+---------+                  +--------+
|         |                  |        |
"name"  ["joe", "jane"]      "grade"    [0, 50]
```

### SquirrelWalk

The TSL package include a helper [SquirrelWalk](/pkg/tsl/tsl_squirrel_walk.go) method that adds search to [squirrel](https://github.com/Masterminds/squirrel)'s SelectBuilder object:

``` go
filter, err := tsl.SquirrelWalk(tree)

sql, args, err := sq.Select("name, city, state").
    From("users").
    Where(filter).
    ToSql()
```

After SQL generation the `sql` var will be:
``` sql
SELECT name, city, state FROM users WHERE name IN (?, ?) AND grade NOT BETWEEN ? AND ?
```

### BSONWalk

The TSL package include a helper [BSONWalk](/pkg/tsl/tsl_bson_walk.go) method that adds search bson filter to [mongo-go-driver](https://github.com/mongodb/mongo-go-driver):

``` go
// Prepare a bson filter
filter, err = tsl.BSONWalk(tree)

// Run query
cur, err := collection.Find(ctx, bson.NewDocument(filter))
```

### Antlr4 grammar

TSL is generated with [Antlr4 tool](https://github.com/antlr/antlr4/), the antlr4 grammar file is [TSL.g4](/TSL.g4),
SQL generation is done using [Squirrel](https://github.com/Masterminds/squirrel) SQL builder.

## Cli tools

[tls_parser](/cmd/tsl_parser), [tls_mongo](/cmd/tsl_mongo) and [tsl_to_sql](/cmd/tsl_to_sql) are example cli tools showcasing the TSL language and TSL golang package.

### tls_parser

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

### tsl_mongo

tsl_mongo include an example using [BSONWalk](/pkg/tsl/tsl_bson_walk.go) method, for building a mongo bson filter.

``` bash
$ ./tsl_mongo -h
Usage of ./tsl_mongo:
  -c string
    	collection name to query on (default "books")
  -d string
    	db name to connect to (default "tsl")
  -i string
    	the tsl string to parse (e.g. "author = 'Jane'") (default "title is not null")
  -p	prepare a book collection for queries
  ...
  -u string
    	url for mongo server (default "mongodb://localhost:27017")
```

``` bash
$ ./tsl_mongo -p -i "title is not null"
{"_id":{"$oid":"5b8083aadda18596ef42a941"},"title":"Book","author":"Joe","spec":{"pages":{"$numberLong":"100"},"raiting":{"$numberLong":"4"}}}
{"_id":{"$oid":"5b8083aadda18596ef42a942"},"title":"Other Book","author":"Jane","spec":{"pages":{"$numberLong":"200"},"raiting":{"$numberLong":"3"}}}
{"_id":{"$oid":"5b8083aadda18596ef42a943"},"title":"Some Other Book","author":"Jane","spec":{"pages":{"$numberLong":"50"},"raiting":{"$numberLong":"5"}}}
{"_id":{"$oid":"5b8083aadda18596ef42a944"},"title":"Some Other Book","author":"Jane","spec":{"pages":{"$numberLong":"50"}}}
```
``` bash
$ ./tsl_mongo -i "title ~= 'Other' and spec.raiting > 4"
{"_id":{"$oid":"5b8083c9e0d411d1f2fbcfa4"},"title":"Some Other Book","author":"Jane","spec":{"pages":{"$numberLong":"50"},"raiting":{"$numberLong":"5"}}}
```

### tsl_to_sql

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

## Grammar

### Keywords
```
and or not is null like between in
```
### Operators
```
= <= >= != ~= ~! <>
```
### Examples
```
name = 'Joe'
```
```
city in ('paris', 'rome', 'milan') or sate = 'spain'
```
```
(name = 'joe' or city = 'rome') and state = 'italy'
```

## Code snippets

### import

``` go
import (
  ...
	"github.com/yaacov/tsl/pkg/tsl"
  ...
)
```

### ParseTSL

``` go
  ...
  // Set a TSL input string
  input = "name='joe' or name='jane'"

  // Parse input string into a TSL tree
  tree, err := tsl.ParseTSL(input)
  ...
```

### SquirrelWalk

``` go
  import (
    ...
    sq "github.com/Masterminds/squirrel"
    "github.com/yaacov/tsl/pkg/tsl"
    ...
  )
  ...
  // Set filter
  filter, err := tsl.SquirrelWalk(tree)

  // Convert TSL tree into SQL string using squirrel sql builder
  sql, args, err := sq.Select("name, city, state").
    From("users").
    Where(filter).
    ToSql()
  ...
```
