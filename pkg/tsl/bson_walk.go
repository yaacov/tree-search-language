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
)

// Returns the identifier setring from an IdentOp operator node.
func identString(n interface{}) string {
	// This is just a string, not and IdentOp.
	if str, ok := n.(Node).Left.(string); ok {
		return str
	}

	// This is an IdentOp.
	return n.(Node).Left.(Node).Left.(string)
}

// bsonFromArray helper method creates a slice of bson values from an interface,
// supported values can be strings or floats.
func bsonFromArray(a interface{}) (values []*bson.Value, err error) {
	for _, v := range a.([]interface{}) {
		if w, ok := v.(string); ok {
			// Check for a string value.
			values = append(values, bson.VC.String(w))
		} else if w, ok := v.(float64); ok {
			// Check for a float value.
			values = append(values, bson.VC.Double(w))
		} else {
			// Not a string or a float,
			// We do not support values other then strings or floats.
			err = fmt.Errorf("un supported value type of var: %v", v)
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
//  cur, _ := collection.Find(ctx, bson.NewDocument(filter))
//
// mongo-go-driver: https://github.com/mongodb/mongo-go-driver
//
func BSONWalk(n Node) (b *bson.Element, err error) {
	var values []*bson.Value
	var l, r *bson.Element

	// Note: tsl function constants should be the same as mongo bson
	// functions (e.g. tsl.AndOp is "$and" in tsl and in mongo bson),
	// this is the reason we can use n.Func as a function name in the mongo
	// bson Element construction.
	switch n.Func {
	case AndOp, OrOp:
		l, err = BSONWalk(n.Left.(Node))
		if err != nil {
			return
		}
		r, err = BSONWalk(n.Right.(Node))
		if err != nil {
			return
		}
		b = bson.EC.ArrayFromElements(n.Func,
			bson.VC.DocumentFromElements(l),
			bson.VC.DocumentFromElements(r),
		)
	case EqOp, NotEqOp, LtOp, LteOp, GtOp, GteOp:
		b = bson.EC.SubDocumentFromElements(identString(n.Left),
			bson.EC.Interface(n.Func, n.Right),
		)
	case InOp:
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.EC.SubDocumentFromElements(identString(n.Left),
			bson.EC.ArrayFromElements("$in", values...),
		)
	case NotInOp:
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.EC.SubDocumentFromElements(identString(n.Left),
			bson.EC.SubDocumentFromElements("$not",
				bson.EC.ArrayFromElements("$in", values...),
			),
		)
	case IsNilOp:
		b = bson.EC.SubDocumentFromElements(identString(n.Left), bson.EC.Boolean("$exists", false))
	case IsNotNilOp:
		b = bson.EC.SubDocumentFromElements(identString(n.Left), bson.EC.Boolean("$exists", true))
	case RegexOp:
		b = bson.EC.Regex(identString(n.Left), n.Right.(string), "")
	case NotRegexOp:
		b = bson.EC.SubDocumentFromElements(identString(n.Left),
			bson.EC.Regex("$not", n.Right.(string), ""),
		)
	case BetweenOp:
		// Mongo does not have a between function, translating sql's between into
		// two none eq operators.
		// Note: sql's between operator is inclusive: begin and end values are included.
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.EC.ArrayFromElements("$and",
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(identString(n.Left), bson.EC.Interface("$gte", values[0])),
			),
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(identString(n.Left), bson.EC.Interface("$lte", values[1])),
			),
		)
	case NotBetweenOp:
		// Mongo does not have a between function, translating sql's not between into
		// two none eq operators.
		values, err = bsonFromArray(n.Right)
		if err != nil {
			return
		}
		b = bson.EC.ArrayFromElements("$or",
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(identString(n.Left), bson.EC.Interface("$lt", values[0])),
			),
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(identString(n.Left), bson.EC.Interface("$gt", values[1])),
			),
		)
	default:
		// If here than the operator is not supported.
		err = fmt.Errorf("un supported operand: %s", n.Func)
	}

	return
}
