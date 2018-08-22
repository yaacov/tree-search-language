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
	"github.com/yaacov/tsl/pkg/parser"
)

// ExitLiteralOps is called when production LiteralOps is exited.
func (l *Listener) ExitLiteralOps(c *parser.LiteralOpsContext) {
	n := Node{
		Func:  opDic[c.LiteralOp().GetText()],
		Left:  c.ColumnName().GetText(),
		Right: c.LiteralValue().GetText(),
	}

	l.push(n)
}

// ExitStringOps is called when production StringOps is exited.
func (l *Listener) ExitStringOps(c *parser.StringOpsContext) {
	n := Node{
		Func:  opDic[c.StringOp().GetText()],
		Left:  c.ColumnName().GetText(),
		Right: c.StringValue().GetText(),
	}

	l.push(n)
}

// ExitLike is called when production Like is exited.
func (l *Listener) ExitLike(c *parser.LikeContext) {
	op := ternaryOp(c.KeyNot() == nil, likeOp, notLikeOp)
	n := Node{
		Func:  op,
		Left:  c.ColumnName().GetText(),
		Right: c.StringValue().GetText(),
	}

	l.push(n)
}

// ExitIsLiteral is called when production IsLiteral is exited.
func (l *Listener) ExitIsLiteral(c *parser.IsLiteralContext) {
	op := ternaryOp(c.KeyNot() == nil, eqOp, notEqOp)
	n := Node{
		Func:  op,
		Left:  c.ColumnName().GetText(),
		Right: c.LiteralValue().GetText(),
	}

	l.push(n)
}

// ExitIsNull is called when production IsNull is exited.
func (l *Listener) ExitIsNull(c *parser.IsNullContext) {
	op := ternaryOp(c.KeyNot() == nil, isNilOp, isNotNilOp)
	n := Node{
		Func: op,
		Left: c.ColumnName().GetText(),
	}

	l.push(n)
}

// ExitIn is called when production In is exited.
func (l *Listener) ExitIn(c *parser.InContext) {
	op := ternaryOp(c.KeyNot() == nil, inOp, notInOp)
	args := literalValuesToArgs(c)
	n := Node{
		Func:  op,
		Left:  c.ColumnName().GetText(),
		Right: args,
	}

	l.push(n)
}

// ExitBetween is called when production Between is exited.
func (l *Listener) ExitBetween(c *parser.BetweenContext) {
	op := ternaryOp(c.KeyNot() == nil, betweenOp, notBetweenOp)
	args := literalValuesToArgs(c)
	n := Node{
		Func:  op,
		Left:  c.ColumnName().GetText(),
		Right: args,
	}

	l.push(n)
}

// ExitNot is called when production Not is exited.
func (l *Listener) ExitNot(c *parser.NotContext) {
	left := l.pop()
	n := Node{
		Func: notOp,
		Left: left,
	}

	l.push(n)
}

// ExitAnd is called when production And is exited.
func (l *Listener) ExitAnd(c *parser.AndContext) {
	left, right := l.pop(), l.pop()
	n := Node{
		Func:  andOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}

// ExitOr is called when production Or is exited.
func (l *Listener) ExitOr(c *parser.OrContext) {
	left, right := l.pop(), l.pop()
	n := Node{
		Func:  orOp,
		Left:  left,
		Right: right,
	}

	l.push(n)
}
