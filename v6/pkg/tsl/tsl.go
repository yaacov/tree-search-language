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

//go:generate bash ../parser/build.sh

package tsl

/*
#cgo CFLAGS: -I../parser

#include <stdlib.h>
#include "tsl.h"
*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// TSLNode represents an tree search language AST (Abstract Syntax Tree) node
type TSLNode struct {
	cNode *C.ast_node
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
	cInput := C.CString(input)
	defer C.free(unsafe.Pointer(cInput))

	cNode := C.parse_input_string(cInput)
	if cNode == nil {
		errStr := C.GoString(C.get_error_string())
		errPos := int(C.get_error_position())
		errCtx := C.GoString(C.get_input_string_at_error())

		// Cleanup parser memory after getting error information
		C.cleanup_parser_memory()

		return nil, fmt.Errorf("%s at position %d:\n%s\n%s", errStr, errPos, input, errCtx)
	}

	// Cleanup parser memory after successful parse
	C.cleanup_parser_memory()
	return &TSLNode{cNode: cNode}, nil
}

// Free releases the memory allocated for the AST
func (n *TSLNode) Free() {
	if n.cNode != nil {
		C.ast_free(n.cNode)
		n.cNode = nil
	}
}

// Type returns the type of the node
func (n *TSLNode) Type() NodeKind {
	return NodeKind(n.cNode._type)
}

// Value returns the node's value based on its type
func (n *TSLNode) Value() interface{} {
	switch n.Type() {
	case KindBooleanLiteral:
		return C.get_boolean_value(n.cNode) != 0
	case KindNumericLiteral:
		// Keep numeric values as float64 consistently
		return float64(C.get_number_value(n.cNode))
	case KindStringLiteral, KindIdentifier:
		return C.GoString(C.get_string_value(n.cNode))
	case KindDateLiteral, KindTimestampLiteral:
		// Parse date strings into time.Time
		dateStr := C.GoString(C.get_string_value(n.cNode))
		if t, err := time.Parse(time.RFC3339, dateStr); err == nil {
			return t
		}
		// Fallback to returning raw string if parsing fails
		return dateStr
	case KindBinaryExpr:
		return TSLExpressionOp{
			Operator: Operator(C.get_binary_op(n.cNode)),
			Left:     &TSLNode{cNode: C.get_binary_left(n.cNode)},
			Right:    &TSLNode{cNode: C.get_binary_right(n.cNode)},
		}
	case KindUnaryExpr:
		return TSLExpressionOp{
			Operator: Operator(C.get_unary_op(n.cNode)),
			Left:     nil,
			Right:    &TSLNode{cNode: C.get_unary_child(n.cNode)},
		}
	case KindArrayLiteral:
		size := int(C.get_array_size(n.cNode))
		values := make([]*TSLNode, size)
		for i := 0; i < size; i++ {
			values[i] = &TSLNode{cNode: C.get_array_element(n.cNode, C.int(i))}
		}
		return TSLArrayLiteral{Values: values}
	case KindNullLiteral:
		return "NULL"
	default:
		return nil
	}
}
