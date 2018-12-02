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
	"fmt"

	"github.com/yaacov/tsl/pkg/tsl"
)

// Walk travel the TSL tree and replace identifiers.
//
// Users can call the Walk method to check and replace identifiers.
//
//   newTree, err = ident.Walk(tree, map[string]string)
//
func Walk(n tsl.Node, identMap map[string]string) (tsl.Node, error) {
	var err error

	// Walk tree.
	switch n.Func {
	case tsl.IdentOp:
		// If we have an identifier, check for it in the identMap.
		if v, ok := identMap[n.Left.(string)]; ok {
			// If valid identifier, use it.
			n.Left = v
			return n, nil
		}

		return n, fmt.Errorf("unknown identifier: %s", n.Left.(string))
	case tsl.StringOp, tsl.NumberOp:
		// This are our leafs.
		return n, nil
	default:
		// Check identifiers on left side.
		if n.Left != nil {
			n.Left, err = Walk(n.Left.(tsl.Node), identMap)
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
			if _, ok := n.Right.(tsl.Node); ok {
				// It's a node.
				n.Right, err = Walk(n.Right.(tsl.Node), identMap)
				if err != nil {
					return n, err
				}
			}
		}

		return n, nil
	}
}
