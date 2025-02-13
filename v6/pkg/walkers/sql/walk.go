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
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

// Walk travel the TSL tree to create squirrel SQL select operators.
//
// Users can call the Walk method inside a squirrel Where to add the query.
//
//	filter, _ := sql.Walk(tree)
//	sql, args, _ := sq.Select("name, city, state").
//	  From("users").
//	  Where(filter).
//	  ToSql()
//
// Squirrel: https://github.com/Masterminds/squirrel
func Walk(n *tsl.TSLNode) (s sq.Sqlizer, err error) {
	switch n.Type() {
	case tsl.KindIdentifier:
		s = sq.Expr(n.Value().(string))
	case tsl.KindNumericLiteral:
		s = sq.Expr("?", n.Value().(float64))
	case tsl.KindDateLiteral, tsl.KindTimestampLiteral:
		// Format time value using SQL timestamp format
		t := n.Value().(time.Time)
		s = sq.Expr("?", t.Format("2006-01-02 15:04:05"))
	case tsl.KindStringLiteral:
		s = sq.Expr("?", n.Value().(string))
	case tsl.KindBooleanLiteral:
		if n.Value().(bool) {
			s = sq.Expr("?", 1)
		} else {
			s = sq.Expr("?", 0)
		}
	case tsl.KindBinaryExpr:
		return binaryStep(n)
	case tsl.KindUnaryExpr:
		return unaryStep(n)
	case tsl.KindNullLiteral:
		// NULL literal is handled as a special case of IS NULL operator
		s = sq.Expr("")
	default:
		err = tsl.UnexpectedLiteralError{Literal: n.Type()}
	}

	return
}

// Helper function to walk array nodes and return values
func walkArrayValues(n *tsl.TSLNode) ([]sq.Sqlizer, error) {
	if n.Type() != tsl.KindArrayLiteral {
		return nil, tsl.UnexpectedTypeError{Type: n.Type()}
	}

	array := n.Value().(tsl.TSLArrayLiteral)
	values := make([]sq.Sqlizer, len(array.Values))
	var err error

	for i, node := range array.Values {
		values[i], err = Walk(node)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}

func binaryStep(n *tsl.TSLNode) (s sq.Sqlizer, err error) {
	var l sq.Sqlizer
	op := n.Value().(tsl.TSLExpressionOp)

	l, err = Walk(op.Left)
	if err != nil {
		return
	}

	// Handle array operations specially
	switch op.Operator {
	case tsl.OpIn:
		values, err := walkArrayValues(op.Right)
		if err != nil {
			return nil, err
		}
		return sq.Expr("? IN ("+placeholders(len(values))+")", append([]interface{}{l}, sqlizersToInterface(values)...)...), nil

	case tsl.OpBetween:
		values, err := walkArrayValues(op.Right)
		if err != nil {
			return nil, err
		}
		if len(values) != 2 {
			return nil, tsl.BetweenOperatorError{Message: "BETWEEN requires exactly two values"}
		}
		return sq.Expr("? BETWEEN ? AND ?", l, values[0], values[1]), nil
	}

	// For non-array operations, handle normally
	r, err := Walk(op.Right)
	if err != nil {
		return
	}

	switch op.Operator {
	// Arithmetic operators
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

	// Comparison operators
	case tsl.OpEQ:
		return sq.Expr("? = ?", l, r), nil
	case tsl.OpNE:
		return sq.Expr("? != ?", l, r), nil
	case tsl.OpLT:
		return sq.Expr("? < ?", l, r), nil
	case tsl.OpLE:
		return sq.Expr("? <= ?", l, r), nil
	case tsl.OpGT:
		return sq.Expr("? > ?", l, r), nil
	case tsl.OpGE:
		return sq.Expr("? >= ?", l, r), nil
	case tsl.OpREQ:
		return sq.Expr("? REGEXP ?", l, r), nil
	case tsl.OpRNE:
		return sq.Expr("NOT (? REGEXP ?)", l, r), nil

	// Logical operators
	case tsl.OpAnd:
		return sq.And{l, r}, nil
	case tsl.OpOr:
		return sq.Or{l, r}, nil

	// String operators
	case tsl.OpLike:
		return sq.Expr("? LIKE ?", l, r), nil
	case tsl.OpILike:
		return sq.Expr("? ILIKE ?", l, r), nil

	// Null operator
	case tsl.OpIs:
		return sq.Expr("? IS NULL", l), nil

	default:
		return nil, tsl.UnexpectedOperatorError{Operator: op.Operator}
	}
}

// unaryStep handles minus and not operators first
func unaryStep(n *tsl.TSLNode) (s sq.Sqlizer, err error) {
	op := n.Value().(tsl.TSLExpressionOp)

	// Get the child node's SQL representation
	right, err := Walk(op.Right)
	if err != nil {
		return nil, err
	}

	// Handle minus and not operators first
	switch op.Operator {
	case tsl.OpUMinus:
		return sq.Expr("-(?)", right), nil
	case tsl.OpNot:
		return sq.Expr("NOT (?)", right), nil
	default:
		return nil, tsl.UnexpectedLiteralError{Literal: op.Operator}
	}
}

// Helper to generate SQL placeholders
func placeholders(n int) string {
	if n <= 0 {
		return ""
	}
	ph := make([]byte, 2*n-1)
	for i := 0; i < len(ph); i += 2 {
		ph[i] = '?'
		if i+1 < len(ph) {
			ph[i+1] = ','
		}
	}
	return string(ph)
}

// Helper to convert []sq.Sqlizer to []interface{}
func sqlizersToInterface(sqlizers []sq.Sqlizer) []interface{} {
	result := make([]interface{}, len(sqlizers))
	for i, s := range sqlizers {
		result[i] = s
	}
	return result
}
