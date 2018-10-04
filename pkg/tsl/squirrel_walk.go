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
)

// binaryStep handle a binary operator step for SquirrelWalk.
func binaryStep(n Node) (s sq.Sqlizer, err error) {
	var l, r sq.Sqlizer

	// Left hand side must be a Node.
	l, err = SquirrelWalk(n.Left.(Node))
	if err != nil {
		return
	}

	// Check if right hand side is Node or string.
	if str, ok := n.Right.(string); ok {
		// Right hand side is a string.
		r = sq.Expr(str)
	} else {
		// Right hand side is a Node.
		r, err = SquirrelWalk(n.Right.(Node))
		if err != nil {
			return
		}
	}

	switch n.Func {
	case AndOp:
		s = sq.And{l, r}
	case OrOp:
		s = sq.Or{l, r}
	case AddOp:
		s = addExpr{l, r}
	case SubtractOp:
		s = subExpr{l, r}
	case MultiplyOp:
		s = mulExpr{l, r}
	case DivideOp:
		s = divExpr{l, r}
	case ModuloOp:
		s = modExpr{l, r}
	default:
		// If here than the operator is not supported.
		err = fmt.Errorf("un supported operand: %s", n.Func)
	}

	return
}

// unaryStep handle a unary operator step for SquirrelWalk.
func unaryStep(n Node) (s sq.Sqlizer, err error) {
	var l sq.Sqlizer
	var sql string

	l, err = SquirrelWalk(n.Left.(Node))
	if err != nil {
		return
	}

	sql, _, err = l.ToSql()
	if err != nil {
		return
	}

	switch n.Func {
	case NotOp:
		s = notExpr{l}
	case EqOp:
		s = sq.Eq{sql: n.Right}
	case NotEqOp:
		s = sq.NotEq{sql: n.Right}
	case LtOp:
		s = sq.Lt{sql: n.Right}
	case LteOp:
		s = sq.LtOrEq{sql: n.Right}
	case GtOp:
		s = sq.Gt{sql: n.Right}
	case GteOp:
		s = sq.GtOrEq{sql: n.Right}
	case InOp:
		// Multiple eq will be translated into IN (?, ? ...).
		s = sq.Eq{sql: n.Right}
	case NotInOp:
		// Multiple not eq will be translated into NOT IN (?, ? ...).
		s = sq.NotEq{n.Left.(string): n.Right}
	case IsNilOp:
		// eq nil will be translated into IS NULL.
		s = sq.Eq{sql: nil}
	case IsNotNilOp:
		// not eq nil will be translated into IS NOT NULL.
		s = sq.NotEq{sql: nil}
	default:
		// If here than the operator is not supported.
		err = fmt.Errorf("un supported operand: %s", n.Func)
	}

	return
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
// Squirrel: https://github.com/Masterminds/squirrel
//
func SquirrelWalk(n Node) (s sq.Sqlizer, err error) {
	switch n.Func {
	case IdentOp:
		s = sq.Expr(n.Left.(string))
	case AndOp, OrOp, AddOp, SubtractOp, MultiplyOp, DivideOp, ModuloOp:
		return binaryStep(n)
	case NotOp, EqOp, NotEqOp, LtOp, LteOp, GtOp, GteOp, InOp, NotInOp, IsNilOp, IsNotNilOp:
		return unaryStep(n)
	case LikeOp:
		t := fmt.Sprintf("%s LIKE ?", n.Left.(string))
		s = sq.Expr(t, n.Right)
	case NotLikeOp:
		t := fmt.Sprintf("%s NOT LIKE ?", n.Left.(string))
		s = sq.Expr(t, n.Right)
	case BetweenOp:
		t := fmt.Sprintf("%s BETWEEN ? AND ?", n.Left.(string))
		s = sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1])
	case NotBetweenOp:
		t := fmt.Sprintf("%s NOT BETWEEN ? AND ?", n.Left.(string))
		s = sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1])
	default:
		// If here than the operator is not supported.
		err = fmt.Errorf("un supported operand: %s", n.Func)
	}

	return
}
