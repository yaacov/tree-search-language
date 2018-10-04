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

// Currently squirrel does not have a NOT operator expression,
// this expresson handles SQL not.
type notExpr []sq.Sqlizer

//nolint
func (n notExpr) ToSql() (sql string, args []interface{}, err error) {
	if len(n) != 1 {
		return "", nil, fmt.Errorf("operator not does not have one argument")
	}

	partSQL, partArgs, err := n[0].ToSql()
	if err != nil {
		return "", nil, err
	}

	args = append(args, partArgs...)
	sql = fmt.Sprintf("NOT (%s)", partSQL)

	return
}

// Currently squirrel does not have a ADD operator expression,
// this expresson hanlhandlesdels SQL add.
type addExpr []sq.Sqlizer

//nolint
func (n addExpr) ToSql() (sql string, args []interface{}, err error) {
	var left string
	var right string
	var partArgs []interface{}

	if len(n) != 2 {
		return "", nil, fmt.Errorf("operator not does not have one argument")
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
	sql = fmt.Sprintf("%s + %s", left, right)

	return
}
