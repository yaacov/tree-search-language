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
		"tags":        []string{"fiction", "bestseller"},
		"price":       29.99,
	}

	// This is the evaluation function that we will use:
	eval := func(name string) (value interface{}, ok bool) {
		value, ok = record[name]
		return
	}

	DescribeTable("Returns the expected result",
		func(text string, expected bool) {
			// Parse the text:
			tree, err := tsl.ParseTSL(text)
			Expect(err).ToNot(HaveOccurred())
			defer tree.Free()

			// Get the value:
			actual, err := Walk(tree, eval)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(expected))
		},

		// String operations
		Entry("equals string", "title = 'A good book'", true),
		Entry("not equals string", "author != 'Jane'", true),
		Entry("like with wildcard", "title like '%good%'", true),
		Entry("ilike case insensitive", "title ilike '%GOOD%'", true),
		Entry("regexp equals", "title ~= 'good.*'", true),
		Entry("regexp not equals", "title !~ '.*bad.*'", true),

		// Numeric operations
		Entry("equals number", "spec.pages = 14", true),
		Entry("greater than", "spec.rating > 4", true),
		Entry("less than or equal", "spec.pages <= 14", true),
		Entry("between numbers", "price between 20 and 30", true),

		// Boolean operations
		Entry("equals boolean", "loaned = true", true),
		Entry("not equals boolean", "loaned != false", true),
		Entry("complex boolean", "(spec.pages > 10 and loaned = true) or spec.rating >= 5", true),

		// Date operations
		Entry("equals date", "date = '2020-01-01T00:00:00Z'", true),
		Entry("greater than date", "date > '2019-12-31T00:00:00Z'", true),
		Entry("between dates", "date between '2019-12-31T00:00:00Z' and '2020-01-02T00:00:00Z'", true),

		// Array operations
		Entry("in array literal", "spec.rating in [3, 4, 5]", true),
		Entry("not in array", "spec.pages in [20, 30, 40]", false),

		// Null checks
		Entry("is null", "missing_field is null", true),
		Entry("is not null", "title is not null", true),

		// Combined operations
		Entry("complex query", "(spec.pages <= spec.rating * 3) and (title like '%book%' or author = 'Joe')", true),
	)
})

var _ = Describe("Walk error cases", func() {
	eval := func(name string) (interface{}, bool) {
		return nil, false
	}

	DescribeTable("Returns appropriate errors",
		func(text string, expectedError string) {
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
			Expect(err.Error()).To(ContainSubstring(expectedError))
		},

		Entry("Type mismatch string/number", "name > 5", "type mismatch"),
		Entry("Type mismatch number/string", "age = 'young'", "type mismatch"),
		Entry("Invalid operator", "name invalid 'test'", "unexpected operator"),
		Entry("Invalid between array", "age between [1]", "between operator requires exactly 2 values"),
	)
})

var _ = Describe("Walk", func() {
	Context("Valid expressions", func() {
		It("Accepts valid arithmetic", func() {
			tree, err := tsl.ParseTSL("price * 2 > 1000")
			Expect(err).ToNot(HaveOccurred())
			err = validateTree(tree)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Accepts valid comparisons", func() {
			tree, err := tsl.ParseTSL("name = 'john'")
			Expect(err).ToNot(HaveOccurred())
			err = validateTree(tree)
			Expect(err).ToNot(HaveOccurred())
		})

		It("Accepts valid logical operations", func() {
			tree, err := tsl.ParseTSL("age > 20 and city = 'London'")
			Expect(err).ToNot(HaveOccurred())
			err = validateTree(tree)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("Invalid expressions", func() {
		It("Rejects invalid arithmetic operands", func() {
			tree, err := tsl.ParseTSL("'hello' + 'world'")
			Expect(err).ToNot(HaveOccurred())
			err = validateTree(tree)
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(TypeMismatchError{}))
		})

		It("Rejects invalid comparison operands", func() {
			tree, err := tsl.ParseTSL("5 = name")
			Expect(err).ToNot(HaveOccurred())
			err = validateTree(tree)
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(TypeMismatchError{}))
		})

		It("Rejects invalid LIKE operands", func() {
			tree, err := tsl.ParseTSL("5 like '5%'")
			Expect(err).ToNot(HaveOccurred())
			err = validateTree(tree)
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(TypeMismatchError{}))
		})
	})
})
