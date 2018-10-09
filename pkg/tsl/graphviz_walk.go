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

package tsl

import (
	"fmt"
	"math/rand"
	"strings"
)

var pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var identStyle = "shape=record color=red"
var numberStyle = "shape=record color=blue"
var stringStyle = "shape=record color=blue"
var opStyle = "shape=box color=black"

// Generate a random string.
func randStr(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

// GraphvizWalk travel the TSL tree to create Graphviz dot nodes.
//
// Users can call the GraphvizWalk method to get the dot file graph string.
//
//   filter, _ := SquirrelWalk(tree)
//   s, err = tsl.GraphvizWalk("", tree, "")
//
// The return string will be a node list for a .dot file:
//
//   XVlB [shape=circle label=<<b>$eq</b>>]
//   zgba [shape=box label=<<b>$ident</b>: "city">]
//   iCMR [shape=box label=<<b>$string</b>: "rome">]
//   XVlB -> { zgba, iCMR }
//
func GraphvizWalk(in string, n Node, nodeID string) (out string, err error) {
	// If node ID is missing, create one.
	if nodeID == "" {
		nodeID = randStr(4)
	}

	// If we have in string, add a new line break.
	if in != "" {
		in = fmt.Sprintf("%s\n", in)
	}

	switch n.Func {
	case IdentOp:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | '%s'\" ]",
			nodeID,
			identStyle,
			n.Func,
			n.Left)
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case StringOp:
		// Add leaf label and value.
		nodeLabel := fmt.Sprintf("%s [%s label=\"%s | '%s'\" ]",
			nodeID,
			stringStyle,
			n.Func,
			n.Left)
		out = fmt.Sprintf("%s%s", in, nodeLabel)
	case NumberOp:
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

			l, err := GraphvizWalk(in, n.Left.(Node), leftID)
			if err != nil {
				return "", err
			}
			st = fmt.Sprintf("%s\n%s", st, l)
		}

		// Add right child.
		if n.Right != nil {
			var nn []Node

			// Check if right hand arg is a node, or an array of nodes.
			if _, ok := n.Right.(Node); ok {
				nn = []Node{n.Right.(Node)}
			} else {
				nn = n.Right.([]Node)
			}

			// Add all nodes to graph.
			for _, v := range nn {
				rightID := randStr(4)
				childrens = append(childrens, rightID)

				r, err := GraphvizWalk(in, v, rightID)
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
