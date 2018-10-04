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

// ExitLiteralOps is called when production LiteralOps is exited.
func (l *Listener) ExitLiteralOps(c *parser.LiteralOpsContext) {
	left := l.pop()

	n := Node{
		Func:  opDic[c.LiteralOp().GetText()],
		Left:  left,
		Right: literalValueToArg(c.LiteralValue().GetText()),
	}

	l.push(n)
}

// ExitColumn is called when entering the Column production.
func (l *Listener) ExitColumn(c *parser.ColumnContext) {
	n := Node{
		Func: IdentOp,
		Left: c.ColumnName().GetText(),
	}

	l.push(n)
}

// ExitColumnNameOp is called when entering the ColumnNameOp production.
func (l *Listener) ExitColumnNameOp(c *parser.ColumnNameOpContext) {
	right, left := l.pop(), l.pop()

	n := Node{
		Func:  opDic[c.NumericOp().GetText()],
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitStringOps is called when production StringOps is exited.
func (l *Listener) ExitStringOps(c *parser.StringOpsContext) {
	left := l.pop()

	n := Node{
		Func:  opDic[c.StringOp().GetText()],
		Left:  left,
		Right: literalValueToArg(c.StringValue().GetText()),
	}

	l.push(n)
}

// ExitLike is called when production Like is exited.
func (l *Listener) ExitLike(c *parser.LikeContext) {
	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, LikeOp, NotLikeOp)

	n := Node{
		Func:  op,
		Left:  left,
		Right: literalValueToArg(c.StringValue().GetText()),
	}

	l.push(n)
}

// ExitIsLiteral is called when production IsLiteral is exited.
func (l *Listener) ExitIsLiteral(c *parser.IsLiteralContext) {
	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, EqOp, NotEqOp)

	n := Node{
		Func:  op,
		Left:  left,
		Right: literalValueToArg(c.LiteralValue().GetText()),
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
	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, InOp, NotInOp)
	args := literalValuesToArgs(c)

	n := Node{
		Func:  op,
		Left:  left,
		Right: args,
	}

	l.push(n)
}

// ExitBetween is called when production Between is exited.
func (l *Listener) ExitBetween(c *parser.BetweenContext) {
	left := l.pop()
	op := ternaryOp(c.KeyNot() == nil, BetweenOp, NotBetweenOp)
	args := literalValuesToArgs(c)

	n := Node{
		Func:  op,
		Left:  left,
		Right: args,
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

// hasLiteralValues is a type that contain literal values.
type hasLiteralValues interface {
	AllLiteralValue() []parser.ILiteralValueContext
	LiteralValue(int) parser.ILiteralValueContext
}

// ternaryOp return lh if conditional is true, rh o/w.
func ternaryOp(conditional bool, lh string, rh string) string {
	if conditional {
		return lh
	}

	return rh
}

// literalValueToArg clean literal value, and return an arg string or float.
func literalValueToArg(v string) (arg interface{}) {
	l := len(v)

	// Check string length.
	if l < 1 {
		return
	}

	// Check for ''.
	if v[0] == '\'' {
		v = v[1 : l-1]
		arg = strings.Replace(v, "''", "'", -1)
		return
	}

	// It's a float.
	arg, _ = strconv.ParseFloat(v, 64)

	return
}

// literalValuesToArgs collect literal values, and create args list.
func literalValuesToArgs(c hasLiteralValues) (args []interface{}) {
	// Get length of literal values list.
	l := len(c.AllLiteralValue())

	// Create the arg list.
	args = make([]interface{}, l)
	for i := 0; i < l; i++ {
		args[i] = literalValueToArg(c.LiteralValue(i).GetText())
	}

	return
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
