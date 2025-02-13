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

// Example for the tsl package.
func Example() {
	// Set a TSL input string.
	input := "name = 'joe' and city != 'rome'"

	// Parse input string into a TSL tree.
	tree, _ := tsl.ParseTSL(input)

	// Set filter
	filter, _ := Walk(tree)

	// Convert TSL tree into SQL string using squirrel sql builder.
	sql, args, _ := sq.Select("name, city, state").
		From("users").
		Where(filter).
		ToSql()

	fmt.Printf("SQL : %s\n", sql)
	fmt.Printf("Args: %v\n", args)

	// Output:
	// SQL : SELECT name, city, state FROM users WHERE (name = ? AND city != ?)
	// Args: [joe rome]
}

func Example_complex() {
	// Complex query with multiple conditions
	input := `
		(salary * 12) > 50000 AND 
		department IN ['IT', 'HR'] AND 
		hire_date BETWEEN 2020-01-01T00:00:00Z and 2023-12-31T23:59:59Z AND
		(manager IS NULL OR title LIKE '%Senior%')
	`

	// Parse input string into a TSL tree
	tree, _ := tsl.ParseTSL(input)

	// Convert to SQL
	filter, _ := Walk(tree)
	sql, args, _ := sq.Select("name, department, salary").
		From("employees").
		Where(filter).
		ToSql()

	fmt.Printf("SQL : %s\n", sql)
	fmt.Printf("Args: %v\n", args)

	// Output:
	// SQL : SELECT name, department, salary FROM employees WHERE ((((salary * ?) > ? AND department IN (?,?)) AND hire_date BETWEEN ? AND ?) AND (manager IS NULL OR title LIKE ?))
	// Args: [12 50000 IT HR 2020-01-01 00:00:00 2023-12-31 23:59:59 %Senior%]
}

func Example_arithmetic() {
	input := "(base_salary + bonus) * tax_rate > 20000"
	tree, _ := tsl.ParseTSL(input)
	filter, _ := Walk(tree)
	sql, args, _ := sq.Select("*").From("salaries").Where(filter).ToSql()

	fmt.Printf("SQL : %s\n", sql)
	fmt.Printf("Args: %v\n", args)

	// Output:
	// SQL : SELECT * FROM salaries WHERE ((base_salary + bonus) * tax_rate) > ?
	// Args: [20000]
}
