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
	"sort"
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
	It("Should extract identifiers from simple expression", func() {
		tree, err := tsl.ParseTSL("name = 'john'")
		Expect(err).ToNot(HaveOccurred())

		idents, err := Walk(tree)
		Expect(err).ToNot(HaveOccurred())
		Expect(idents).To(HaveLen(1))
		Expect(idents).To(ContainElement("name"))
	})

	It("Should extract multiple identifiers", func() {
		tree, err := tsl.ParseTSL("age > 20 and city = 'London'")
		Expect(err).ToNot(HaveOccurred())

		idents, err := Walk(tree)
		Expect(err).ToNot(HaveOccurred())
		sort.Strings(idents)
		Expect(idents).To(Equal([]string{"age", "city"}))
	})

	It("Should handle complex expressions", func() {
		tree, err := tsl.ParseTSL("(name = 'john' or age > 20) and (city = 'London' or country = 'UK')")
		Expect(err).ToNot(HaveOccurred())

		idents, err := Walk(tree)
		Expect(err).ToNot(HaveOccurred())
		sort.Strings(idents)
		Expect(idents).To(Equal([]string{"age", "city", "country", "name"}))
	})

	It("Should not duplicate identifiers", func() {
		tree, err := tsl.ParseTSL("name = 'john' and name != 'doe'")
		Expect(err).ToNot(HaveOccurred())

		idents, err := Walk(tree)
		Expect(err).ToNot(HaveOccurred())
		Expect(idents).To(HaveLen(1))
		Expect(idents).To(ContainElement("name"))
	})

	It("Should handle arithmetic expressions", func() {
		tree, err := tsl.ParseTSL("salary * 2 > bonus + 1000")
		Expect(err).ToNot(HaveOccurred())

		idents, err := Walk(tree)
		Expect(err).ToNot(HaveOccurred())
		sort.Strings(idents)
		Expect(idents).To(Equal([]string{"bonus", "salary"}))
	})
})
