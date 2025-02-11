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
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

// Update nodesToStrings to handle null values and edge cases
func nodesToStrings(in interface{}) (s []interface{}) {
	// Handle nil case
	if in == nil {
		return nil
	}

	// Handle array literals
	if arr, ok := in.(tsl.TSLArrayLiteral); ok {
		for _, node := range arr.Values {
			switch node.Type() {
			case tsl.KindBooleanLiteral:
				v := 0
				if node.Value().(bool) {
					v = 1
				}
				s = append(s, v)
			case tsl.KindDateLiteral, tsl.KindTimestampLiteral:
				d := node.Value().(time.Time).Format(time.RFC3339)
				s = append(s, d)
			default:
				s = append(s, node.Value())
			}
		}
		return
	}

	// Handle single values
	switch v := in.(type) {
	case bool:
		if v {
			return []interface{}{1}
		}
		return []interface{}{0}
	case time.Time:
		return []interface{}{v.Format(time.RFC3339)}
	default:
		return []interface{}{v}
	}
}

// Walk travel the TSL tree to create squirrel SQL select operators.
func Walk(n *tsl.TSLNode) (s sq.Sqlizer, err error) {
	switch n.Type() {
	case tsl.KindIdentifier:
		s = sq.Expr(n.Value().(string))
	case tsl.KindNumericLiteral:
		f := strconv.FormatFloat(n.Value().(float64), 'g', -1, 64)
		s = sq.Expr(f)
	case tsl.KindDateLiteral, tsl.KindTimestampLiteral:
		f := n.Value().(time.Time).Format(time.RFC3339)
		s = sq.Expr(f)
	case tsl.KindStringLiteral:
		s = sq.Expr(n.Value().(string))
	case tsl.KindBooleanLiteral:
		value := "0"
		if n.Value().(bool) {
			value = "1"
		}
		s = sq.Expr(value)
	case tsl.KindBinaryExpr:
		op := n.Value().(tsl.TSLExpressionOp)
		switch op.Operator {
		case tsl.OpAnd:
			return binaryStep(op.Left, op.Right, tsl.OpAnd)
		case tsl.OpOr:
			return binaryStep(op.Left, op.Right, tsl.OpOr)
		case tsl.OpPlus, tsl.OpMinus, tsl.OpStar, tsl.OpSlash, tsl.OpPercent:
			l, err := Walk(op.Left)
			if err != nil {
				return nil, err
			}
			r, err := Walk(op.Right)
			if err != nil {
				return nil, err
			}
			switch op.Operator {
			case tsl.OpPlus:
				return sq.Expr("(? + ?)", l, r), nil
			case tsl.OpMinus:
				return sq.Expr("(? - ?)", l, r), nil
			case tsl.OpStar:
				return sq.Expr("(? * ?)", l, r), nil
			case tsl.OpSlash:
				return sq.Expr("(? / ?)", l, r), nil
			case tsl.OpPercent:
				return sq.Expr("(? % ?)", l, r), nil
			}
		default:
			return unaryStep(n)
		}
	case tsl.KindUnaryExpr:
		return unaryStep(n)
	default:
		err = tsl.UnexpectedLiteralError{Literal: n.Type()}
	}

	return
}

// Update binaryStep to handle the new node types
func binaryStep(left, right *tsl.TSLNode, op tsl.Operator) (s sq.Sqlizer, err error) {
	var l, r sq.Sqlizer

	l, err = Walk(left)
	if err != nil {
		return
	}

	r, err = Walk(right)
	if err != nil {
		return
	}

	switch op {
	case tsl.OpAnd:
		s = sq.And{l, r}
	case tsl.OpOr:
		s = sq.Or{l, r}
	case tsl.OpPlus:
		s = addExpr{l, r}
	case tsl.OpMinus:
		s = subExpr{l, r}
	case tsl.OpStar:
		s = mulExpr{l, r}
	case tsl.OpSlash:
		s = divExpr{l, r}
	case tsl.OpPercent:
		s = modExpr{l, r}
	default:
		err = tsl.UnexpectedLiteralError{Literal: op}
	}

	return
}

// Update unaryStep to properly handle right node values
func unaryStep(n *tsl.TSLNode) (s sq.Sqlizer, err error) {
	op := n.Value().(tsl.TSLExpressionOp)
	var sql string
	var l sq.Sqlizer

	// Get the node's SQL representation
	l, err = Walk(op.Left)
	if err != nil {
		return
	}

	sql, _, err = l.ToSql()
	if err != nil {
		return
	}

	// Handle the NOT operator separately as it doesn't need right values
	if op.Operator == tsl.OpNot {
		return notExpr{l}, nil
	}

	// Convert the right node value to strings
	right := nodesToStrings(op.Right.Value())

	// Handle IS NULL case specially
	if op.Operator == tsl.OpIs {
		if len(right) == 0 {
			return sq.Eq{sql: nil}, nil
		}
		return sq.NotEq{sql: nil}, nil
	}

	// For operators that require values, ensure we have them
	if len(right) == 0 {
		return nil, tsl.UnexpectedLiteralError{Literal: "empty right operand"}
	}

	switch op.Operator {
	case tsl.OpEQ:
		return sq.Eq{sql: right[0]}, nil
	case tsl.OpNE:
		return sq.NotEq{sql: right[0]}, nil
	case tsl.OpLT:
		return sq.Lt{sql: right[0]}, nil
	case tsl.OpLE:
		return sq.LtOrEq{sql: right[0]}, nil
	case tsl.OpGT:
		return sq.Gt{sql: right[0]}, nil
	case tsl.OpGE:
		return sq.GtOrEq{sql: right[0]}, nil
	case tsl.OpIn:
		return sq.Eq{sql: right}, nil
	case tsl.OpLike:
		t := fmt.Sprintf("%s LIKE ?", sql)
		return sq.Expr(t, right[0]), nil
	case tsl.OpILike:
		t := fmt.Sprintf("%s ILIKE ?", sql)
		return sq.Expr(t, right[0]), nil
	case tsl.OpBetween:
		if len(right) < 2 {
			return nil, tsl.UnexpectedLiteralError{Literal: "BETWEEN requires two values"}
		}
		t := fmt.Sprintf("%s BETWEEN ? AND ?", sql)
		return sq.Expr(t, right[0], right[1]), nil
	default:
		return nil, tsl.UnexpectedLiteralError{Literal: op.Operator}
	}
}
