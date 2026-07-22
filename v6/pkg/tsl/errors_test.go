package tsl

import (
	"strings"
	"testing"
)

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantMsg string
		wantPos int
	}{
		{"missing operand", "name = ", "syntax error", 6},
		{"unterminated string", "'unterminated", "Unterminated string", 0},
		{"unexpected character", "name @ value", "Unexpected character", 5},
		{"leading operator", "= value", "syntax error", 0},
		{"empty input", "", "syntax error", 0},
		{"trailing keyword", "a > 1 and", "syntax error", 6},
		{"lone tilde", "a ~ b", "Unexpected character", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseTSL(tt.input)
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			se, ok := err.(*SyntaxError)
			if !ok {
				t.Fatalf("expected *SyntaxError, got %T: %v", err, err)
			}

			if !strings.Contains(se.Message, tt.wantMsg) {
				t.Errorf("message %q does not contain %q", se.Message, tt.wantMsg)
			}

			if se.Position != tt.wantPos {
				t.Errorf("position = %d, want %d", se.Position, tt.wantPos)
			}

			full := se.Error()
			if tt.input != "" && !strings.Contains(full, tt.input) {
				t.Errorf("formatted error does not contain original input %q:\n%s", tt.input, full)
			}
			if !strings.Contains(full, "^") {
				t.Errorf("formatted error missing caret pointer:\n%s", full)
			}
		})
	}
}
