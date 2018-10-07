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
	"strings"

	"github.com/yaacov/tsl/pkg/parser"
)

// Listener is a Tree search listener.
type Listener struct {
	*parser.BaseTSLListener

	Stack []Node
	Err   error
}

// GetTree return the parsed tree, if exist.
func (l *Listener) GetTree() (n Node, err error) {
	// Check stack size.
	if len(l.Stack) != 1 {
		err = fmt.Errorf("unexpected operator stack")
		return
	}

	n = l.Stack[0]
	return
}

// ExitColumnIdentifier is called when exiting the ColumnIdentifier production.
func (l *Listener) ExitColumnIdentifier(c *parser.ColumnIdentifierContext) {
	n := Node{
		Func: IdentOp,
		Left: c.ColumnName().GetText(),
	}

	l.push(n)
}

// ExitNumberLiteral is called when exiting the NumberLiteral production.
func (l *Listener) ExitNumberLiteral(c *parser.NumberLiteralContext) {
	n := Node{
		Func: NumberOp,
		Left: c.SignedNumber().GetText(),
	}

	l.push(n)
}

// ExitStringLiteral is called when exiting the StringLiteral production.
func (l *Listener) ExitStringLiteral(c *parser.StringLiteralContext) {
	n := Node{
		Func: StringOp,
		Left: stringValueToArg(c.StringValue().GetText()),
	}

	l.push(n)
}

// ExitMulOps is called when production MulOps is exited.
func (l *Listener) ExitMulOps(c *parser.MulOpsContext) {
	right, left := l.pop(), l.pop()

	// Check right op is not a string.
	if right.Func == StringOp {
		l.Err = fmt.Errorf("unexpected a string literal in math function")
		return
	}

	n := Node{
		Func:  MultiplyOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitLiteralOps is called when production LiteralOps is exited.
func (l *Listener) ExitDivOps(c *parser.DivOpsContext) {
	right, left := l.pop(), l.pop()

	// Check right op is not a string.
	if right.Func == StringOp {
		l.Err = fmt.Errorf("unexpected a string literal in math function")
		return
	}

	n := Node{
		Func:  DivideOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitLiteralOps is called when production LiteralOps is exited.
func (l *Listener) ExitModOps(c *parser.ModOpsContext) {
	right, left := l.pop(), l.pop()

	// Check right op is not a string.
	if right.Func == StringOp {
		l.Err = fmt.Errorf("unexpected a string literal in math function")
		return
	}

	n := Node{
		Func:  ModuloOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitLiteralOps is called when production LiteralOps is exited.
func (l *Listener) ExitAddOps(c *parser.AddOpsContext) {
	right, left := l.pop(), l.pop()

	// Check right op is not a string.
	if right.Func == StringOp {
		l.Err = fmt.Errorf("unexpected a string literal in math function")
		return
	}

	n := Node{
		Func:  AddOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitLiteralOps is called when production LiteralOps is exited.
func (l *Listener) ExitSubOps(c *parser.SubOpsContext) {
	right, left := l.pop(), l.pop()

	// Check right op is not a string.
	if right.Func == StringOp {
		l.Err = fmt.Errorf("unexpected a string literal in math function")
		return
	}

	n := Node{
		Func:  SubtractOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitLiteralOps is called when production LiteralOps is exited.
func (l *Listener) ExitLiteralOps(c *parser.LiteralOpsContext) {
	right, left := l.pop(), l.pop()

	n := Node{
		Func:  opDic[c.LiteralOp().GetText()],
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitStringOps is called when production StringOps is exited.
func (l *Listener) ExitStringOps(c *parser.StringOpsContext) {
	right, left := l.pop(), l.pop()

	// Check right op is a string.
	if right.Func != StringOp {
		l.Err = fmt.Errorf("expected a string literal, found %v", right.Func)
		return
	}

	n := Node{
		Func:  opDic[c.StringOp().GetText()],
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitLike is called when production Like is exited.
func (l *Listener) ExitLike(c *parser.LikeContext) {
	right, left := l.pop(), l.pop()
	op := ternaryOp(c.KeyNot() == nil, LikeOp, NotLikeOp)

	// Check right op is a string.
	if right.Func != StringOp {
		l.Err = fmt.Errorf("expected a string literal, found %v", right.Func)
		return
	}

	n := Node{
		Func:  op,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitIsLiteral is called when production IsLiteral is exited.
func (l *Listener) ExitIsLiteral(c *parser.IsLiteralContext) {
	right, left := l.pop(), l.pop()
	op := ternaryOp(c.KeyNot() == nil, EqOp, NotEqOp)

	n := Node{
		Func:  op,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitIsNull is called when production IsNull is exited.
func (l *Listener) ExitIsNull(c *parser.IsNullContext) {
	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, IsNilOp, IsNotNilOp)

	n := Node{
		Func: op,
		Left: left,
	}

	l.push(n)
}

// ExitIn is called when production In is exited.
func (l *Listener) ExitIn(c *parser.InContext) {
	right := l.popLiterals([]Node{})
	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, InOp, NotInOp)

	n := Node{
		Func:  op,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitBetween is called when production Between is exited.
func (l *Listener) ExitBetween(c *parser.BetweenContext) {
	right := []Node{l.pop(), l.pop()}
	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, BetweenOp, NotBetweenOp)

	n := Node{
		Func:  op,
		Left:  left,
		Right: []Node{right[1], right[0]},
	}

	l.push(n)
}

// ExitNot is called when production Not is exited.
func (l *Listener) ExitNot(c *parser.NotContext) {
	left := l.pop()

	n := Node{
		Func: NotOp,
		Left: left,
	}

	l.push(n)
}

// ExitAnd is called when production And is exited.
func (l *Listener) ExitAnd(c *parser.AndContext) {
	right, left := l.pop(), l.pop()
	n := Node{
		Func:  AndOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitOr is called when production Or is exited.
func (l *Listener) ExitOr(c *parser.OrContext) {
	right, left := l.pop(), l.pop()
	n := Node{
		Func:  OrOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ternaryOp return lh if conditional is true, rh o/w.
func ternaryOp(conditional bool, lh string, rh string) string {
	if conditional {
		return lh
	}

	return rh
}

// stringValueToArg clean string value, and return an arg string.
func stringValueToArg(v string) (arg interface{}) {
	l := len(v)

	// Check string length.
	if l < 1 {
		return
	}

	arg = strings.Replace(v[1:l-1], "''", "'", -1)
	return
}

// popLiterals collect literal values, and create args list.
func (l *Listener) popLiterals(in []Node) (out []Node) {
	// Get a literal from stack.
	p := l.pop()

	// Run recurtion on literals.
	if p.Func == StringOp || p.Func == NumberOp {
		return l.popLiterals(append(in, p))
	}

	// If p is not a literal, add it back to stack and exit.
	l.push(p)
	return in
}

// push is a helper function for pushing new node to the listener Stack.
func (l *Listener) push(i Node) {
	l.Stack = append(l.Stack, i)
}

// pop is a helper function for poping a node from the listener Stack.
func (l *Listener) pop() (n Node) {
	// Check that we have nodes in the stack.
	size := len(l.Stack)
	if size < 1 {
		l.Err = fmt.Errorf("operator stack is empty")
		return
	}

	// Pop the last value from the Stack.
	n, l.Stack = l.Stack[size-1], l.Stack[:size-1]

	return
}