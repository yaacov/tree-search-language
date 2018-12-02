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

// Package walkers examples TSL tree walking, and helps to generate SQL and MongoDB
// query filters. sql.Walk and mongo.Walk methods can be used to create such filters.
//
// Squirrel walk code:  https://github.com/yaacov/tsl/blob/master/pkg/walkers/sql/walk.go
//
// SQL generation example:  https://github.com/yaacov/tsl/blob/master/cmd/tsl_sqlite/main.go
//
// Usage:
//   filter, err := sql.Walk(tree)
//
//   sql, args, err := sq.Select("name, city, state").
//        From("users").
//        Where(filter).
//        ToSql()
//
// BSON walk code: https://github.com/yaacov/tsl/blob/master/pkg/walkers/mongo/walk.go
//
// BSON generation example:  https://github.com/yaacov/tsl/blob/master/cmd/tsl_mongo/main.go
//
// Usage:
//   // Prepare a bson filter
//   filter, err = mongo.Walk(tree)
//
//   // Run query
//   cur, err := collection.Find(ctx, filter)
//
// squirrel: https://github.com/Masterminds/squirrel
//
// mongo-go-driver: https://github.com/mongodb/mongo-go-driver
package walkers
