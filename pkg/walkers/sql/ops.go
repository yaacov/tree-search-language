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

package sql

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/yaacov/tree-search-language/pkg/tsl"
)

// Currently squirrel does not have a some needed operator expression.

// notExpr handles SQL not.
type notExpr []sq.Sqlizer

//nolint
func (n notExpr) ToSql() (sql string, args []interface{}, err error) {
	if len(n) != 1 {
		return "", nil, tsl.UnexpectedLiteralError{}
	}

	partSQL, partArgs, err := n[0].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, partArgs...)
	sql = fmt.Sprintf("NOT (%s)", partSQL)

	return
}

// mathExpToSQL take a math expresion and return it's string, args and error.
func mathExpToSQL(n []sq.Sqlizer, mathOp string) (sql string, args []interface{}, err error) {
	var left string
	var right string
	var partArgs []interface{}

	if len(n) != 2 {
		return "", nil, tsl.UnexpectedLiteralError{}
	}

	left, partArgs, err = n[0].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, partArgs...)

	right, partArgs, err = n[1].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, partArgs...)
	sql = fmt.Sprintf("(%s %s %s)", left, mathOp, right)

	return
}

// addExpr hanlhandlesdels SQL add.
type addExpr []sq.Sqlizer

//nolint
func (n addExpr) ToSql() (sql string, args []interface{}, err error) {
	return mathExpToSQL(n, "+")
}

// subExpr hanlhandlesdels SQL subtract.
type subExpr []sq.Sqlizer

//nolint
func (n subExpr) ToSql() (sql string, args []interface{}, err error) {
	return mathExpToSQL(n, "-")
}

// mulExpr hanlhandlesdels SQL multiply.
type mulExpr []sq.Sqlizer

//nolint
func (n mulExpr) ToSql() (sql string, args []interface{}, err error) {
	return mathExpToSQL(n, "*")
}

// divExpr hanlhandlesdels SQL divide.
type divExpr []sq.Sqlizer

//nolint
func (n divExpr) ToSql() (sql string, args []interface{}, err error) {
	return mathExpToSQL(n, "/")
}

// modExpr hanlhandlesdels SQL modulo.
type modExpr []sq.Sqlizer

//nolint
func (n modExpr) ToSql() (sql string, args []interface{}, err error) {
	return mathExpToSQL(n, "%")
}
