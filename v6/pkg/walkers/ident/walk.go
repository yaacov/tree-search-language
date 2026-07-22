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
//		return nil, fmt.Errorf("column not found")
//	}
//
//	// Check and replace user identifiers with the SQL table column names,
//	// and return the new tree.
//	//
//	// If the input tree contains: `spec.pages > 100 AND author = "Joe"`,
//	// the new tree will contain: `pages > 100 AND author = "Joe"`
//	newTree, err = ident.Walk(tree, check)
func Walk(n *tsl.TSLNode, check func(s string) (string, error)) (*tsl.TSLNode, error) {
	if n == nil {
		return nil, nil
	}

	// Create a deep copy of the input tree to avoid mutations
	treeCopy := n.Clone()
	if treeCopy == nil {
		return nil, fmt.Errorf("failed to clone input tree")
	}

	newTree, err := walkAndReplace(treeCopy, check)
	if err != nil {
		return nil, err
	}

	return newTree, nil
}

// processIdentifier handles the common logic for processing identifier nodes
func processIdentifier(n *tsl.TSLNode, check func(s string) (string, error)) (*tsl.TSLNode, error) {
	ident := n.Value().(string)

	newIdent, err := check(ident)
	if err != nil {
		return nil, err
	}

	return tsl.ParseTSL(newIdent)
}

func walkAndReplace(n *tsl.TSLNode, check func(s string) (string, error)) (*tsl.TSLNode, error) {
	if n == nil {
		return nil, nil
	}

	switch n.Type() {
	case tsl.KindBinaryExpr:
		op := n.Value().(tsl.TSLExpressionOp)

		// Process both sides of the binary expression
		var processedLeft, processedRight *tsl.TSLNode
		var err error

		if op.Left.Type() == tsl.KindIdentifier {
			processedLeft, err = processIdentifier(op.Left, check)
			if err != nil {
				return nil, err
			}
		} else {
			processedLeft, err = walkAndReplace(op.Left, check)
			if err != nil {
				return nil, err
			}
		}

		if op.Right.Type() == tsl.KindIdentifier {
			processedRight, err = processIdentifier(op.Right, check)
			if err != nil {
				return nil, err
			}
		} else {
			processedRight, err = walkAndReplace(op.Right, check)
			if err != nil {
				return nil, err
			}
		}

		// Update the binary expression with the processed children
		n.SetLeft(processedLeft)
		n.SetRight(processedRight)
		return n, nil

	case tsl.KindUnaryExpr:
		op := n.Value().(tsl.TSLExpressionOp)

		// Process the operand
		var processedRight *tsl.TSLNode
		var err error

		if op.Right.Type() == tsl.KindIdentifier {
			processedRight, err = processIdentifier(op.Right, check)
			if err != nil {
				return nil, err
			}
		} else {
			processedRight, err = walkAndReplace(op.Right, check)
			if err != nil {
				return nil, err
			}
		}

		// Update the unary expression with the processed child
		n.SetRight(processedRight)
		return n, nil

	case tsl.KindIdentifier:
		return processIdentifier(n, check)

	case tsl.KindArrayLiteral:
		arr := n.Value().(tsl.TSLArrayLiteral)
		newValues := make([]*tsl.TSLNode, len(arr.Values))
		for i, child := range arr.Values {
			processed, err := walkAndReplace(child, check)
			if err != nil {
				return nil, err
			}
			newValues[i] = processed
		}
		n.SetArrayValues(newValues)
		return n, nil

	default:
		return n, nil
	}
}
