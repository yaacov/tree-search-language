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

// AddToSelect adds a TSL tree into a squirrel SelectBuilder.
func AddToSelect(s sq.SelectBuilder, tree Node) sq.SelectBuilder {
	return s.Where(SquirrelWalk(tree))
}

// SquirrelWalk travel the TSL tree to create squirrel SQL select operators.
//
// Users can call the SquirrelWalk method inside a squirrel Where to add the query.
//
// 	sql, args, _ := sq.Select("name, city, state").
// 		Where(SquirrelWalk(tree)).
// 		From("users").
// 		ToSql()
//
func SquirrelWalk(n Node) sq.Sqlizer {
	switch n.Func {
	case AndOp:
		return sq.And{SquirrelWalk(n.Left.(Node)), SquirrelWalk(n.Right.(Node))}
	case OrOp:
		return sq.Or{SquirrelWalk(n.Left.(Node)), SquirrelWalk(n.Right.(Node))}
	case NotOp:
		return notExpr{SquirrelWalk(n.Left.(Node))}
	case EqOp:
		return sq.Eq{n.Left.(string): n.Right}
	case NotEqOp:
		return sq.NotEq{n.Left.(string): n.Right}
	case LtOp:
		return sq.Lt{n.Left.(string): n.Right}
	case LteOp:
		return sq.LtOrEq{n.Left.(string): n.Right}
	case GtOp:
		return sq.Gt{n.Left.(string): n.Right}
	case GteOp:
		return sq.GtOrEq{n.Left.(string): n.Right}
	case InOp:
		return sq.Eq{n.Left.(string): n.Right}
	case NotInOp:
		return sq.NotEq{n.Left.(string): n.Right}
	case IsNilOp:
		return sq.Eq{n.Left.(string): nil}
	case IsNotNilOp:
		return sq.NotEq{n.Left.(string): nil}
	case LikeOp:
		t := fmt.Sprintf("%s LIKE ?", n.Left.(string))
		return sq.Expr(t, n.Right)
	case NotLikeOp:
		t := fmt.Sprintf("%s NOT LIKE ?", n.Left.(string))
		return sq.Expr(t, n.Right)
	case BetweenOp:
		t := fmt.Sprintf("%s BETWEEN ? AND ?", n.Left.(string))
		return sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1])
	case NotBetweenOp:
		t := fmt.Sprintf("%s NOT BETWEEN ? AND ?", n.Left.(string))
		return sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1])
	}

	return sq.And{}
}
