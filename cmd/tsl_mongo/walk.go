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

// Package main
package main

import (
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/yaacov/tsl/pkg/tsl"
)

func bsonFromArray(a interface{}) (values []*bson.Value) {
	for _, v := range a.([]interface{}) {
		if w, ok := v.(string); ok {
			values = append(values, bson.VC.String(w))
		}
		if w, ok := v.(float64); ok {
			values = append(values, bson.VC.Double(w))
		}
	}

	return
}

func bsonWalk(n tsl.Node) *bson.Element {
	switch n.Func {
	case tsl.AndOp, tsl.OrOp:
		return bson.EC.ArrayFromElements(n.Func,
			bson.VC.DocumentFromElements(bsonWalk(n.Left.(tsl.Node))),
			bson.VC.DocumentFromElements(bsonWalk(n.Right.(tsl.Node))),
		)
	case tsl.EqOp, tsl.NotEqOp, tsl.LtOp, tsl.LteOp, tsl.GtOp, tsl.GteOp:
		return bson.EC.SubDocumentFromElements(n.Left.(string), bson.EC.Interface(n.Func, n.Right))
	case tsl.InOp:
		return bson.EC.SubDocumentFromElements(n.Left.(string),
			bson.EC.ArrayFromElements("$in",
				bsonFromArray(n.Right)...,
			),
		)
	case tsl.NotInOp:
		return bson.EC.SubDocumentFromElements(n.Left.(string),
			bson.EC.SubDocumentFromElements("$not",
				bson.EC.ArrayFromElements("$in",
					bsonFromArray(n.Right)...,
				),
			),
		)
	case tsl.IsNilOp:
		return bson.EC.SubDocumentFromElements(n.Left.(string), bson.EC.Boolean("$exists", false))
	case tsl.IsNotNilOp:
		return bson.EC.SubDocumentFromElements(n.Left.(string), bson.EC.Boolean("$exists", true))
	case tsl.RegexOp:
		return bson.EC.Regex(n.Left.(string), n.Right.(string), "")
	case tsl.NotRegexOp:
		return bson.EC.SubDocumentFromElements(n.Left.(string),
			bson.EC.Regex("$not", n.Right.(string), ""),
		)
	case tsl.BetweenOp:
		right := bsonFromArray(n.Right)
		return bson.EC.ArrayFromElements("$and",
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(n.Left.(string), bson.EC.Interface("$gte", right[0])),
			),
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(n.Left.(string), bson.EC.Interface("$lte", right[1])),
			),
		)
	case tsl.NotBetweenOp:
		right := bsonFromArray(n.Right)
		return bson.EC.ArrayFromElements("$or",
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(n.Left.(string), bson.EC.Interface("$lt", right[0])),
			),
			bson.VC.DocumentFromElements(
				bson.EC.SubDocumentFromElements(n.Left.(string), bson.EC.Interface("$gt", right[1])),
			),
		)
	default:
		// We do not suppoert any other function code
		log.Fatal(fmt.Errorf("%s is not supported", n.Func))
	}

	return bson.EC.String("", "")
}
