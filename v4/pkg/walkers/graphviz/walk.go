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

	"github.com/yaacov/tree-search-language/v4/pkg/tsl"
)

// Character poll for random string genrator.
const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Default styles for .dot file output.
const identStyle = "shape=record color=red"
const numberStyle = "shape=record color=blue"
const stringStyle = "shape=record color=blue"
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
//   s, err = graphviz.Walk("", tree, "")
//
// The return string will be a node list for a .dot file:
//
//   tcuA [shape=box color=black label="$eq"]
//   xhxK [shape=record color=red label="$ident | 'city'" ]
//   QFDa [shape=record color=blue label="$string | 'rome'" ]
//   tcuA -> { xhxK, QFDa }
//
// For valid Graphviz dot file, the nodes must be wrapped in a `digraph` object:
//
//   s, err = graphviz.Walk("", tree, "")
//   s = fmt.Sprintf("digraph {\n%s\n}\n", s)
func Walk(in string, n tsl.Node, nodeID string) (out string, err error) {
	// If node ID is missing, create one.
	if nodeID == "" {
		nodeID = randStr(4)
	}

	// If we have in string, add a new line break.
	if in != "" {
		in = fmt.Sprintf("%s\n", in)
	}

	switch n.Func {
	case tsl.IdentOp:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | '%s'\" ]",
			nodeID,
			identStyle,
			n.Func,
			n.Left)
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case tsl.StringOp:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | '%s'\" ]",
			nodeID,
			stringStyle,
			n.Func,
			n.Left)
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case tsl.NumberOp:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | %g\" ]",
			nodeID,
			numberStyle,
			n.Func,
			n.Left)
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	default:
		// Add node label.
		st := fmt.Sprintf("%s [%s label=\"%s\"]",
			nodeID,
			opStyle,
			n.Func)
		childrens := []string{}

		// Add left child.
		if n.Left != nil {
			leftID := randStr(4)
			childrens = append(childrens, leftID)

			l, err := Walk(in, n.Left.(tsl.Node), leftID)
			if err != nil {
				return "", err
			}
			st = fmt.Sprintf("%s\n%s", st, l)
		}

		// Add right child.
		if n.Right != nil {
			var nn []tsl.Node

			// Check if right hand arg is a node, or an array of nodes.
			if n.Func == tsl.ArrayOp {
				nn = n.Right.([]tsl.Node)
			} else {
				nn = []tsl.Node{n.Right.(tsl.Node)}
			}

			// Add all nodes to graph.
			for _, v := range nn {
				rightID := randStr(4)
				childrens = append(childrens, rightID)

				r, err := Walk(in, v, rightID)
				if err != nil {
					return "", err
				}
				st = fmt.Sprintf("%s\n%s", st, r)
			}
		}

		// Add parrent node.
		in = fmt.Sprintf("%s%s\n%s -> { %s }",
			in,
			st,
			nodeID,
			strings.Join(childrens, ", "))

		return in, nil
	}

	return
}
