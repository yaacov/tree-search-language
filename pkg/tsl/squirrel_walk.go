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

func nodesToStrings(in interface{}) (s []string) {
	var nn []Node
	var ok bool

	// Assume input is []Node or Node.
	nn, ok = in.([]Node)
	if !ok {
		nn = []Node{in.(Node)}
	}

	// Assume all Nodes are Leafs.
	for _, n := range nn {
		s = append(s, n.Left.(string))
	}

	return
}

// binaryStep handle a binary operator step for SquirrelWalk.
func binaryStep(n Node) (s sq.Sqlizer, err error) {
	var l, r sq.Sqlizer

	// Get left hand side node.
	l, err = SquirrelWalk(n.Left.(Node))
	if err != nil {
		return
	}

	// Get right hand side node.
	r, err = SquirrelWalk(n.Right.(Node))
	if err != nil {
		return
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

	right := nodesToStrings(n.Right)

	switch n.Func {
	case NotOp:
		s = notExpr{l}
	case EqOp:
		s = sq.Eq{sql: right[0]}
	case NotEqOp:
		s = sq.NotEq{sql: right[0]}
	case LtOp:
		s = sq.Lt{sql: right[0]}
	case LteOp:
		s = sq.LtOrEq{sql: right[0]}
	case GtOp:
		s = sq.Gt{sql: right[0]}
	case GteOp:
		s = sq.GtOrEq{sql: right[0]}
	case InOp:
		// Multiple eq will be translated into IN (?, ? ...).
		s = sq.Eq{sql: right}
	case NotInOp:
		// Multiple not eq will be translated into NOT IN (?, ? ...).
		s = sq.NotEq{sql: right}
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
	var l sq.Sqlizer
	var sql string

	switch n.Func {
	case IdentOp:
		s = sq.Expr(n.Left.(string))
	case NumberOp:
		s = sq.Expr(n.Left.(string))
	case StringOp:
		s = sq.Expr(n.Left.(string))
	case AndOp, OrOp, AddOp, SubtractOp, MultiplyOp, DivideOp, ModuloOp:
		return binaryStep(n)
	case NotOp, EqOp, NotEqOp, LtOp, LteOp, GtOp, GteOp, InOp, NotInOp, IsNilOp, IsNotNilOp:
		return unaryStep(n)
	case LikeOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}

		sql, _, err = l.ToSql()
		if err != nil {
			return
		}

		t := fmt.Sprintf("%s LIKE ?", sql)
		s = sq.Expr(t, n.Right)
	case NotLikeOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}

		sql, _, err = l.ToSql()
		if err != nil {
			return
		}

		t := fmt.Sprintf("%s NOT LIKE ?", sql)
		s = sq.Expr(t, n.Right.(Node).Left.(string))
	case BetweenOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}

		sql, _, err = l.ToSql()
		if err != nil {
			return
		}

		t := fmt.Sprintf("%s BETWEEN ? AND ?", sql)
		s = sq.Expr(t, n.Right.([]Node)[0].Left.(string), n.Right.([]Node)[1].Left.(string))
	case NotBetweenOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}

		sql, _, err = l.ToSql()
		if err != nil {
			return
		}

		t := fmt.Sprintf("%s NOT BETWEEN ? AND ?", sql)
		s = sq.Expr(t, n.Right.([]Node)[0].Left.(string), n.Right.([]Node)[1].Left.(string))
	default:
		// If here than the operator is not supported.
		err = fmt.Errorf("un supported operand: %s", n.Func)
	}

	return
}
