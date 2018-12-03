

<p align="center">
  <img src="https://raw.githubusercontent.com/yaacov/tsl/master/img/search-248.png" alt="TSL Logo">
</p>

# Tree Search Language (TSL) Imperial Walkers (IW)

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
