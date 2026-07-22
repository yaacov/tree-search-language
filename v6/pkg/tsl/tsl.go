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
	"github.com/yaacov/tree-search-language/v6/pkg/parser"
)

// TSLNode represents a tree search language AST (Abstract Syntax Tree) node.
// This provides the public interface for TSL parsing results.
// Use the accessor methods (Type, Value, Clone) to inspect the tree,
// and SetLeft/SetRight to modify expression children.
type TSLNode struct {
	node *Node
}

// TSLExpressionOp represents both unary and binary operations
type TSLExpressionOp struct {
	Operator    Operator
	Left, Right *TSLNode
}

// TSLArrayLiteral represents an array of nodes
type TSLArrayLiteral struct {
	Values []*TSLNode
}

// ParseTSL parses a TSL expression and returns the AST root node
func ParseTSL(input string) (*TSLNode, error) {
	parserNode, err := parser.Parse(input)
	if err != nil {
		// Return a TSL-specific error with position information
		if parseErr, ok := err.(*parser.ParseError); ok {
			return nil, &SyntaxError{
				Message:  parseErr.Message,
				Position: parseErr.Position,
				Context:  "",
				Input:    input,
			}
		}
		return nil, err
	}

	// Create TSL node from parsed input
	tslNode := wrapParserNode(parserNode)
	return &TSLNode{node: tslNode}, nil
}

// Clone creates a deep copy of the TSLNode and its children
func (n *TSLNode) Clone() *TSLNode {
	if n == nil || n.node == nil {
		return nil
	}
	return &TSLNode{node: n.node.Clone()}
}

// Type returns the type of the node
func (n *TSLNode) Type() Kind {
	if n == nil || n.node == nil {
		return Kind(-1)
	}
	return n.node.Kind
}

// Value returns the node's value based on its type
func (n *TSLNode) Value() interface{} {
	if n == nil || n.node == nil {
		return nil
	}

	switch n.node.Kind {
	case KindBooleanLiteral, KindNumericLiteral, KindStringLiteral,
		KindIdentifier, KindDateLiteral, KindTimestampLiteral:
		return n.node.Value
	case KindBinaryExpr:
		var left, right *TSLNode
		if n.node.Left != nil {
			left = &TSLNode{node: n.node.Left}
		}
		if n.node.Right != nil {
			right = &TSLNode{node: n.node.Right}
		}
		return TSLExpressionOp{
			Operator: n.node.Operator,
			Left:     left,
			Right:    right,
		}
	case KindUnaryExpr:
		var right *TSLNode
		if n.node.Right != nil {
			right = &TSLNode{node: n.node.Right}
		}
		return TSLExpressionOp{
			Operator: n.node.Operator,
			Left:     nil,
			Right:    right,
		}
	case KindArrayLiteral:
		values := make([]*TSLNode, len(n.node.Children))
		for i, child := range n.node.Children {
			values[i] = &TSLNode{node: child}
		}
		return TSLArrayLiteral{Values: values}
	case KindNullLiteral:
		return "NULL"
	default:
		return nil
	}
}

// SetLeft replaces the left child of a binary expression node
func (n *TSLNode) SetLeft(child *TSLNode) {
	if n == nil || n.node == nil || child == nil {
		return
	}
	n.node.Left = child.node
}

// SetRight replaces the right child of a binary or unary expression node
func (n *TSLNode) SetRight(child *TSLNode) {
	if n == nil || n.node == nil || child == nil {
		return
	}
	n.node.Right = child.node
}
