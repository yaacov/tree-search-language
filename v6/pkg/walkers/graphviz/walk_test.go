package graphviz

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

func TestWalk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Graphviz walker")
}

var _ = Describe("Walk", func() {
	DescribeTable("should generate correct graphviz output",
		func(input string, expectedContent []string) {
			tree, err := tsl.ParseTSL(input)
			Expect(err).ToNot(HaveOccurred())
			defer tree.Free()

			got, err := Walk("", tree, "")
			Expect(err).ToNot(HaveOccurred())

			for _, want := range expectedContent {
				Expect(got).To(ContainSubstring(want))
			}
		},
		Entry("simple equality",
			"name = 'john'",
			[]string{
				"[shape=box color=black label=\"EQ\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'name'\" ]",
				"[shape=record color=blue label=\"STRING | 'john'\" ]",
			}),
		Entry("numeric comparison",
			"age > 25",
			[]string{
				"[shape=box color=black label=\"GT\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'age'\" ]",
				"[shape=record color=blue label=\"NUMBER | 25\" ]",
			}),
		Entry("logical operation",
			"active = true and age > 18",
			[]string{
				"[shape=box color=black label=\"AND\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'active'\" ]",
				"[shape=record color=red label=\"IDENTIFIER | 'age'\" ]",
				"[shape=record color=blue label=\"NUMBER | 18\" ]",
			}),
		Entry("array operation",
			"status in ['active', 'pending']",
			[]string{
				"[shape=box color=black label=\"IN\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'status'\" ]",
				"[shape=box color=green label=\"ARRAY\"]",
				"[shape=record color=blue label=\"STRING | 'active'\" ]",
				"[shape=record color=blue label=\"STRING | 'pending'\" ]",
			}),
		Entry("date comparison",
			"created_at > 2023-01-01",
			[]string{
				"[shape=box color=black label=\"GT\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'created_at'\" ]",
				"[shape=record color=orange label=\"DATE | 2023-01-01\" ]",
			}),
		Entry("timestamp comparison",
			"updated_at < 2023-01-01T15:04:05Z",
			[]string{
				"[shape=box color=black label=\"LT\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'updated_at'\" ]",
				"[shape=record color=orange label=\"TIMESTAMP | 2023-01-01 15:04:05 +0000 UTC\" ]",
			}),
		Entry("complex nested expression",
			"(age > 18 and status in ['active', 'pending']) or (created_at > 2023-01-01 and not is_deleted)",
			[]string{
				"[shape=box color=black label=\"OR\"]",
				"[shape=box color=black label=\"AND\"]",
				"[shape=box color=black label=\"GT\"]",
				"[shape=box color=black label=\"IN\"]",
				"[shape=box color=black label=\"NOT\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'age'\" ]",
				"[shape=record color=blue label=\"NUMBER | 18\" ]",
				"[shape=box color=green label=\"ARRAY\"]",
				"[shape=record color=orange label=\"DATE | 2023-01-01\" ]",
				"[shape=record color=red label=\"IDENTIFIER | 'is_deleted'\" ]",
			}),
		Entry("mixed type comparisons",
			"count >= 100 and active = true and name like '%test%' and tags in ['a', 'b']",
			[]string{
				"[shape=box color=black label=\"AND\"]",
				"[shape=box color=black label=\"GE\"]",
				"[shape=box color=black label=\"EQ\"]",
				"[shape=box color=black label=\"LIKE\"]",
				"[shape=box color=black label=\"IN\"]",
				"[shape=record color=blue label=\"NUMBER | 100\" ]",
				"[shape=record color=purple label=\"BOOLEAN | true\" ]",
				"[shape=record color=blue label=\"STRING | '%test%'\" ]",
				"[shape=box color=green label=\"ARRAY\"]",
			}),
	)
})
