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
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	sq "github.com/Masterminds/squirrel"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

func TestSQLWalker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SQL walker")
}

var _ = Describe("Walk", func() {
	DescribeTable("Generates the expected SQL and arguments",
		func(input string, expectedSQL string, expectedArgs ...interface{}) {
			// Parse the input:
			tree, err := tsl.ParseTSL(input)
			Expect(err).ToNot(HaveOccurred())

			// Convert the tree to a SQL filter:
			filter, err := Walk(tree)
			Expect(err).ToNot(HaveOccurred())

			// Convert the filter to SQL text:
			actualSQL, actualArgs, err := sq.Select("name, city, state").
				From("users").
				Where(filter).
				ToSql()
			Expect(err).ToNot(HaveOccurred())

			// Ensure that the arguments are not nil:
			if actualArgs == nil {
				actualArgs = []interface{}{}
			}

			// Check that the generated SQL and the arguments are the expected ones:
			Expect(actualSQL).To(Equal(expectedSQL))
			Expect(actualArgs).To(Equal(expectedArgs))
		},

		Entry(
			"Search by name and city",
			"name = 'joe' and city != 'rome'",
			"SELECT name, city, state FROM users WHERE (name = ? AND city != ?)",
			"joe", "rome",
		),

		Entry(
			"Boolean literal",
			"name = 'joe' and isCarpenter = TRUE",
			"SELECT name, city, state FROM users WHERE (name = ? AND isCarpenter = ?)",
			"joe", 1,
		),

		Entry(
			"Date literal",
			"name = 'joe' and date = 2020-01-01T00:00:01Z",
			"SELECT name, city, state FROM users WHERE (name = ? AND date = ?)",
			"joe", "2020-01-01 00:00:01",
		),

		Entry(
			"Addition",
			"salary + bonus > 50000",
			"SELECT name, city, state FROM users WHERE (salary + bonus) > ?",
			50000.0,
		),

		Entry(
			"Multiplication",
			"hours * rate = 1000",
			"SELECT name, city, state FROM users WHERE (hours * rate) = ?",
			1000.0,
		),

		Entry(
			"Complex arithmetic",
			"(salary + bonus) * 0.3 > 20000",
			"SELECT name, city, state FROM users WHERE ((salary + bonus) * ?) > ?",
			0.3, 20000.0,
		),

		Entry(
			"IN operator",
			"city IN ['rome', 'paris', 'london']",
			"SELECT name, city, state FROM users WHERE city IN (?,?,?)",
			"rome", "paris", "london",
		),

		Entry(
			"BETWEEN operator",
			"age BETWEEN 20 and 30",
			"SELECT name, city, state FROM users WHERE age BETWEEN ? AND ?",
			20.0, 30.0,
		),

		Entry(
			"NULL check",
			"email IS NULL",
			"SELECT name, city, state FROM users WHERE email IS NULL",
		),

		Entry(
			"LIKE operator",
			"name LIKE '%smith%'",
			"SELECT name, city, state FROM users WHERE name LIKE ?",
			"%smith%",
		),

		Entry(
			"ILIKE operator",
			"name ILIKE '%smith%'",
			"SELECT name, city, state FROM users WHERE name ILIKE ?",
			"%smith%",
		),

		Entry(
			"Regular expression",
			"email ~= '.*@gmail.com'",
			"SELECT name, city, state FROM users WHERE email REGEXP ?",
			".*@gmail.com",
		),

		Entry(
			"Regular expression negation",
			"email ~! '.*@gmail.com'",
			"SELECT name, city, state FROM users WHERE NOT (email REGEXP ?)",
			".*@gmail.com",
		),

		Entry(
			"Complex arithmetic",
			"(salary * 12) + bonus BETWEEN 50000 and 100000",
			"SELECT name, city, state FROM users WHERE ((salary * ?) + bonus) BETWEEN ? AND ?",
			12.0, 50000.0, 100000.0,
		),
	)
})
