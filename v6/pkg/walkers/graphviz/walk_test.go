package graphviz

import (
	"strings"
	"testing"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

func TestWalk(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		contains []string
	}{
		{
			name:    "simple equality",
			input:   "name = 'john'",
			wantErr: false,
			contains: []string{
				"[shape=box color=black label=\"EQ\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'name'\" ]",
				"[shape=record color=blue label=\"STRING | 'john'\" ]",
			},
		},
		{
			name:    "numeric comparison",
			input:   "age > 25",
			wantErr: false,
			contains: []string{
				"[shape=box color=black label=\"GT\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'age'\" ]",
				"[shape=record color=blue label=\"NUMBER | 25\" ]",
			},
		},
		{
			name:    "logical operation",
			input:   "active = true and age > 18",
			wantErr: false,
			contains: []string{
				"[shape=box color=black label=\"AND\"]",
				"[shape=record color=red label=\"IDENTIFIER | 'active'\" ]",
				"[shape=record color=red label=\"IDENTIFIER | 'age'\" ]",
				"[shape=record color=blue label=\"NUMBER | 18\" ]",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input into a TSL tree
			tree, err := tsl.ParseTSL(tt.input)
			if err != nil {
				t.Fatalf("failed to parse input: %v", err)
			}
			defer tree.Free()

			// Generate the graphviz output
			got, err := Walk("", tree, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("Walk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check if all expected strings are in the output
			for _, want := range tt.contains {
				if !strings.Contains(got, want) {
					t.Errorf("Walk() output missing expected content:\nwant: %s\ngot: %s", want, got)
				}
			}
		})
	}
}
