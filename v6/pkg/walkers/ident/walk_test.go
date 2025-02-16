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

package ident

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

func TestWalk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ident walker")
}

var _ = Describe("Walk", func() {
	var columnMap map[string]string

	BeforeEach(func() {
		columnMap = map[string]string{
			"name":        "user_name",
			"age":         "user_age",
			"city":        "address_city",
			"country":     "address_country",
			"salary":      "emp_salary",
			"bonus":       "emp_bonus",
			"spec.pages":  "pages",
			"spec.rating": "rating",
		}
	})

	check := func(s string) (string, error) {
		if v, ok := columnMap[s]; ok {
			return v, nil
		}
		return s, fmt.Errorf("column not found: %s", s)
	}

	It("Should extract and replace identifiers in simple expression", func() {
		tree, err := tsl.ParseTSL("name = 'john'")
		Expect(err).ToNot(HaveOccurred())
		defer tree.Free()

		newTree, err := Walk(tree, check)
		Expect(err).ToNot(HaveOccurred())
		defer newTree.Free()

		// Verify the replacement
		Expect(newTree.Type()).To(Equal(tsl.KindBinaryExpr))
		op := newTree.Value().(tsl.TSLExpressionOp)
		Expect(op.Left.Value()).To(Equal("user_name"))
	})

	It("Should extract and replace multiple identifiers", func() {
		tree, err := tsl.ParseTSL("age > 20 and city = 'London'")
		Expect(err).ToNot(HaveOccurred())
		defer tree.Free()

		newTree, err := Walk(tree, check)
		Expect(err).ToNot(HaveOccurred())
		defer newTree.Free()

		// Verify the replacement in the new tree
		op := newTree.Value().(tsl.TSLExpressionOp)
		leftOp := op.Left.Value().(tsl.TSLExpressionOp)
		rightOp := op.Right.Value().(tsl.TSLExpressionOp)
		Expect(leftOp.Left.Value()).To(Equal("user_age"))
		Expect(rightOp.Left.Value()).To(Equal("address_city"))
	})

	It("Should handle complex nested expressions", func() {
		tree, err := tsl.ParseTSL("(name = 'john' or age > 20) and (city = 'London' or country = 'UK')")
		Expect(err).ToNot(HaveOccurred())
		defer tree.Free()

		newTree, err := Walk(tree, check)
		Expect(err).ToNot(HaveOccurred())
		defer newTree.Free()
		// Tree structure should be preserved, just identifiers replaced
		Expect(newTree.Type()).To(Equal(tsl.KindBinaryExpr))
	})

	It("Should return error for unknown identifiers", func() {
		tree, err := tsl.ParseTSL("unknown_column > 100")
		Expect(err).ToNot(HaveOccurred())
		defer tree.Free()

		newTree, err := Walk(tree, check)
		if err == nil {
			defer newTree.Free()
		}
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("column not found: unknown_column"))
	})

	It("Should handle nested field names", func() {
		tree, err := tsl.ParseTSL("spec.pages > 100 and spec.rating = 5")
		Expect(err).ToNot(HaveOccurred())
		defer tree.Free()

		newTree, err := Walk(tree, check)
		Expect(err).ToNot(HaveOccurred())
		defer newTree.Free()

		// Verify the replacements
		op := newTree.Value().(tsl.TSLExpressionOp)
		leftOp := op.Left.Value().(tsl.TSLExpressionOp)
		rightOp := op.Right.Value().(tsl.TSLExpressionOp)
		Expect(leftOp.Left.Value()).To(Equal("pages"))
		Expect(rightOp.Left.Value()).To(Equal("rating"))
	})
})
