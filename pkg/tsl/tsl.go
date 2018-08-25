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

	sq "github.com/Masterminds/squirrel"
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/yaacov/tsl/pkg/parser"
)

// ParseTSL parses the input string into TSL tree.
func ParseTSL(input string) (tree Node, err error) {
	// Setup the ErrorListener
	errorListener := NewErrorListener()

	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewTSLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewTSLParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)

	// Parse the expression (by walking the tree)
	var listener Listener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// Check for errors
	err = errorListener.Err
	if err != nil {
		return
	}

	// Get the parsed tree
	tree, err = listener.GetTree()

	return
}

// GetTree return the parsed tree, if exist.
func (l *Listener) GetTree() (n Node, err error) {
	// Check stack size
	if len(l.Stack) < 1 {
		err = fmt.Errorf("operator stack is empty")
		return
	}

	return l.Stack[0], l.Err
}

// SquirrelWalk travel the TSL tree to create squirrel SQL select operators.
//
// Users can call the SquirrelWalk method inside a squirrel Where to add the query.
//
//  filter, _ := SquirrelWalk(tree)
//  sql, args, _ := sq.Select("name, city, state").
//    From("users").
//    Where(filter).
//    ToSql()
//
func SquirrelWalk(n Node) (s sq.Sqlizer, err error) {
	var l, r sq.Sqlizer

	switch n.Func {
	case AndOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		r, err = SquirrelWalk(n.Right.(Node))
		if err != nil {
			return
		}
		return sq.And{l, r}, nil
	case OrOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		r, err = SquirrelWalk(n.Right.(Node))
		if err != nil {
			return
		}
		return sq.Or{l, r}, nil
	case NotOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		return notExpr{l}, nil
	case EqOp:
		return sq.Eq{n.Left.(string): n.Right}, nil
	case NotEqOp:
		return sq.NotEq{n.Left.(string): n.Right}, nil
	case LtOp:
		return sq.Lt{n.Left.(string): n.Right}, nil
	case LteOp:
		return sq.LtOrEq{n.Left.(string): n.Right}, nil
	case GtOp:
		return sq.Gt{n.Left.(string): n.Right}, nil
	case GteOp:
		return sq.GtOrEq{n.Left.(string): n.Right}, nil
	case InOp:
		return sq.Eq{n.Left.(string): n.Right}, nil
	case NotInOp:
		return sq.NotEq{n.Left.(string): n.Right}, nil
	case IsNilOp:
		return sq.Eq{n.Left.(string): nil}, nil
	case IsNotNilOp:
		return sq.NotEq{n.Left.(string): nil}, nil
	case LikeOp:
		t := fmt.Sprintf("%s LIKE ?", n.Left.(string))
		return sq.Expr(t, n.Right), nil
	case NotLikeOp:
		t := fmt.Sprintf("%s NOT LIKE ?", n.Left.(string))
		return sq.Expr(t, n.Right), nil
	case BetweenOp:
		t := fmt.Sprintf("%s BETWEEN ? AND ?", n.Left.(string))
		return sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1]), nil
	case NotBetweenOp:
		t := fmt.Sprintf("%s NOT BETWEEN ? AND ?", n.Left.(string))
		return sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1]), nil
	default:
		err = fmt.Errorf("un supported operand: %s", n.Func)
		return
	}

	return
}
