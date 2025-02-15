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

// Base styles for shared properties
const baseRecordStyle = "shape=record"
const baseBoxStyle = "shape=box"

// Node specific styles combining base styles with colors
const identStyle = baseRecordStyle + " color=red"
const numberStyle = baseRecordStyle + " color=blue"
const stringStyle = baseRecordStyle + " color=blue"
const booleanStyle = baseRecordStyle + " color=purple"
const dateStyle = baseRecordStyle + " color=orange"
const timestampStyle = baseRecordStyle + " color=orange"
const opStyle = baseBoxStyle + " color=black"
const arrayStyle = baseBoxStyle + " color=green"

// Generate a random string of specified length using only letters
func randStr(l int) string {
	const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

// formatLeafNodeWithInput creates a graphviz node string for leaf nodes including input string
func formatLeafNodeWithInput(in string, nodeID string, style string, nodeType tsl.Kind, value interface{}) string {
	if in != "" {
		in = fmt.Sprintf("%s\n", in)
	}
	nodeStr := fmt.Sprintf("%s [%s label=\"%s | %v\" ]",
		nodeID,
		style,
		nodeType.String(),
		value)
	return fmt.Sprintf("%s%s", in, nodeStr)
}

// formatOperatorNode creates a graphviz node string for operator nodes
func formatOperatorNode(nodeID string, operator string) string {
	return fmt.Sprintf("%s [%s label=\"%s\"]",
		nodeID,
		opStyle,
		operator)
}

// handleChildren processes child nodes and returns their formatted string and IDs
func handleChildren(in string, children []*tsl.TSLNode) (string, []string, error) {
	childrenIDs := []string{}
	result := ""

	for _, child := range children {
		childID := randStr(4)
		childrenIDs = append(childrenIDs, childID)

		r, err := Walk(in, child, childID)
		if err != nil {
			return "", nil, err
		}
		result = fmt.Sprintf("%s\n%s", result, r)
	}

	return result, childrenIDs, nil
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
	if nodeID == "" {
		nodeID = randStr(4)
	}

	switch n.Type() {
	case tsl.KindIdentifier:
		out = formatLeafNodeWithInput(in, nodeID, identStyle, n.Type(), fmt.Sprintf("'%s'", n.Value()))
	case tsl.KindStringLiteral:
		out = formatLeafNodeWithInput(in, nodeID, stringStyle, n.Type(), fmt.Sprintf("'%s'", n.Value()))
	case tsl.KindNumericLiteral:
		out = formatLeafNodeWithInput(in, nodeID, numberStyle, n.Type(), n.Value())
	case tsl.KindBooleanLiteral:
		out = formatLeafNodeWithInput(in, nodeID, booleanStyle, n.Type(), n.Value())
	case tsl.KindDateLiteral:
		out = formatLeafNodeWithInput(in, nodeID, dateStyle, n.Type(), n.Value())
	case tsl.KindTimestampLiteral:
		out = formatLeafNodeWithInput(in, nodeID, timestampStyle, n.Type(), n.Value())
	case tsl.KindBinaryExpr:
		expr := n.Value().(tsl.TSLExpressionOp)
		st := formatOperatorNode(nodeID, expr.Operator.String())

		childrenStr, childrenIDs, err := handleChildren(in, []*tsl.TSLNode{expr.Left, expr.Right})
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s%s%s\n%s -> { %s }", in, st, childrenStr, nodeID, strings.Join(childrenIDs, ", ")), nil

	case tsl.KindUnaryExpr:
		expr := n.Value().(tsl.TSLExpressionOp)
		st := formatOperatorNode(nodeID, expr.Operator.String())

		childrenStr, childrenIDs, err := handleChildren(in, []*tsl.TSLNode{expr.Right})
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s%s%s\n%s -> { %s }", in, st, childrenStr, nodeID, strings.Join(childrenIDs, ", ")), nil

	case tsl.KindArrayLiteral:
		array := n.Value().(tsl.TSLArrayLiteral)
		st := formatOperatorNode(nodeID, n.Type().String())
		// Use arrayStyle instead of opStyle for arrays
		st = strings.Replace(st, opStyle, arrayStyle, 1)

		childrenStr, childrenIDs, err := handleChildren(in, array.Values)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s%s%s\n%s -> { %s }", in, st, childrenStr, nodeID, strings.Join(childrenIDs, ", ")), nil
	}

	return
}
