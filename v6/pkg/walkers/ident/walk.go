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

// Package ident helps extract identifiers from TSL trees.
package ident

import (
	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

// Walk traverses the TSL tree and replaces identifiers using the check function.
//
// Users can call the Walk method to check and replace identifiers.
// The function returns the modified tree, a list of all identifiers found in the tree,
// and any error encountered.
//
// Example:
//
//	columnNamesMap :=  map[string]string{
//		"title":       "title",
//		"author":      "author",
//		"spec.pages":  "pages",
//		"spec.rating": "rating",
//	}
//
//	func check(s string) (string, error) {
//		// Check for column name in map.
//		if v, ok := columnNamesMap[s]; ok {
//			return v, nil
//		}
//
//		// If not found return string as is, and an error.
//		return s, fmt.Errorf("column not found")
//	}
//
//	// Check and replace user identifiers with the SQL table column names,
//	// and get a list of all identifiers used in the tree.
//	//
//	// If the input tree contains: `spec.pages > 100 AND author = "Joe"`,
//	// the identifiers list will contain: ["spec.pages", "author"]
//	//
//	newTree, identifiers, err = ident.Walk(tree, check)
func Walk(n *tsl.TSLNode, check func(s string) (string, error)) (*tsl.TSLNode, []string, error) {
	if n == nil {
		return nil, nil, nil
	}

	identifiers := make(map[string]bool)
	newTree, err := walkAndReplace(n, check, identifiers)
	if err != nil {
		return nil, nil, err
	}

	// Convert map to slice
	identList := make([]string, 0, len(identifiers))
	for ident := range identifiers {
		identList = append(identList, ident)
	}

	return newTree, identList, nil
}

// walkAndReplace recursively traverses the tree, replacing identifiers
func walkAndReplace(n *tsl.TSLNode, check func(s string) (string, error), identifiers map[string]bool) (*tsl.TSLNode, error) {
	if n == nil {
		return nil, nil
	}

	switch n.Type() {
	case tsl.KindIdentifier:
		ident := n.Value().(string)
		// Add to the map of found identifiers
		identifiers[ident] = true

		// Check and possibly replace the identifier
		newIdent, err := check(ident)
		if err != nil {
			return nil, err
		}

		// Parse the new identifier as a TSL expression
		newNode, err := tsl.ParseTSL(newIdent)
		if err != nil {
			return nil, err
		}
		return newNode, nil

	case tsl.KindBinaryExpr:
		op := n.Value().(tsl.TSLExpressionOp)

		// Process both sides of the binary expression
		left, err := walkAndReplace(op.Left, check, identifiers)
		if err != nil {
			return nil, err
		}
		op.Left = left

		right, err := walkAndReplace(op.Right, check, identifiers)
		if err != nil {
			return nil, err
		}
		op.Right = right

		return n, nil

	case tsl.KindUnaryExpr:
		op := n.Value().(tsl.TSLExpressionOp)

		// Process the operand
		right, err := walkAndReplace(op.Right, check, identifiers)
		if err != nil {
			return nil, err
		}
		op.Right = right

		return n, nil

	default:
		return n, nil
	}
}
