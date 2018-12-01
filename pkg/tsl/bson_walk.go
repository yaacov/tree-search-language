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

package tsl

import (
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Returns the identifier setring from an IdentOp operator node.
func identString(n interface{}) string {
	// This is an identifier.
	if str, ok := n.(Node).Left.(string); ok {
		return str
	}

	// This is not an identifier.
	return ""
}

// bsonFromArray helper method creates a slice of bson values from an interface,
// supported values can be strings or floats.
func bsonFromArray(a interface{}) (values []interface{}, err error) {
	for _, v := range a.([]Node) {
		// Check node value type.
		if s, ok := v.Left.(string); ok {
			// Node value is string.
			values = append(values, s)
		} else if f, ok := v.Left.(float64); ok {
			// Node value is float.
			values = append(values, f)
		} else {
			// Not a string or a float,
			// We do not support values other then strings or floats.
			err = fmt.Errorf("un supported value type of var: %v", v.Left)
			return
		}
	}

	// If here this is a valid array.
	return
}

// BSONWalk travel the TSL tree to create mongo-go-driver bson select operators.
//
// Users can call the BSONWalk method to filter a mongo Find.
//
//  // Prepare filter
//  filter, _ := BSONWalk(tree)
//
//  // Run query
//  cur, _ := collection.Find(ctx, bson.D(filter))
//
// mongo-go-driver: https://github.com/mongodb/mongo-go-driver
//
func BSONWalk(n Node) (b bson.E, err error) {
	var values []interface{}
	var l, r bson.E

	// Note: tsl function constants should be the same as mongo bson
	// functions (e.g. tsl.AndOp is "$and" in tsl and in mongo bson),
	// this is the reason we can use n.Func as a function name in the mongo
	// bson Element construction.
	switch n.Func {
	case OrOp, AndOp:
		l, err = BSONWalk(n.Left.(Node))
		if err != nil {
			return
		}
		r, err = BSONWalk(n.Right.(Node))
		if err != nil {
			return
		}
		b = bson.E{n.Func, bson.A{bson.D{l}, bson.D{r}}}
	case EqOp, NotEqOp, LtOp, LteOp, GtOp, GteOp:
		b = bson.E{identString(n.Left), bson.D{{n.Func, n.Right.(Node).Left}}}
	case InOp, NotInOp:
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.E{identString(n.Left), bson.D{{n.Func, values}}}
	case IsNilOp:
		b = bson.E{identString(n.Left), bson.D{{"$exists", false}}}
	case IsNotNilOp:
		b = bson.E{identString(n.Left), bson.D{{"$exists", true}}}
	case RegexOp:
		b = bson.E{identString(n.Left), primitive.Regex{n.Right.(Node).Left.(string), ""}}
	case NotRegexOp:
		b = bson.E{identString(n.Left),
			bson.D{{"$not", primitive.Regex{n.Right.(Node).Left.(string), ""}}}}
	case BetweenOp:
		// Mongo does not have a between function, translating sql's between into
		// two none eq operators.
		// Note: sql's between operator is inclusive: begin and end values are included.
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.E{"$and",
			bson.A{
				bson.D{{identString(n.Left), bson.D{{"$gte", values[0]}}}},
				bson.D{{identString(n.Left), bson.D{{"$lte", values[1]}}}},
			}}
	case NotBetweenOp:
		// Mongo does not have a between function, translating sql's not between into
		// two none eq operators.
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.E{"$or",
			bson.A{
				bson.D{{identString(n.Left), bson.D{{"$lt", values[0]}}}},
				bson.D{{identString(n.Left), bson.D{{"$gt", values[1]}}}},
			}}
	default:
		// If here than the operator is not supported.
		err = fmt.Errorf("un supported operand: %s", n.Func)
	}

	return
}
