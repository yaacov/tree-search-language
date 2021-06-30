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
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/yaacov/tree-search-language/v5/pkg/tsl"
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
		"spec.pages":  14,
		"spec.rating": 5,
		"loaned":      true,
		"date":        date,
	}

	// This is the evaluation function that we will use to extract fields from the record:
	eval := func(name string) (value interface{}, ok bool) {
		value, ok = record[name]
		return
	}

	DescribeTable("Returns the expected result",
		func(text string, expected bool) {
			// Parse the text:
			tree, err := tsl.ParseTSL(text)
			Expect(err).ToNot(HaveOccurred())

			// Get the value:
			actual, err := Walk(tree, eval)
			Expect(err).ToNot(HaveOccurred())
			Expect(actual).To(Equal(expected))
		},

		// NOT combined with LIKE and ILIKE:
		Entry("like good", "title like '%good%'", true),
		Entry("like bad", "title like '%bad%'", false),
		Entry("ilike GOOD", "title ilike '%GOOD%'", true),
		Entry("ilike BAD", "title ilike '%BAD%'", false),
		Entry("not like good", "title not like '%good%'", false),
		Entry("not like bad", "title not like '%bad%'", true),
		Entry("not ilike GOOD", "title not ilike '%GOOD%'", false),
		Entry("not ilike BAD", "title not ilike '%BAD%'", true),

		// Two identifiers
		Entry("more pages", "spec.pages <= spec.rating", false),
		Entry("add pages to number", "spec.pages = (spec.rating + 9)", true),
		Entry("multiply pages with number", "spec.pages < (spec.rating * 3)", true),

		// Booleans
		Entry("booleans", "spec.pages < 20 and loaned = true", true),
		Entry("booleans false", "spec.pages < 20 and loaned != true", false),

		// dates
		Entry("dates", "date = 2020-01-01T00:00:00Z", true),
		Entry("dates", "date > 2020-01-02T00:00:00Z", false),
		Entry("dates", "date between 2019-12-30T00:00:00Z and 2020-01-02T00:00:00Z", true),

		// short dates
		Entry("dates", "date = 2020-01-01", true),
		Entry("dates", "date > 2020-01-02", false),
		Entry("dates", "date between 2019-12-30 and 2020-01-02", true),
	)
})

func evalFactory(record map[string]interface{}) EvalFunc {
	return func(name string) (value interface{}, ok bool) {
		value, ok = record[name]
		return
	}
}

var _ = Describe("Walk has-many relationship", func() {
	// Subscriptions has many ReservedResource
	// ReservedResource has string field ResourceName
	text := "subscription.managed = 'true' and subscription.status not in ('Deprovisioned','Deleted') and reserved_resource.resource_name in ('resourceA', 'resourceB')\n"

	// Parse the text:
	tree, err := tsl.ParseTSL(text)
	Expect(err).ToNot(HaveOccurred())

	subscriptionWithResourceAB := map[string]interface{}{
		"subscription.managed": true,
		"subscription.status":  "Active",
		"reserved_resource.resource_name": []string{
			"resourceA",
			"resourceB",
		},
	}
	subscriptionWithResourceB := map[string]interface{}{
		"subscription.managed": true,
		"subscription.status":  "Active",
		"reserved_resource.resource_name": []string{
			"resourceB",
		},
	}
	subscriptionWithoutResourceAB := map[string]interface{}{
		"subscription.managed": true,
		"subscription.status":  "Active",
		"reserved_resource.resource_name": []string{
			"resourceC",
		},
	}

	eval := evalFactory(subscriptionWithResourceAB)
	actual, err := Walk(tree, eval)
	Expect(err).ToNot(HaveOccurred())
	Expect(actual).To(Equal(true))

	eval = evalFactory(subscriptionWithResourceB)
	actual, err = Walk(tree, eval)
	Expect(err).ToNot(HaveOccurred())
	Expect(actual).To(Equal(true))

	eval = evalFactory(subscriptionWithoutResourceAB)
	actual, err = Walk(tree, eval)
	Expect(err).ToNot(HaveOccurred())
	Expect(actual).To(Equal(false))
})
