package tsl

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v3"
)

func TestMarshal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TSL Marshal Suite")
}

var _ = Describe("TSL Node Marshaling", func() {
	DescribeTable("marshaling TSL nodes to JSON and YAML",
		func(input string, expectedJSON string, expectedYAML string) {
			node, err := ParseTSL(input)
			Expect(err).NotTo(HaveOccurred())
			defer node.Free()

			// Test JSON marshaling
			jsonBytes, err := json.Marshal(node)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(jsonBytes)).To(MatchJSON(expectedJSON))

			// Test YAML marshaling
			yamlBytes, err := yaml.Marshal(node)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(yamlBytes)).To(Equal(expectedYAML))
		},
		Entry("simple identifier",
			"name",
			`{"type":"IDENTIFIER","value":"name"}`,
			"type: IDENTIFIER\nvalue: name\n",
		),
		Entry("binary expression",
			"age > 20",
			`{"type":"BINARY_EXP","operator":"GT","left":{"type":"IDENTIFIER","value":"age"},"right":{"type":"NUMBER","value":20}}`,
			"type: BINARY_EXP\noperator: GT\nleft:\n    type: IDENTIFIER\n    value: age\nright:\n    type: NUMBER\n    value: 20\n",
		),
		Entry("array literal",
			"tags in [\"a\", \"b\"]",
			`{"type":"BINARY_EXP","operator":"IN","left":{"type":"IDENTIFIER","value":"tags"},"right":{"type":"ARRAY","values":[{"type":"STRING","value":"a"},{"type":"STRING","value":"b"}]}}`,
			"type: BINARY_EXP\noperator: IN\nleft:\n    type: IDENTIFIER\n    value: tags\nright:\n    type: ARRAY\n    values:\n        - type: STRING\n          value: a\n        - type: STRING\n          value: b\n",
		),
	)
})
