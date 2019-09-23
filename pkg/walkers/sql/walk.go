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

// Package sql helps to create SQL squirrel filters using the TSL package.
package sql

import (
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"

	"github.com/yaacov/tree-search-language/pkg/tsl"
)

func nodesToStrings(in interface{}) (s []interface{}) {
	var nn []tsl.Node

	// Check for nil
	if in == nil {
		return
	}

	// Assume in is a node.
	inNode := in.(tsl.Node)

	// Check for array node type.
	if inNode.Func == tsl.ArrayOp {
		// Take all nodes array from array RSE.
		nn = inNode.Right.([]tsl.Node)
	} else {
		// Make a node array out of this node.
		nn = []tsl.Node{inNode}
	}

	// Assume all Nodes are Leafs.
	for _, n := range nn {
		s = append(s, n.Left)
	}

	return
}

// binaryStep handle a binary operator step for Walk.
func binaryStep(n tsl.Node) (s sq.Sqlizer, err error) {
	var l, r sq.Sqlizer

	// Get left hand side node.
	l, err = Walk(n.Left.(tsl.Node))
	if err != nil {
		return
	}

	// Get right hand side node.
	r, err = Walk(n.Right.(tsl.Node))
	if err != nil {
		return
	}

	switch n.Func {
	case tsl.AndOp:
		s = sq.And{l, r}
	case tsl.OrOp:
		s = sq.Or{l, r}
	case tsl.AddOp:
		s = addExpr{l, r}
	case tsl.SubtractOp:
		s = subExpr{l, r}
	case tsl.MultiplyOp:
		s = mulExpr{l, r}
	case tsl.DivideOp:
		s = divExpr{l, r}
	case tsl.ModuloOp:
		s = modExpr{l, r}
	default:
		// If here than the operator is not supported.
		err = tsl.UnexpectedLiteralError{Literal: n.Func}
	}

	return
}

// unaryStep handle a unary operator step for Walk.
func unaryStep(n tsl.Node) (s sq.Sqlizer, err error) {
	var l sq.Sqlizer
	var sql string

	l, err = Walk(n.Left.(tsl.Node))
	if err != nil {
		return
	}

	sql, _, err = l.ToSql()
	if err != nil {
		return
	}

	right := nodesToStrings(n.Right)

	switch n.Func {
	case tsl.NotOp:
		s = notExpr{l}
	case tsl.EqOp:
		s = sq.Eq{sql: right[0]}
	case tsl.NotEqOp:
		s = sq.NotEq{sql: right[0]}
	case tsl.LtOp:
		s = sq.Lt{sql: right[0]}
	case tsl.LteOp:
		s = sq.LtOrEq{sql: right[0]}
	case tsl.GtOp:
		s = sq.Gt{sql: right[0]}
	case tsl.GteOp:
		s = sq.GtOrEq{sql: right[0]}
	case tsl.InOp:
		// Multiple eq will be translated into IN (?, ? ...).
		s = sq.Eq{sql: right}
	case tsl.NotInOp:
		// Multiple not eq will be translated into NOT IN (?, ? ...).
		s = sq.NotEq{sql: right}
	case tsl.IsNilOp:
		// eq nil will be translated into IS NULL.
		s = sq.Eq{sql: nil}
	case tsl.IsNotNilOp:
		// not eq nil will be translated into IS NOT NULL.
		s = sq.NotEq{sql: nil}
	case tsl.LikeOp:
		t := fmt.Sprintf("%s LIKE ?", sql)
		s = sq.Expr(t, right[0])
	case tsl.ILikeOp:
		t := fmt.Sprintf("%s ILIKE ?", sql)
		s = sq.Expr(t, right[0])
	case tsl.BetweenOp:
		t := fmt.Sprintf("%s BETWEEN ? AND ?", sql)
		s = sq.Expr(t, right[0], right[1])
	case tsl.NotBetweenOp:
		t := fmt.Sprintf("%s NOT BETWEEN ? AND ?", sql)
		s = sq.Expr(t, right[0], right[1])
	default:
		// If here than the operator is not supported.
		err = tsl.UnexpectedLiteralError{Literal: n.Func}
	}

	return
}

// Walk travel the TSL tree to create squirrel SQL select operators.
//
// Users can call the Walk method inside a squirrel Where to add the query.
//
//  filter, _ := sql.Walk(tree)
//  sql, args, _ := sq.Select("name, city, state").
//    From("users").
//    Where(filter).
//    ToSql()
//
// Squirrel: https://github.com/Masterminds/squirrel
//
func Walk(n tsl.Node) (s sq.Sqlizer, err error) {
	switch n.Func {
	case tsl.IdentOp:
		s = sq.Expr(n.Left.(string))
	case tsl.NumberOp:
		f := strconv.FormatFloat(n.Left.(float64), 'g', -1, 64)
		s = sq.Expr(f)
	case tsl.StringOp:
		s = sq.Expr(n.Left.(string))
	case tsl.AndOp, tsl.OrOp, tsl.AddOp, tsl.SubtractOp, tsl.MultiplyOp, tsl.DivideOp,
		tsl.ModuloOp:
		return binaryStep(n)
	case tsl.NotOp, tsl.EqOp, tsl.NotEqOp, tsl.LtOp, tsl.LteOp, tsl.GtOp, tsl.GteOp,
		tsl.InOp, tsl.NotInOp, tsl.IsNilOp, tsl.IsNotNilOp:
		return unaryStep(n)
	case tsl.LikeOp, tsl.ILikeOp, tsl.BetweenOp, tsl.NotBetweenOp:
		return unaryStep(n)
	default:
		// If here than the operator is not supported.
		err = tsl.UnexpectedLiteralError{Literal: n.Func}
	}

	return
}
