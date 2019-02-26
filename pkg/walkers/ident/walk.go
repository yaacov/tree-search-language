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

// Package ident helps to replace identifiers in a TSL tree.
package ident

import (
	"github.com/yaacov/tree-search-language/pkg/tsl"
)

// Walk travel the TSL tree and replace identifiers.
//
// Users can call the Walk method to check and replace identifiers.
//
// Example:
//  	columnNamesMap :=  map[string]string{
//  		"title":       "title",
//  		"author":      "author",
//  		"spec.pages":  "pages",
//  		"spec.rating": "rating",
//  	}
//
//  	func check(s string) (string, error) {
//  		// Chekc for column name in map.
//  		if v, ok := columnNamesMap[s]; ok {
//  			return v, nil
//  		}
//
//  		// If not found return string as is, and an error.
//  		return s, fmt.Errorf("column not found")
//  	}
//
//  	// Check and replace user identifiers with the SQL table column names.
//  	//
//  	// SQL table columns are "title, author, pages and rating", but for
//  	// users pages and ratings are fields of an internal struct called
//  	// spec (e.g. {"title": "Book", "author": "Joe", "spec": {"pages": 5, "rating": 5}}).
//  	//
//  	newTree, err = ident.Walk(tree, check)
//
func Walk(n tsl.Node, checkColumnName func(string) (string, error)) (tsl.Node, error) {
	var err error
	var v string

	// Walk tree.
	switch n.Func {
	case tsl.IdentOp:
		// If we have an identifier, check for it in the identMap.
		if v, err = checkColumnName(n.Left.(string)); err == nil {
			// If valid identifier, use it.
			n.Left = v
			return n, nil
		}

		return n, err
	case tsl.StringOp, tsl.NumberOp:
		// This are our leafs.
		return n, nil
	default:
		// Check identifiers on left side.
		if n.Left != nil {
			n.Left, err = Walk(n.Left.(tsl.Node), checkColumnName)
			if err != nil {
				return n, err
			}
		}

		// Check identifiers on right side.
		if n.Right != nil {
			// Check if right hand arg is a node, or an array of nodes.
			//
			// If it's an array of nodes.
			// We assume that all are leafs, no nead to walk on them.
			if n.Func != tsl.ArrayOp {
				// It's a node.
				n.Right, err = Walk(n.Right.(tsl.Node), checkColumnName)
				if err != nil {
					return n, err
				}
			}
		}

		return n, nil
	}
}
