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

			// Check that the generated SQL and the arguments are the expected ones:
			Expect(actualSQL).To(Equal(expectedSQL))
			Expect(actualArgs).To(Equal(expectedArgs))
		},

		Entry(
			"Search by name and city",
			"name = 'joe' and city != 'rome'",
			"SELECT name, city, state FROM users WHERE (name = ? AND city <> ?)",
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
			"joe", "2020-01-01T00:00:01Z",
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
			"SELECT name, city, state FROM users WHERE ((salary + bonus) * 0.3) > ?",
			20000.0,
		),
	)
})
