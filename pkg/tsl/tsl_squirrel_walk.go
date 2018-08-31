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
	case EqOp:
		s = sq.Eq{n.Left.(string): n.Right}
	case NotEqOp:
		s = sq.NotEq{n.Left.(string): n.Right}
	case LtOp:
		s = sq.Lt{n.Left.(string): n.Right}
	case LteOp:
		s = sq.LtOrEq{n.Left.(string): n.Right}
	case GtOp:
		s = sq.Gt{n.Left.(string): n.Right}
	case GteOp:
		s = sq.GtOrEq{n.Left.(string): n.Right}
	case InOp:
		s = sq.Eq{n.Left.(string): n.Right}
	case NotInOp:
		s = sq.NotEq{n.Left.(string): n.Right}
	case IsNilOp:
		s = sq.Eq{n.Left.(string): nil}
	case IsNotNilOp:
		s = sq.NotEq{n.Left.(string): nil}
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
		err = fmt.Errorf("un supported operand: %s", n.Func)
	}

	return
}
