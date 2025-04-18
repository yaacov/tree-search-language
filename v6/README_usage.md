# TSL Usage Guide

A set examples showing how to use TSL in various scenarios.

---

## 1. Parsing and validating filter expressions

Use case: receive a filter string from a REST API, validate its syntax, and print a human‐readable tree or report errors.

```go
import (
  "fmt"
  "github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

func parseFilter(input string) {
  tree, err := tsl.ParseTSL(input)
  if err != nil {
    // SyntaxError contains Message, Position, Context
    fmt.Printf("Invalid filter: %s\n", err)
    return
  }
  defer tree.Free()

  // If we reach here, syntax is valid; tree can be marshaled to JSON/YAML
  data, _ := tree.MarshalJSON()
  fmt.Printf("Parsed filter tree:\n%s\n", data)
}
```

**Explanation**  
- `ParseTSL` returns a syntax‐checked AST.  
- On error, you get precise position and context.  
- A valid tree can be serialized for debugging or logging.

---

## 2. In‑memory record filtering

Use case: filter a slice of user records in Go without touching a database.

```go
import "github.com/yaacov/tree-search-language/v6/pkg/walkers/semantics"

records := []map[string]interface{}{
  {"name":"alice","age":30},
  {"name":"bob","age":25},
}

tree, _ := tsl.ParseTSL("age >= 28")
defer tree.Free()

for _, rec := range records {
  eval := func(field string) (interface{}, bool) {
    val, ok := rec[field]
    return val, ok
  }
  match, _ := semantics.Walk(tree, eval)
  if match.(bool) {
    // rec passed the filter
    fmt.Println("Matches:", rec)
  }
}
```

**Explanation**  
- Supply a lookup function (`eval`) that returns `value, ok`.  
- `Walk` applies the expression tree to each record.  
- Great for in‑process filtering of JSON, CSV, or config objects.

---

## 3. Generating SQL WHERE clauses

Use case: convert a user‐supplied filter into a parameterized WHERE clause using Squirrel.

```go
import (
  sq "github.com/Masterminds/squirrel"
  "github.com/yaacov/tree-search-language/v6/pkg/walkers/sql"
)

tree, _ := tsl.ParseTSL("status = 'active' AND score > 50")
defer tree.Free()

filter, _ := sql.Walk(tree) // returns a squirrel.Sqlizer
query, args, _ := sq.Select("*").
  From("players").
  Where(filter).
  ToSql()

// query: "SELECT * FROM players WHERE (status = ? AND score > ?)"
// args:  ["active", 50]
```

**Explanation**  
- `sql.Walk` produces a Squirrel filter object.  
- The returned SQL is safe against injection (parameters are placeholders).  
- You can tack this onto any SELECT/UPDATE/DELETE builder.

---

## 4. Visualizing expression trees

Use case: debug complex filters by exporting to Graphviz DOT and rendering a diagram.

```go
import (
  "fmt"
  "github.com/yaacov/tree-search-language/v6/pkg/walkers/graphviz"
)

tree, _ := tsl.ParseTSL("name LIKE '%doe%' OR (age BETWEEN 20 AND 30)")
defer tree.Free()

dotContent, _ := graphviz.Walk("", tree, "FilterTree")
dot := fmt.Sprintf("digraph FilterTree {\n%s\n}", dotContent)
fmt.Println(dot)
// Pipe output into `dot -Tpng` to see the visual tree
```

**Explanation**  
- `graphviz.Walk` returns a snippet of DOT nodes and edges.  
- Wrap it in a `digraph { … }` block.  
- Use standard Graphviz tools to generate PNG/SVG for documentation.

---

## 5. Renaming fields (identifier mapping)

Use case: let end users filter on logical names, map them to actual database or JSON fields.

```go
import "github.com/yaacov/tree-search-language/v6/pkg/walkers/ident"

mapping := map[string]string{
  "user":    "customers.name",
  "balance": "accounts.current_balance",
}

mapper := func(id string) (string, error) {
  if col, ok := mapping[id]; ok {
    return col, nil
  }
  return "", fmt.Errorf("unknown field: %s", id)
}

tree, _ := tsl.ParseTSL("user = 'alice' AND balance > 1000")
defer tree.Free()

newTree, _ := ident.Walk(tree, mapper)
defer newTree.Free()

// newTree now uses "customers.name" and "accounts.current_balance"
```

**Explanation**  
- `ident.Walk` clones the AST and applies your mapping function to each identifier.  
- Invalid identifiers cause an error early in the pipeline.  
- Ideal for decoupling front‑end field names from internal schemas.

---
