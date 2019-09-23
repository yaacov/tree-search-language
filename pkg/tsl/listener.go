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
	"strconv"
	"strings"

	"github.com/yaacov/tree-search-language/pkg/parser"
)

// Listener is a Tree search listener.
type Listener struct {
	*parser.BaseTSLListener

	Stack []Node
	Errs  []error
}

// GetTree return the parsed tree, if exist.
func (l *Listener) GetTree() (n Node, err error) {
	// Check for errors.
	if len(l.Errs) > 0 {
		err = l.Errs[0]
		return
	}

	// Check stack size.
	if len(l.Stack) != 1 {
		err = StackError{}
		return
	}

	n = l.Stack[0]
	return
}

// ExitColumnIdentifier is called when exiting the ColumnIdentifier production.
func (l *Listener) ExitColumnIdentifier(c *parser.ColumnIdentifierContext) {
	l.exitLiteral(IdentOp, c.ColumnName().GetText())
}

// ExitNumberLiteral is called when exiting the NumberLiteral production.
func (l *Listener) ExitNumberLiteral(c *parser.NumberLiteralContext) {
	// Check for a float value.
	f, err := strconv.ParseFloat(c.SignedNumber().GetText(), 64)
	if err != nil {
		l.Errs = append(l.Errs, err)
	}

	l.exitLiteral(NumberOp, f)
}

// ExitStringLiteral is called when exiting the StringLiteral production.
func (l *Listener) ExitStringLiteral(c *parser.StringLiteralContext) {
	// StringValue must be a string of format \'.*'\,
	// length must be greater or equal to 2.
	s := c.StringValue().GetText()
	ln := len(s)
	v := strings.Replace(s[1:ln-1], "''", "'", -1)

	l.exitLiteral(StringOp, v)
}

// ExitMulOps is called when production multiply op is exited.
func (l *Listener) ExitMulOps(c *parser.MulOpsContext) {
	l.exitMathOps(MultiplyOp)
}

// ExitDivOps is called when production div op is exited.
func (l *Listener) ExitDivOps(c *parser.DivOpsContext) {
	l.exitMathOps(DivideOp)
}

// ExitModOps is called when production modulo op is exited.
func (l *Listener) ExitModOps(c *parser.ModOpsContext) {
	l.exitMathOps(ModuloOp)
}

// ExitAddOps is called when production add op is exited.
func (l *Listener) ExitAddOps(c *parser.AddOpsContext) {
	l.exitMathOps(AddOp)
}

// ExitSubOps is called when production subtract op is exited.
func (l *Listener) ExitSubOps(c *parser.SubOpsContext) {
	l.exitMathOps(SubtractOp)
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
		l.Errs = append(l.Errs, UnexpectedLiteralError{ExpectedType: "string", Literal: right.Left})
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

	// Check the operator to use according to the given LIKE or ILIKE keyword:
	var op string
	switch c.LikeOp().GetStart().GetTokenType() {
	case parser.TSLLexerK_LIKE:
		op = LikeOp
	case parser.TSLLexerK_ILIKE:
		op = ILikeOp
	default:
		// This will never happen, as the grammar doesn't allow it, however this will catch
		// errors if in the future we add new operators to the grammar but we forget to
		// update this code.
		l.Errs = append(l.Errs, fmt.Errorf("expected LIKE or ILIKE"))
		return
	}

	// Check right op is a string.
	if right.Func != StringOp {
		l.Errs = append(l.Errs, UnexpectedLiteralError{ExpectedType: "string", Literal: right.Left})
		return
	}

	// Create the node for the LIKE or ILIKE operation:
	n := Node{
		Func:  op,
		Left:  left,
		Right: right,
	}

	// If there is a NOT before the LIKE or ILIKE operations then create an
	// additional node for the negation:
	if c.KeyNot() != nil {
		n = Node{
			Func: NotOp,
			Left: n,
		}
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
	right := Node{
		Func:  ArrayOp,
		Right: l.popLiterals([]Node{}),
	}
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
	nodes := []Node{l.pop(), l.pop()}
	right := Node{
		Func:  ArrayOp,
		Right: []Node{nodes[1], nodes[0]},
	}

	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, BetweenOp, NotBetweenOp)

	n := Node{
		Func:  op,
		Left:  left,
		Right: right,
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

// exitMathOps is called when production math operator is exited.
func (l *Listener) exitMathOps(op string) {
	right, left := l.pop(), l.pop()

	// Check right op is not a string.
	if right.Func == StringOp {
		l.Errs = append(l.Errs, UnexpectedLiteralError{ExpectedType: "float", Literal: right.Left})
		return
	}

	n := Node{
		Func:  op,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitColumnIdentifier is called when exiting a literal production.
func (l *Listener) exitLiteral(f string, v interface{}) {
	n := Node{Func: f, Left: v}
	l.push(n)
}

// ternaryOp return lh if conditional is true, rh o/w.
func ternaryOp(conditional bool, lh string, rh string) string {
	if conditional {
		return lh
	}

	return rh
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
		l.Errs = append(l.Errs, StackError{})
		return
	}

	// Pop the last value from the Stack.
	n, l.Stack = l.Stack[size-1], l.Stack[:size-1]

	return
}
