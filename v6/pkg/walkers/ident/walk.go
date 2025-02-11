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

// Walk traverses a TSL tree and returns a slice of unique identifiers found in the tree.
func Walk(n *tsl.TSLNode) ([]string, error) {
	identifiers := make(map[string]bool)
	err := walk(n, identifiers)
	if err != nil {
		return nil, err
	}

	// Convert map keys to slice
	result := make([]string, 0, len(identifiers))
	for ident := range identifiers {
		result = append(result, ident)
	}

	return result, nil
}

// walk is a helper function that recursively traverses the tree
func walk(n *tsl.TSLNode, identifiers map[string]bool) error {
	if n == nil {
		return nil
	}

	switch n.Type() {
	case tsl.KindIdentifier:
		// Add identifier to the map
		identifiers[n.Value().(string)] = true
	case tsl.KindBinaryExpr:
		// Process binary expression
		op := n.Value().(tsl.TSLExpressionOp)
		if err := walk(op.Left, identifiers); err != nil {
			return err
		}
		if err := walk(op.Right, identifiers); err != nil {
			return err
		}
	case tsl.KindUnaryExpr:
		// Process unary expression
		op := n.Value().(tsl.TSLExpressionOp)
		if err := walk(op.Left, identifiers); err != nil {
			return err
		}
	}

	return nil
}
