

<p align="center">
  <img src="https://raw.githubusercontent.com/yaacov/tsl/master/img/search-248.png" alt="TSL Logo">
</p>

# Tree Search Language (TSL)

Tree Search Language (TSL) is a wonderful search language, With similar grammar to SQL's
where part. implementing query based search engines was never that easy.

[![Go Report Card](https://goreportcard.com/badge/github.com/yaacov/tsl)](https://goreportcard.com/report/github.com/yaacov/tsl)
[![Build Status](https://travis-ci.org/yaacov/tsl.svg?branch=master)](https://travis-ci.org/yaacov/tsl)
[![GoDoc](https://godoc.org/github.com/yaacov/tsl/pkg/tsl?status.svg)](https://godoc.org/github.com/yaacov/tsl/pkg/tsl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

The TSL language grammar is similar to SQL syntax.

## Code examples

For complete working code examples see the cli tools direcotry in the [/cmd](/cmd).

## Install

#### Installing the different packages

See some code snippets [here](https://github.com/yaacov/tsl#code-snippets) and [here](https://github.com/yaacov/tsl#sqlwalk).

``` bash
# Install the base package
go get "github.com/yaacov/tsl/pkg/tsl"

# Install all walkers
go get "github.com/yaacov/tsl/pkg/walkers/..."

# Or pick the walker needed
go get "github.com/yaacov/tsl/pkg/walkers/sql"
go get "github.com/yaacov/tsl/pkg/walkers/mongo"
go get "github.com/yaacov/tsl/pkg/walkers/ident"
go get "github.com/yaacov/tsl/pkg/walkers/graphviz"
```

#### Installing the command line examples

See CLI tools usage [here](https://github.com/yaacov/tsl#cli-tools).

``` bash
go get -v "github.com/yaacov/tsl/cmd/tsl_parser"
go get -v "github.com/yaacov/tsl/cmd/tsl_mongo"
go get -v "github.com/yaacov/tsl/cmd/tsl_sqlite"
```

## Syntax examples

Images created using the `tsl_parser` CLI example and Graphviz's `dot` utility:
``` bash
./tsl_parser -i "name like '%joe%' and (city = 'paris' or city = 'milan')" -o dot > file.dot

dot file.dot -Tpng > image.png
```

#### Operator precedence
``` sql
name like '%joe%' and (city = 'paris' or city = 'milan')
```
![TSL](/img/example_a.png?raw=true "example tree")

#### Operators with multiple arguments
``` sql
name in ('joe', 'jane') and grade not between 0 and 50
```
![TSL](/img/example_b.png?raw=true "example tree")

#### Math operators
``` sql
memory.total - memory.cache > 2000 and cpu.usage > 50
```
![TSL](/img/example_c.png?raw=true "example tree")

#### More math operators
``` sql
(net.rx + net.tx) / 1000 > 3 or net.rx / 1000 > 6
```
![TSL](/img/example_d.png?raw=true "example tree")


For code examples see the cli tools in the [/cmd](/cmd) direcotry.

## Usage examples

##### ParseTSL

The TSL package include the [ParseTSL](https://godoc.org/github.com/yaacov/tsl/pkg/tsl#ParseTSL) method for parsing TSL into a search tree:
``` go
tree, err := tsl.ParseTSL("name in ('joe', 'jane') and grade not between 0 and 50")
```

After parsing the TSL tree will look like this (image created using the `tsl_parser` cli utility using `.dot` output option):

![TSL](/img/example01.png?raw=true "example tree")

##### sql.Walk

The TSL package include a helper [sql.Walk](/pkg/walkers/sql/walk.go) method that adds search to [squirrel](https://github.com/Masterminds/squirrel)'s SelectBuilder object:

``` go
import (
    ...
    "github.com/yaacov/tsl/pkg/walkers/sql"
    ...
)

// Prepare squirrel filter.
filter, err := sql.Walk(tree)

// Create an SQL query.
sql, args, err := sq.Select("name", "city", "state").
    From("users").
    Where(filter).
    ToSql()
```

After SQL generation the `sql` and `args` vars will be:
``` sql
SELECT name, city, state FROM users WHERE (name IN (?,?) AND grade NOT BETWEEN ? AND ?)
```

``` json
["joe", "jane", 0, 50]

```

##### mongo.Walk

The TSL package include a helper [mongo.Walk](/pkg/walkers/mongo/walk.go) method that adds search bson filter to [mongo-go-driver](https://github.com/mongodb/mongo-go-driver):

``` go
import (
    ...
    "github.com/yaacov/tsl/pkg/walkers/mongo"
    ...
)

// Prepare a bson filter.
filter, err = mongo.Walk(tree)

// Run query.
cur, err := collection.Find(ctx, filter)
```

##### graphviz.Walk

The TSL package include a helper [graphviz.Walk](/pkg/walkers/graphviz/walk.go) method that exports `.dot` file nodes :

``` go
import (
    ...
    "github.com/yaacov/tsl/pkg/walkers/graphviz"
    ...
)

// Prepare .dot file nodes as a string.
s, err = graphviz.Walk("", tree, "")

// Wrap the nodes in a digraph wrapper.
s = fmt.Sprintf("digraph {\n%s\n}\n", s)
```

##### ident.Walk

The TSL package include a helper [ident.Walk](/pkg/walkers/ident/walk.go) method that checks and mapps identifier names:

``` go
import (
    ...
    "github.com/yaacov/tsl/pkg/walkers/ident"
    ...
)
...

// columnNamesMap mapps between user namespace and the SQL column names.
var columnNamesMap = map[string]string{
	"title":       "title",
	"author":      "author",
	"spec.pages":  "pages",
	"spec.rating": "rating",
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
...

// Check and replace user identifiers with the SQL table column names.
tree, err = ident.Walk(tree, checkColumnName)
...
```

## Cli tools

[tls_parser](/cmd/tsl_parser), [tls_mongo](/cmd/tsl_mongo) and [tsl_sqlite](/cmd/tsl_sqlite) are example cli tools showcasing the TSL language and TSL golang package.

##### tls_parser

``` bash
$ ./tsl_parser -h
Usage of ./tls_parser:
  -i string
    	the tsl string to parse (e.g. "animal = 'kitty'")
  -o string
    	output format [json/yaml/prettyjson/sql/dot] (default "json")
```


``` bash
$ ./tsl_parser -i "(name = 'joe' or name = 'jane') and city = 'rome'" -o sql
```
```
sql:  SELECT * FROM table_name WHERE ((name = ? OR name = ?) AND city = ?)
args: [joe jane rome]
```

``` bash
./tsl_parser -i "(name = 'joe' or name = 'jane') and city = 'rome'" -o prettyjson
```

``` json
{
  "func": "$and",
  "left": {
    "func": "$or",
    "left": {
      "func": "$eq",
      "left": {
        "func": "$ident",
        "left": "name"
      },
      "right": {
        "func": "$string",
        "left": "joe"
      }
    },
    "right": {
      "func": "$eq",
      "left": {
        "func": "$ident",
        "left": "name"
      },
      "right": {
        "func": "$string",
        "left": "jane"
      }
    }
  },
  "right": {
    "func": "$eq",
    "left": {
      "func": "$ident",
      "left": "city"
    },
    "right": {
      "func": "$string",
      "left": "rome"
    }
  }
}
```

``` bash
./tsl_parser -i "city = 'rome'" -o dot
```

``` dot
digraph {
root [shape=box color=black label="$eq"]
XVlB [shape=record color=red label="$ident | 'city'" ]
zgba [shape=record color=blue label="$string | 'rome'" ]
root -> { XVlB, zgba }
}
```

##### tsl_mongo

tsl_mongo include an example using [mongo.NWalk](/pkg/walkers/mongo/walk.go) method, for building a mongo bson filter.

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
$ ./tsl_mongo -p -i "title is not null" | jq
```
``` json
{
  "title": "Book",
  "author": "Joe",
  "spec": {
    "pages": 100,
    "rating": 4
  }
}
```
``` bash
$ ./tsl_mongo -i "title ~= 'Other' and spec.rating > 1" | jq
```
``` json
{
  "title": "Other Book",
  "author": "Jane",
  "spec": {
    "pages": 200,
    "rating": 3
  }
}
```

##### tsl_sqlite

``` bash
./tsl_sqlite -h
Usage of ./tsl_sqlite:
  -f string
    	the sqlite database file name (default "./sqlite.db")
  -i string
    	the tsl string to parse (e.g. "Title = 'Book'")
  -p	prepare a book collection for queries
```

``` bash
$ SQL="title like '%Book%' and spec.pages > 100"
$ ./tsl_sqlite -i "$SQL" -p | jq
```
``` json
{
  "title": "Other Book",
  "author": "Jane",
  "spec": {
    "pages": 200,
    "rating": 3
  }
}
{
  "title": "Good Book",
  "author": "Joe",
  "spec": {
    "pages": 150,
    "rating": 4
  }
}
```

## Grammar

##### Antlr4 grammar

TSL is generated using [Antlr4 tool](https://github.com/antlr/antlr4/), the antlr4 grammar file is [TSL.g4](/TSL.g4).

##### Keywords
```
and or not is null like between in
```
##### Operators
```
= <= >= != ~= ~! <> + - * / %
```
##### Examples
```
name = 'Joe'
```
```
city in ('paris', 'rome', 'milan') or sate = 'spain'
```
```
(name = 'joe' or city = 'rome') and state = 'italy'
```
```
net.tx + net.rx > 2000 or mem.total - mem.usage < 1000
```

## Code snippets


``` go
import "github.com/yaacov/tsl/pkg/tsl"
```

##### ParseTSL

ParseTSL takes a string input and generate a search tree object, the function
returns the root Node of the tree.

``` go
...
// Set a TSL input string.
input = "name='joe' or name='jane'"

// Parse input string into a TSL tree.
tree, err := tsl.ParseTSL(input)
...
```

sql.Walk and mongo.Walk are example methods the demonstrate traversing ( walk ) the search tree.

##### sql.Walk

sql.Walk takes the base Node ( tree ) of the search tree, and return a Squirrel SQL filter object.

``` go
import (
    ...
    sq "github.com/Masterminds/squirrel"
    "github.com/yaacov/tsl/pkg/tsl"
    "github.com/yaacov/tsl/pkg/walkers/sql"
    ...
)
...
// Set filter.
filter, err := sql.Walk(tree)

// Convert TSL tree into SQL string using squirrel select builder.
sql, args, err := sq.Select("name, city, state").
    From("users").
    Where(filter).
    ToSql()
...
```

##### mongo.Walk

mongo.Walk takes the base Node ( tree ) of the search tree, and return a MongoDB BSON object.

``` go
...
import (
  ...
  "github.com/yaacov/tsl/pkg/walkers/mongo"
  ...
)
...
// Prepare a bson filter.
filter, err = mongo.Walk(tree)

// Run query.
cur, err := collection.Find(ctx, filter)
defer cur.Close(ctx)

// Loop on query elements.
for cur.Next(ctx) {
...
```

##### ident.Walk

ident.Walk helps to check validity of identifiers and replace them if necessary.

``` go
...
import (
  ...
  "github.com/yaacov/tsl/pkg/walkers/ident"
  ...
)
...
// check that the column idetifier is valid.
// It returns a mapped valid column name, and an error if name is not valid.
func check(s string) (string, error) {
  return s, nil
}
...
// Check and replace user identifiers with the SQL table column names.
tree, err = ident.Walk(tree, check)
...
```

[ awesome logo image by [gophers...](https://github.com/egonelbre/gophers) ]
