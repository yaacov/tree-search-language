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

// Package graphviz helps to create graphviz dot files using the TSL package.
package graphviz

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

// Character poll for random string genrator.
const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Default styles for .dot file output.
const identStyle = "shape=record color=red"
const numberStyle = "shape=record color=blue"
const stringStyle = "shape=record color=blue"
const dateStyle = "shape=record color=orange"
const opStyle = "shape=box color=black"

// Generate a random string.
func randStr(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

// Walk travel the TSL tree to create Graphviz dot nodes.
//
// Users can call the Walk method to get the dot file graph string.
//
//	s, err = graphviz.Walk("", tree, "")
//
// The return string will be a node list for a .dot file:
//
//	tcuA [shape=box color=black label="$eq"]
//	xhxK [shape=record color=red label="$ident | 'city'" ]
//	QFDa [shape=record color=blue label="$string | 'rome'" ]
//	tcuA -> { xhxK, QFDa }
//
// For valid Graphviz dot file, the nodes must be wrapped in a `digraph` object:
//
//	s, err = graphviz.Walk("", tree, "")
//	s = fmt.Sprintf("digraph {\n%s\n}\n", s)
func Walk(in string, n *tsl.TSLNode, nodeID string) (out string, err error) {
	// If node ID is missing, create one.
	if nodeID == "" {
		nodeID = randStr(4)
	}

	// If we have in string, add a new line break.
	if in != "" {
		in = fmt.Sprintf("%s\n", in)
	}

	switch n.Type() {
	case tsl.KindIdentifier:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | '%s'\" ]",
			nodeID,
			identStyle,
			n.Type().String(),
			n.Value())
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case tsl.KindStringLiteral:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | '%s'\" ]",
			nodeID,
			stringStyle,
			n.Type().String(),
			n.Value())
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case tsl.KindNumericLiteral:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | %g\" ]",
			nodeID,
			numberStyle,
			n.Type().String(),
			n.Value())
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case tsl.KindBooleanLiteral, tsl.KindDateLiteral, tsl.KindTimestampLiteral:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | %v\" ]",
			nodeID,
			dateStyle,
			n.Type().String(),
			n.Value())
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case tsl.KindBinaryExpr, tsl.KindUnaryExpr:
		expr := n.Value().(tsl.TSLExpressionOp)
		// Add node label.
		st := fmt.Sprintf("%s [%s label=\"%s\"]",
			nodeID,
			opStyle,
			expr.Operator.String())
		childrens := []string{}

		// Add left child for binary expressions
		if expr.Left != nil {
			leftID := randStr(4)
			childrens = append(childrens, leftID)

			l, err := Walk(in, expr.Left, leftID)
			if err != nil {
				return "", err
			}
			st = fmt.Sprintf("%s\n%s", st, l)
		}

		// Add right child
		if expr.Right != nil {
			rightID := randStr(4)
			childrens = append(childrens, rightID)

			r, err := Walk(in, expr.Right, rightID)
			if err != nil {
				return "", err
			}
			st = fmt.Sprintf("%s\n%s", st, r)
		}

		// Add parent node
		in = fmt.Sprintf("%s%s\n%s -> { %s }",
			in,
			st,
			nodeID,
			strings.Join(childrens, ", "))

		return in, nil
	case tsl.KindArrayLiteral:
		array := n.Value().(tsl.TSLArrayLiteral)
		// Add node label
		st := fmt.Sprintf("%s [%s label=\"ARRAY\"]",
			nodeID,
			opStyle)
		childrens := []string{}

		// Add all array elements
		for _, elem := range array.Values {
			elemID := randStr(4)
			childrens = append(childrens, elemID)

			e, err := Walk(in, elem, elemID)
			if err != nil {
				return "", err
			}
			st = fmt.Sprintf("%s\n%s", st, e)
		}

		// Add parent node
		in = fmt.Sprintf("%s%s\n%s -> { %s }",
			in,
			st,
			nodeID,
			strings.Join(childrens, ", "))

		return in, nil
	}

	return
}
