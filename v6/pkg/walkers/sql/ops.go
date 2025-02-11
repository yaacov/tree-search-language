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
	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

// Currently squirrel does not have a some needed operator expression.

// notExpr handles SQL not.
type notExpr []sq.Sqlizer

// nolint
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

// addExpr handles SQL addition
type addExpr []sq.Sqlizer

func (a addExpr) ToSql() (sql string, args []interface{}, err error) {
	if len(a) != 2 {
		return "", nil, tsl.UnexpectedLiteralError{}
	}

	leftSQL, leftArgs, err := a[0].ToSql()
	if err != nil {
		return "", nil, err
	}
	rightSQL, rightArgs, err := a[1].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, leftArgs...)
	args = append(args, rightArgs...)
	sql = fmt.Sprintf("(%s + %s)", leftSQL, rightSQL)

	return
}

// subExpr handles SQL subtraction
type subExpr []sq.Sqlizer

func (s subExpr) ToSql() (sql string, args []interface{}, err error) {
	if len(s) != 2 {
		return "", nil, tsl.UnexpectedLiteralError{}
	}

	leftSQL, leftArgs, err := s[0].ToSql()
	if err != nil {
		return "", nil, err
	}
	rightSQL, rightArgs, err := s[1].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, leftArgs...)
	args = append(args, rightArgs...)
	sql = fmt.Sprintf("(%s - %s)", leftSQL, rightSQL)

	return
}

// mulExpr handles SQL multiplication
type mulExpr []sq.Sqlizer

func (m mulExpr) ToSql() (sql string, args []interface{}, err error) {
	if len(m) != 2 {
		return "", nil, tsl.UnexpectedLiteralError{}
	}

	leftSQL, leftArgs, err := m[0].ToSql()
	if err != nil {
		return "", nil, err
	}
	rightSQL, rightArgs, err := m[1].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, leftArgs...)
	args = append(args, rightArgs...)
	sql = fmt.Sprintf("(%s * %s)", leftSQL, rightSQL)

	return
}

// divExpr handles SQL division
type divExpr []sq.Sqlizer

func (d divExpr) ToSql() (sql string, args []interface{}, err error) {
	if len(d) != 2 {
		return "", nil, tsl.UnexpectedLiteralError{}
	}

	leftSQL, leftArgs, err := d[0].ToSql()
	if err != nil {
		return "", nil, err
	}
	rightSQL, rightArgs, err := d[1].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, leftArgs...)
	args = append(args, rightArgs...)
	sql = fmt.Sprintf("(%s / %s)", leftSQL, rightSQL)

	return
}

// modExpr handles SQL modulo
type modExpr []sq.Sqlizer

func (m modExpr) ToSql() (sql string, args []interface{}, err error) {
	if len(m) != 2 {
		return "", nil, tsl.UnexpectedLiteralError{}
	}

	leftSQL, leftArgs, err := m[0].ToSql()
	if err != nil {
		return "", nil, err
	}
	rightSQL, rightArgs, err := m[1].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, leftArgs...)
	args = append(args, rightArgs...)
	sql = fmt.Sprintf("(%s %% %s)", leftSQL, rightSQL)

	return
}
