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
// Disable linting because this function swith is too large:
// nolint
func SquirrelWalk(n Node) (s sq.Sqlizer, err error) {
	var l, r sq.Sqlizer
	var sql string

	switch n.Func {
	case IdentOp:
		s = sq.Expr(n.Left.(string))
	case AndOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		r, err = SquirrelWalk(n.Right.(Node))
		if err != nil {
			return
		}
		s = sq.And{l, r}
	case OrOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		r, err = SquirrelWalk(n.Right.(Node))
		if err != nil {
			return
		}
		s = sq.Or{l, r}
	case NotOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		s = notExpr{l}
	case AddOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		r, err = SquirrelWalk(n.Right.(Node))
		if err != nil {
			return
		}
		s = addExpr{l, r}
	case EqOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		s = sq.Eq{sql: n.Right}
	case NotEqOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		s = sq.NotEq{sql: n.Right}
	case LtOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		s = sq.Lt{sql: n.Right}
	case LteOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		s = sq.LtOrEq{sql: n.Right}
	case GtOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		s = sq.Gt{sql: n.Right}
	case GteOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		s = sq.GtOrEq{sql: n.Right}
	case InOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		// Multiple eq will be translated into IN (?, ? ...).
		s = sq.Eq{sql: n.Right}
	case NotInOp:
		// Multiple not eq will be translated into NOT IN (?, ? ...).
		s = sq.NotEq{n.Left.(string): n.Right}
	case IsNilOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		// eq nil will be translated into IS NULL.
		s = sq.Eq{sql: nil}
	case IsNotNilOp:
		l, err = SquirrelWalk(n.Left.(Node))
		if err != nil {
			return
		}
		sql, _, err = l.ToSql()
		// not eq nil will be translated into IS NOT NULL.
		s = sq.NotEq{sql: nil}
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
