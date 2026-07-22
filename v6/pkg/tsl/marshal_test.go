package tsl

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo/v2"
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
		Entry("unary expression",
			"not active",
			`{"type":"UNARY_EXP","operator":"NOT","right":{"type":"IDENTIFIER","value":"active"}}`,
			"type: UNARY_EXP\noperator: NOT\nright:\n    type: IDENTIFIER\n    value: active\n",
		),
		Entry("boolean literal true",
			"active = true",
			`{"type":"BINARY_EXP","operator":"EQ","left":{"type":"IDENTIFIER","value":"active"},"right":{"type":"BOOLEAN","value":true}}`,
			"type: BINARY_EXP\noperator: EQ\nleft:\n    type: IDENTIFIER\n    value: active\nright:\n    type: BOOLEAN\n    value: true\n",
		),
		Entry("null literal",
			"name is null",
			`{"type":"BINARY_EXP","operator":"IS","left":{"type":"IDENTIFIER","value":"name"},"right":{"type":"NULL","value":"NULL"}}`,
			"type: BINARY_EXP\noperator: IS\nleft:\n    type: IDENTIFIER\n    value: name\nright:\n    type: \"NULL\"\n    value: \"NULL\"\n",
		),
		Entry("date literal",
			"created > 2023-01-01",
			`{"type":"BINARY_EXP","operator":"GT","left":{"type":"IDENTIFIER","value":"created"},"right":{"type":"DATE","value":"2023-01-01"}}`,
			"type: BINARY_EXP\noperator: GT\nleft:\n    type: IDENTIFIER\n    value: created\nright:\n    type: DATE\n    value: \"2023-01-01\"\n",
		),
		Entry("nested expression",
			"(a = 1 or b = 2) and c = 3",
			`{"type":"BINARY_EXP","operator":"AND","left":{"type":"BINARY_EXP","operator":"OR","left":{"type":"BINARY_EXP","operator":"EQ","left":{"type":"IDENTIFIER","value":"a"},"right":{"type":"NUMBER","value":1}},"right":{"type":"BINARY_EXP","operator":"EQ","left":{"type":"IDENTIFIER","value":"b"},"right":{"type":"NUMBER","value":2}}},"right":{"type":"BINARY_EXP","operator":"EQ","left":{"type":"IDENTIFIER","value":"c"},"right":{"type":"NUMBER","value":3}}}`,
			"type: BINARY_EXP\noperator: AND\nleft:\n    type: BINARY_EXP\n    operator: OR\n    left:\n        type: BINARY_EXP\n        operator: EQ\n        left:\n            type: IDENTIFIER\n            value: a\n        right:\n            type: NUMBER\n            value: 1\n    right:\n        type: BINARY_EXP\n        operator: EQ\n        left:\n            type: IDENTIFIER\n            value: b\n        right:\n            type: NUMBER\n            value: 2\nright:\n    type: BINARY_EXP\n    operator: EQ\n    left:\n        type: IDENTIFIER\n        value: c\n    right:\n        type: NUMBER\n        value: 3\n",
		),
	)
})
