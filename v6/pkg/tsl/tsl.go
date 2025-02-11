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
	"unsafe"
)

// TSLNode represents an tree search language AST (Abstract Syntax Tree) node
type TSLNode struct {
	cNode *C.ast_node
}

// ExpressionOp represents both unary and binary operations
type ExpressionOp struct {
	Operator    OperatorType
	Left, Right *TSLNode
}

// ArrayLiteral represents an array of nodes
type ArrayLiteral struct {
	Values []*TSLNode
}

// Parse parses a TSL expression and returns the AST root node
func Parse(input string) (*TSLNode, error) {
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
		return float64(C.get_number_value(n.cNode))
	case KindStringLiteral, KindIdentifier, KindDateLiteral, KindTimestampLiteral:
		return C.GoString(C.get_string_value(n.cNode))
	case KindBinaryExpr:
		return ExpressionOp{
			Operator: OperatorType(C.get_binary_op(n.cNode)),
			Left:     &TSLNode{cNode: C.get_binary_left(n.cNode)},
			Right:    &TSLNode{cNode: C.get_binary_right(n.cNode)},
		}
	case KindUnaryExpr:
		return ExpressionOp{
			Operator: OperatorType(C.get_unary_op(n.cNode)),
			Left:     nil,
			Right:    &TSLNode{cNode: C.get_unary_child(n.cNode)},
		}
	case KindArrayLiteral:
		size := int(C.get_array_size(n.cNode))
		values := make([]*TSLNode, size)
		for i := 0; i < size; i++ {
			values[i] = &TSLNode{cNode: C.get_array_element(n.cNode, C.int(i))}
		}
		return ArrayLiteral{Values: values}
	case KindNullLiteral:
		return "NULL"
	default:
		return nil
	}
}
