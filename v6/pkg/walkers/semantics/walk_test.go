// Copyright 2019 Yaacov Zamir <kobi.zamir@gmail.com>
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

package semantics

import (
	"strings"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

func TestWalk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Semantic walker")
}

var _ = Describe("Walk", func() {
	// This is the record that we will use for all the tests:
	date, _ := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")
	record := map[string]interface{}{
		"title":       "A good book",
		"author":      "Joe",
		"spec.pages":  14.0,
		"spec.rating": 5.0,
		"loaned":      true,
		"date":        date,
		"tags":        []interface{}{"fiction", "bestseller"},
		"price":       29.99,
		"numbers":     []interface{}{1.0, 2.0, 3.0},
		"booleans":    []interface{}{true, false, true},
	}

	// This is the evaluation function that we will use:
	eval := func(name string) (value interface{}, ok bool) {
		value, ok = record[name]
		return
	}

	DescribeTable("Returns the expected result",
		func(text string, expected interface{}) {
			// Parse the text:
			tree, err := tsl.ParseTSL(text)
			Expect(err).ToNot(HaveOccurred())
			defer tree.Free()

			// Get the value:
			actual, err := Walk(tree, eval)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(expected))
		},

		// Direct identifier evaluation
		Entry("string identifier", "title", "A good book"),
		Entry("numeric identifier", "spec.pages", 14.0),
		Entry("boolean identifier", "loaned", true),
		Entry("date identifier", "date", date),
		Entry("price identifier", "price", 29.99),
		Entry("numeric array identifier", "numbers", []interface{}{1.0, 2.0, 3.0}),
		Entry("string array identifier", "tags", []interface{}{"fiction", "bestseller"}),
		Entry("boolean array identifier", "booleans", []interface{}{true, false, true}),

		// String operations
		Entry("equals string", "title = 'A good book'", true),
		Entry("not equals string", "author != 'Jane'", true),
		Entry("like with wildcard", "title like '%good%'", true),
		Entry("ilike case insensitive", "title ilike '%GOOD%'", true),
		Entry("regexp equals", "title ~= 'good.*'", true),
		Entry("regexp not equals", "title ~! '.*bad.*'", true),

		// Numeric operations
		Entry("equals number", "spec.pages = 14", true),
		Entry("greater than", "spec.rating > 4", true),
		Entry("less than or equal", "spec.pages <= 14", true),
		Entry("between numbers", "price between 20 and 30", true),
		Entry("addition", "spec.pages + 1 = 15", true),
		Entry("subtraction", "spec.pages - 4 = 10", true),
		Entry("multiplication", "spec.rating * 2 = 10", true),
		Entry("division", "price / 2 = 14.995", true),
		Entry("modulus", "spec.pages % 5 = 4", true),

		// Unary minus operations
		Entry("unary minus less than", "-spec.rating < 0", true),
		Entry("unary minus greater than", "-spec.rating > -10", true),
		Entry("unary minus equals", "-spec.rating = -5", true),
		Entry("unary minus with expression", "-(spec.rating - 10) = 5", true),
		Entry("unary minus in complex comparison", "-spec.pages < -spec.rating", true),
		Entry("unary minus with addition", "(-spec.rating + 10) > 0", true),

		// Boolean operations
		Entry("equals boolean", "loaned = true", true),
		Entry("not equals boolean", "loaned != false", true),
		Entry("complex boolean", "(spec.pages > 10 and loaned = true) or spec.rating >= 5", true),
		Entry("boolean and", "loaned and spec.rating > 4", true),
		Entry("boolean or", "loaned or spec.rating < 4", true),

		// Date operations
		Entry("equals date", "date = '2020-01-01T00:00:00Z'", true),
		Entry("greater than date", "date > '2019-12-31T00:00:00Z'", true),
		Entry("between dates", "date between '2019-12-31T00:00:00Z' and '2020-01-02T00:00:00Z'", true),
		Entry("date before", "date < '2021-01-01T00:00:00Z'", true),
		Entry("date after", "date > '2019-01-01T00:00:00Z'", true),

		// Array operations
		Entry("in array literal", "spec.rating in [3, 4, 5]", true),
		Entry("not in array", "spec.pages in [20, 30, 40]", false),
		Entry("in array identifier", "2 in numbers", true),
		Entry("not in array identifier", "5 in numbers", false),
		Entry("in array expression", "5 in (numbers + 2)", true),
		Entry("complex in array", "spec.rating in (numbers + 2)", true),

		// Null checks
		Entry("is null", "price is null", false),
		Entry("is not null", "title is not null", true),

		// Combined operations
		Entry("complex query", "(spec.pages <= spec.rating * 3) and (title like '%book%' or author = 'Joe')", true),
		Entry("nested query", "(spec.pages > 10 and (loaned = true or spec.rating >= 5))", true),

		// Array operations with unary operators
		Entry("unary minus on identifier array", "-numbers", []interface{}{-1.0, -2.0, -3.0}),
		Entry("not on identifier boolean array", "not booleans", []interface{}{false, true, false}),

		// Binary operations on arrays
		Entry("array plus number", "numbers + 10", []interface{}{11.0, 12.0, 13.0}),
		Entry("array minus number", "numbers - 1", []interface{}{0.0, 1.0, 2.0}),
		Entry("array multiplied by number", "numbers * 2", []interface{}{2.0, 4.0, 6.0}),
		Entry("array divided by number", "numbers / 2", []interface{}{0.5, 1.0, 1.5}),
		Entry("array modulus", "numbers % 2", []interface{}{1.0, 0.0, 1.0}),
		Entry("array equality", "numbers = 2", []interface{}{false, true, false}),
		Entry("array inequality", "numbers != 2", []interface{}{true, false, true}),
		Entry("array greater than", "numbers > 1", []interface{}{false, true, true}),
		Entry("array less than", "numbers < 3", []interface{}{true, true, false}),
		Entry("array less than or equal", "numbers <= 2", []interface{}{true, true, false}),
		Entry("array greater than or equal", "numbers >= 2", []interface{}{false, true, true}),
		Entry("boolean array and boolean", "booleans and true", []interface{}{true, false, true}),
		Entry("boolean array or boolean", "booleans or false", []interface{}{true, false, true}),

		// String array operations
		Entry("string array equality", "tags = 'fiction'", []interface{}{true, false}),
		Entry("string array inequality", "tags != 'fiction'", []interface{}{false, true}),
		Entry("string array like", "tags like 'fic%'", []interface{}{true, false}),
		Entry("string array ilike", "tags ilike 'FIC%'", []interface{}{true, false}),
		Entry("string array regex equals", "tags ~= '^f.*'", []interface{}{true, false}),
		Entry("string array regex not equals", "tags ~! '^b.*'", []interface{}{true, false}),
		Entry("any with string array like", "any (tags like '%sell%')", true),
		Entry("all with string array like", "all (tags like '%i%')", false),
		Entry("literal string array operations", "['fiction', 'nonfiction', 'bestseller'] = 'bestseller'", []interface{}{false, false, true}),
		Entry("literal string array like", "['abc', 'def', 'ghi'] like '%e%'", []interface{}{false, true, false}),

		// Literal array operations
		Entry("literal array addition", "[1, 2, 3] + 4", []interface{}{5.0, 6.0, 7.0}),
		Entry("literal array subtraction", "[5, 6, 7] - 2", []interface{}{3.0, 4.0, 5.0}),
		Entry("literal array multiplication", "[2, 3, 4] * 3", []interface{}{6.0, 9.0, 12.0}),
		Entry("literal array division", "[10, 20, 30] / 10", []interface{}{1.0, 2.0, 3.0}),
		Entry("literal array modulus", "[10, 11, 12] % 3", []interface{}{1.0, 2.0, 0.0}),
		Entry("literal array comparison", "[1, 5, 10] > 4", []interface{}{false, true, true}),
		Entry("literal array equals", "[1, 2, 3] = 2", []interface{}{false, true, false}),
		Entry("literal array not equals", "[1, 2, 3] != 2", []interface{}{true, false, true}),

		// Nested arrays and complex operations
		Entry("nested array operations", "([1, 2, 3] + 1) * 2", []interface{}{4.0, 6.0, 8.0}),

		// Array operators ANY, ALL, LEN
		Entry("any with array identifier", "any (numbers > 1)", true),
		Entry("any with array identifier (false case)", "any (numbers > 5)", false),
		Entry("all with array identifier (true case)", "all (numbers > 0)", true),
		Entry("all with array identifier (false case)", "all (numbers > 1)", false),
		Entry("len with array identifier", "len numbers", 3.0),
		Entry("len with literal array", "len [1, 2, 3, 4, 5]", 5.0),
		Entry("any with boolean array", "any booleans", true),
		Entry("all with boolean array", "all booleans", false),
		Entry("len in comparison", "len numbers = 3", true),
		Entry("any with complex expression", "any (numbers * 2 > 5)", true),
		Entry("all with complex expression", "all ((numbers + 10) > 10)", true),
		Entry("any with string array identifier", "any (tags = 'fiction')", true),
	)
})

var _ = Describe("Walk error cases", func() {
	eval := func(name string) (interface{}, bool) {
		return nil, false
	}

	DescribeTable("Returns appropriate errors",
		func(text string, expectedError interface{}) {
			tree, err := tsl.ParseTSL(text)
			if err != nil {
				// Skip syntax errors as they are expected for invalid operators
				if strings.Contains(err.Error(), "syntax error") {
					return
				}
			}
			Expect(err).ToNot(HaveOccurred())
			defer tree.Free()

			_, err = Walk(tree, eval)
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(expectedError))
		},

		Entry("is null", "missing_field is null", tsl.KeyNotFoundError{}),
		Entry("Type mismatch string/number", "name > 5", tsl.KeyNotFoundError{}),
		Entry("Type mismatch number/string", "age = 'young'", tsl.KeyNotFoundError{}),
		Entry("Invalid operator", "name invalid 'test'", tsl.UnexpectedOperatorError{}),
		Entry("Invalid between array", "age between [1]", tsl.BetweenOperatorError{}),
	)
})
