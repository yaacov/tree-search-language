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
	"fmt"

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
//	// the new tree will contain: `pages > 100 AND author = "Joe"`
//	newTree, identifiers, err = ident.Walk(tree, check)
//	defer newTree.Free()
func Walk(n *tsl.TSLNode, check func(s string) (string, error)) (*tsl.TSLNode, []string, error) {
	if n == nil {
		return nil, nil, nil
	}

	// Create a deep copy of the input tree to avoid mutations
	treeCopy := n.Clone()
	if treeCopy == nil {
		return nil, nil, fmt.Errorf("failed to clone input tree")
	}

	identifiers := make(map[string]bool)
	newTree, err := walkAndReplace(treeCopy, check, identifiers)
	if err != nil {
		treeCopy.Free() // Clean up the copy if there's an error
		return nil, nil, err
	}

	// Convert map to slice
	identList := make([]string, 0, len(identifiers))
	for ident := range identifiers {
		identList = append(identList, ident)
	}

	return newTree, identList, nil
}

// processIdentifier handles the common logic for processing identifier nodes
func processIdentifier(n *tsl.TSLNode, identifiers map[string]bool, check func(s string) (string, error)) (*tsl.TSLNode, error) {
	ident := n.Value().(string)
	identifiers[ident] = true

	newIdent, err := check(ident)
	if err != nil {
		return nil, err
	}

	return tsl.ParseTSL(newIdent)
}

func walkAndReplace(n *tsl.TSLNode, check func(s string) (string, error), identifiers map[string]bool) (*tsl.TSLNode, error) {
	if n == nil {
		return nil, nil
	}

	switch n.Type() {
	case tsl.KindBinaryExpr:
		op := n.Value().(tsl.TSLExpressionOp)

		// Process both sides of the binary expression
		left := op.Left
		if left.Type() == tsl.KindIdentifier {
			newNode, err := processIdentifier(left, identifiers, check)
			if err != nil {
				return nil, err
			}
			n.AttachLeft(newNode)
		} else {
			_, err := walkAndReplace(left, check, identifiers)
			if err != nil {
				return nil, err
			}
		}

		right := op.Right
		if right.Type() == tsl.KindIdentifier {
			newNode, err := processIdentifier(right, identifiers, check)
			if err != nil {
				return nil, err
			}
			n.AttachRight(newNode)
		} else {
			_, err := walkAndReplace(right, check, identifiers)
			if err != nil {
				return nil, err
			}
		}

		return n, nil

	case tsl.KindUnaryExpr:
		op := n.Value().(tsl.TSLExpressionOp)

		// Process the operand
		right := op.Right
		if right.Type() == tsl.KindIdentifier {
			newNode, err := processIdentifier(right, identifiers, check)
			if err != nil {
				return nil, err
			}
			n.AttachChild(newNode)
		} else {
			_, err := walkAndReplace(right, check, identifiers)
			if err != nil {
				return nil, err
			}
		}

		return n, nil

	case tsl.KindIdentifier:
		return processIdentifier(n, identifiers, check)

	default:
		return n, nil
	}
}
