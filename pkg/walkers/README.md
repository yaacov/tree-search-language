

<p align="center">
  <img src="https://raw.githubusercontent.com/yaacov/tsl/master/img/search-248.png" alt="TSL Logo">
</p>

# Tree Search Language (TSL) Imperial Walkers (IW)

##### sql

The `sql` package include a helper sql.Walk ([code](/pkg/walkers/sql/walk.go), [doc](https://godoc.org/github.com/yaacov/tsl/pkg/walkers/sql#Walk)) method that adds search to [squirrel](https://github.com/Masterminds/squirrel)'s SelectBuilder object.

##### mongo

The `mongo` package include a helper mongo.Walk ([code](/pkg/walkers/mongo/walk.go), [doc](https://godoc.org/github.com/yaacov/tsl/pkg/walkers/mongo#Walk)) method that adds search bson filter to [mongo-go-driver](https://github.com/mongodb/mongo-go-driver).

##### graphviz

The `graphviz` package include a helper graphviz.Walk ([code](/pkg/walkers/graphviz/walk.go), [doc](https://godoc.org/github.com/yaacov/tsl/pkg/walkers/graphviz#Walk)) method that exports `.dot` file nodes.

##### ident

The 'ident' package include a helper ident.Walk ([code](/pkg/walkers/ident/walk.go), [doc](https://godoc.org/github.com/yaacov/tsl/pkg/walkers/ident#Walk)) method that checks and mapps identifier names.
