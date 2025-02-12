package tsl

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestTSLNodeMarshaling(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantJSON string
		wantYAML string
	}{
		{
			name:     "Simple identifier",
			input:    "name",
			wantJSON: `{"type":"IDENTIFIER","value":"name"}`,
			wantYAML: "type: IDENTIFIER\nvalue: name\n",
		},
		{
			name:     "Binary expression",
			input:    "age > 20",
			wantJSON: `{"type":"BINARY_EXP","operator":"GT","left":{"type":"IDENTIFIER","value":"age"},"right":{"type":"NUMBER","value":20}}`,
			wantYAML: "type: BINARY_EXP\noperator: GT\nleft:\n  type: IDENTIFIER\n  value: age\nright:\n  type: NUMBER\n  value: 20\n",
		},
		{
			name:     "Array literal",
			input:    "tags in [\"a\", \"b\"]",
			wantJSON: `{"type":"BINARY_EXP","operator":"IN","left":{"type":"IDENTIFIER","value":"tags"},"right":{"type":"ARRAY","values":[{"type":"STRING","value":"a"},{"type":"STRING","value":"b"}]}}`,
			wantYAML: "type: BINARY_EXP\noperator: IN\nleft:\n  type: IDENTIFIER\n  value: tags\nright:\n  type: ARRAY\n  values:\n  - type: STRING\n    value: a\n  - type: STRING\n    value: b\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input into a TSL tree
			node, err := ParseTSL(tt.input)
			assert.NoError(t, err)
			defer node.Free()

			// Test JSON marshaling
			jsonBytes, err := json.Marshal(node)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.wantJSON, string(jsonBytes))

			// Test YAML marshaling
			yamlBytes, err := yaml.Marshal(node)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantYAML, string(yamlBytes))
		})
	}
}
