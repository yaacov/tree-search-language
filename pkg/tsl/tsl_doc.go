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

// Package tsl describe and parse the Tree Search Language (TSL),
// TSL is a wonderful search language, With similar grammar to SQL's
// where part.
//
// TSL grammar examples:
//  "name = 'joe' or name = 'jane'"
//  "city in ('paris', 'rome', 'milan') and state != 'spane'"
//
// The package provide the ParseTSL method to convert TSL string into TSL tree.
//
// TSL tree can be use to create search engines on data sources, for
// example bsonWalk method converts TSL tree into bson for MongoDB based search:
// https://github.com/yaacov/tsl/blob/master/cmd/tsl_mongo/walk.go
package tsl
