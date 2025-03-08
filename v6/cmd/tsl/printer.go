package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

// ASTPrinter handles the formatted output of TSL abstract syntax trees
type ASTPrinter struct {
	indentSize int
}

// NewASTPrinter creates a new ASTPrinter with default settings
func NewASTPrinter() *ASTPrinter {
	return &ASTPrinter{
		indentSize: 2,
	}
}

// Print outputs a node and its children with proper indentation
func (p *ASTPrinter) Print(node *tsl.TSLNode, level int) {
	if node == nil {
		return
	}

	t := node.Type()

	switch t {
	case tsl.KindBinaryExpr:
		p.printBinaryOp(node, level)
	case tsl.KindUnaryExpr:
		p.printUnaryOp(node, level)
	case tsl.KindArrayLiteral:
		p.printArray(node, level)
	case tsl.KindTimestampLiteral:
		p.printTimestamp(node, level)
	case tsl.KindStringLiteral:
		p.printString(node, level)
	default:
		// Leaf nodes
		p.printIndented(level, "[%s]: %v\n", t.String(), node.Value())
	}
}

// printString formats and prints string values with escaped special characters
func (p *ASTPrinter) printString(node *tsl.TSLNode, level int) {
	str, ok := node.Value().(string)
	if !ok {
		p.printIndented(level, "[%s]: %v\n", node.Type().String(), node.Value())
		return
	}

	// Use Go's built-in escaping functionality
	quoted := strconv.Quote(str)
	// Remove the surrounding quotes to get just the escaped string
	escaped := quoted[1 : len(quoted)-1]

	p.printIndented(level, "[%s]: %s\n", node.Type().String(), escaped)
}

func (p *ASTPrinter) printBinaryOp(node *tsl.TSLNode, level int) {
	v := node.Value().(tsl.TSLExpressionOp)

	p.printIndented(level, "[%s]\n", v.Operator.String())
	p.Print(v.Left, level+1)
	p.Print(v.Right, level+1)
}

func (p *ASTPrinter) printUnaryOp(node *tsl.TSLNode, level int) {
	v := node.Value().(tsl.TSLExpressionOp)

	p.printIndented(level, "[%s]\n", v.Operator.String())
	if v.Right != nil {
		p.Print(v.Right, level+1)
	}
}

func (p *ASTPrinter) printArray(node *tsl.TSLNode, level int) {
	v := node.Value().(tsl.TSLArrayLiteral)

	p.printIndented(level, "[%s]:\n", node.Type().String())
	for _, elem := range v.Values {
		p.Print(elem, level+1)
	}
}

// printTimestamp formats and prints timestamp values with timezone offset but without timezone name
func (p *ASTPrinter) printTimestamp(node *tsl.TSLNode, level int) {
	timestamp, ok := node.Value().(time.Time)
	if !ok {
		p.printIndented(level, "[%s]: %v\n", node.Type().String(), node.Value())
		return
	}

	// Format with timezone offset but no timezone name
	// 2006-01-02T15:04:05-07:00 is the Go layout for RFC3339 without timezone name
	formattedTime := timestamp.Format("2006-01-02T15:04:05-07:00")
	p.printIndented(level, "[%s]: %s\n", node.Type().String(), formattedTime)
}

func (p *ASTPrinter) printIndented(level int, format string, a ...interface{}) {
	indent := p.getIndentation(level)
	fmt.Printf(indent+format, a...)
}

func (p *ASTPrinter) getIndentation(level int) string {
	return fmt.Sprintf("%*s", level*p.indentSize, "")
}
