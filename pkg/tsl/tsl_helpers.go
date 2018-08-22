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
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

// ToSQL converts a TSL tree into SQL statement.
func ToSQL(n Node, table string, usePgsql bool) (sql string, args []interface{}, err error) {
	var s sq.SelectBuilder

	if usePgsql {
		// If we are using PostgreSQL style use $ instead of ?
		// for placeholders.
		psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
		s = psql.Select("*").From(table)
	} else {
		s = sq.Select("*").From(table)
	}

	s = s.Where(walk(n))

	return s.ToSql()
}

// GetTree return the parsed tree, if exist.
func (l *Listener) GetTree() (n Node, err error) {
	// Check stack size.
	if len(l.Stack) < 1 {
		err = fmt.Errorf("operator stack is empty")
		return
	}

	return l.Stack[0], l.Err
}

// ternaryOp return lh if conditional is true, rh o/w.
func ternaryOp(conditional bool, lh string, rh string) string {
	if conditional {
		return lh
	}

	return rh
}

// literalValueToArg clean literal value, and return an arg string or float.
func literalValueToArg(v string) (arg interface{}) {
	l := len(v)

	// Check string length
	if l < 1 {
		return
	}

	// Check for ''
	if v[0] == '\'' {
		v = v[1 : l-1]
		arg = strings.Replace(v, "''", "'", -1)
		return
	}

	// It's a float
	arg, _ = strconv.ParseFloat(v, 64)

	return
}

// literalValuesToArgs collect literal values, and create args list.
func literalValuesToArgs(c hasLiteralValues) (args []interface{}) {
	// Get length of literal values list.
	l := len(c.AllLiteralValue())

	// Create the arg list.
	args = make([]interface{}, l)
	for i := 0; i < l; i++ {
		args[i] = literalValueToArg(c.LiteralValue(i).GetText())
	}

	return
}

func (l *Listener) push(i Node) {
	l.Stack = append(l.Stack, i)
}

func (l *Listener) pop() (n Node) {
	// Check that we have nodes in the stack.
	size := len(l.Stack)
	if size < 1 {
		l.Err = fmt.Errorf("operator stack is empty")
		return
	}

	// Pop the last value from the Stack.
	n, l.Stack = l.Stack[size-1], l.Stack[:size-1]

	return
}

func walk(n Node) sq.Sqlizer {
	switch n.Func {
	case andOp:
		return sq.And{walk(n.Right.(Node)), walk(n.Left.(Node))}
	case orOp:
		return sq.Or{walk(n.Right.(Node)), walk(n.Left.(Node))}
	case notOp:
		return sq.Or{walk(n.Right.(Node)), walk(n.Left.(Node))}
	case eqOp:
		return sq.Eq{n.Left.(string): n.Right}
	case notEqOp:
		return sq.NotEq{n.Left.(string): n.Right}
	case ltOp:
		return sq.Lt{n.Left.(string): n.Right}
	case lteOp:
		return sq.LtOrEq{n.Left.(string): n.Right}
	case gtOp:
		return sq.Gt{n.Left.(string): n.Right}
	case gteOp:
		return sq.GtOrEq{n.Left.(string): n.Right}
	case inOp:
		return sq.Eq{n.Left.(string): n.Right}
	case notInOp:
		return sq.NotEq{n.Left.(string): n.Right}
	case isNilOp:
		return sq.Eq{n.Left.(string): nil}
	case isNotNilOp:
		return sq.NotEq{n.Left.(string): nil}
	case likeOp:
		t := fmt.Sprintf("%s LIKE ?", n.Left.(string))
		return sq.Expr(t, n.Right)
	case notLikeOp:
		t := fmt.Sprintf("%s NOT LIKE ?", n.Left.(string))
		return sq.Expr(t, n.Right)
	case betweenOp:
		t := fmt.Sprintf("%s BETWEEN ? AND ?", n.Left.(string))
		return sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1])
	case notBetweenOp:
		t := fmt.Sprintf("%s NOT BETWEEN ? AND ?", n.Left.(string))
		return sq.Expr(t, n.Right.([]interface{})[0], n.Right.([]interface{})[1])
	}

	return sq.And{}
}
