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

// Package mongo helps to create mongo BSON filters using the TSL package.
package mongo

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/yaacov/tree-search-language/pkg/tsl"
)

// Returns the identifier setring from an IdentOp operator node.
func identString(n interface{}) string {
	// This is an identifier.
	if str, ok := n.(tsl.Node).Left.(string); ok {
		return str
	}

	// This is not an identifier.
	return ""
}

// bsonFromArray helper method creates a slice of bson values from an interface,
// supported values can be strings or floats.
func bsonFromArray(a interface{}) (values []interface{}, err error) {
	n := a.(tsl.Node)

	// Check that this is an array node.
	if n.Func != tsl.ArrayOp {
		err = tsl.UnexpectedLiteralError{Literal: n.Func}
		return
	}

	nodes := n.Right.([]tsl.Node)
	for _, v := range nodes {
		// Check node value type.
		switch l := v.Left.(type) {
		case string, float64:
			// Node value is string or float.
			values = append(values, l)
		default:
			// Not a string or a float,
			// We do not support values other then strings or floats.
			err = tsl.UnexpectedLiteralError{Literal: v.Left}
			return
		}
	}

	// If here this is a valid array.
	return
}

// Walk travel the TSL tree to create mongo-go-driver bson select operators.
//
// Users can call the Walk method to filter a mongo Find.
//
//  // Prepare filter
//  filter, _ := mongo.Walk(tree)
//
//  // Run query
//  cur, _ := collection.Find(ctx, filter)
//
// mongo-go-driver: https://github.com/mongodb/mongo-go-driver
//
func Walk(n tsl.Node) (b bson.D, err error) {
	var values []interface{}
	var l, r bson.D

	// Note: tsl function constants should be the same as mongo bson
	// functions (e.g. tsl.AndOp is "$and" in tsl and in mongo bson),
	// this is the reason we can use n.Func as a function name in the mongo
	// bson Element construction.
	switch n.Func {
	case tsl.OrOp, tsl.AndOp:
		l, err = Walk(n.Left.(tsl.Node))
		if err != nil {
			return
		}
		r, err = Walk(n.Right.(tsl.Node))
		if err != nil {
			return
		}
		b = bson.D{{n.Func, bson.A{l, r}}}
	case tsl.EqOp, tsl.NotEqOp, tsl.LtOp, tsl.LteOp, tsl.GtOp, tsl.GteOp:
		b = bson.D{{identString(n.Left), bson.D{{n.Func, n.Right.(tsl.Node).Left}}}}
	case tsl.InOp, tsl.NotInOp:
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.D{{identString(n.Left), bson.D{{n.Func, values}}}}
	case tsl.IsNilOp:
		b = bson.D{{identString(n.Left), bson.D{{"$exists", false}}}}
	case tsl.IsNotNilOp:
		b = bson.D{{identString(n.Left), bson.D{{"$exists", true}}}}
	case tsl.RegexOp:
		b = bson.D{{identString(n.Left), primitive.Regex{n.Right.(tsl.Node).Left.(string), ""}}}
	case tsl.NotRegexOp:
		b = bson.D{{identString(n.Left),
			bson.D{{"$not", primitive.Regex{n.Right.(tsl.Node).Left.(string), ""}}}}}
	case tsl.BetweenOp:
		// Mongo does not have a between function, translating sql's between into
		// two none eq operators.
		// Note: sql's between operator is inclusive: begin and end values are included.
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.D{{"$and",
			bson.A{
				bson.D{{identString(n.Left), bson.D{{"$gte", values[0]}}}},
				bson.D{{identString(n.Left), bson.D{{"$lte", values[1]}}}},
			}}}
	case tsl.NotBetweenOp:
		// Mongo does not have a between function, translating sql's not between into
		// two none eq operators.
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.D{{"$or",
			bson.A{
				bson.D{{identString(n.Left), bson.D{{"$lt", values[0]}}}},
				bson.D{{identString(n.Left), bson.D{{"$gt", values[1]}}}},
			}}}
	default:
		// If here than the operator is not supported.
		err = tsl.UnexpectedLiteralError{Literal: n.Func}
	}

	return
}
