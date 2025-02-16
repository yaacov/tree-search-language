

<p align="center">
  <img src="https://raw.githubusercontent.com/yaacov/tree-search-language/master/v6/img/search-248.png" alt="TSL Logo">
</p>

# Tree Search Language (TSL) Imperial Walkers (IW)

##### sql

The `sql` package include a helper `sql.Walk` ([code](/pkg/walkers/sql/walk.go), [doc](https://pkg.go.dev/github.com/yaacov/tree-search-language/v5/pkg/walkers/sql#Walk)) method that adds search to [squirrel](https://github.com/Masterminds/squirrel)'s `SelectBuilder` object.

##### graphviz

The `graphviz` package include a helper `graphviz.Walk` ([code](/pkg/walkers/graphviz/walk.go), [doc](https://pkg.go.dev/github.com/yaacov/tree-search-language/v5/pkg/walkers/graphviz#Walk)) method that exports `.dot` file nodes.

##### ident

The `ident` package include a helper `ident.Walk` ([code](/pkg/walkers/ident/walk.go), [doc](https://pkg.go.dev/github.com/yaacov/tree-search-language/v5/pkg/walkers/ident#Walk)) method that checks and mapps identifier names.

##### semantics

The `semantics` package include a helper `semantics.Walk` ([code](/pkg/walkers/semantics/walk.go), [doc](https://pkg.go.dev/github.com/yaacov/tree-search-language/v5/pkg/walkers/semantics#Walk)) method that reduce one data record to a bolean value (`true` or `false`) using a `tsl tree`.
