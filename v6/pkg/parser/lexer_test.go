package parser

import (
	"testing"
)

func TestLexerTokenTypes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{"identifier", "name", []int{IDENTIFIER, EOF}},
		{"string single quotes", "'hello'", []int{STRING_LITERAL, EOF}},
		{"string double quotes", `"hello"`, []int{STRING_LITERAL, EOF}},
		{"number integer", "42", []int{NUMERIC_LITERAL, EOF}},
		{"number float", "3.14", []int{NUMERIC_LITERAL, EOF}},
		{"number scientific", "1e10", []int{NUMERIC_LITERAL, EOF}},
		{"number size suffix", "5Ki", []int{NUMERIC_LITERAL, EOF}},
		{"date", "2023-01-01", []int{DATE, EOF}},
		{"rfc3339", "2023-01-01T15:04:05Z", []int{RFC3339, EOF}},
		{"operators", "= != < <= > >=", []int{EQ, NE, LT, LE, GT, GE, EOF}},
		{"regex operators", "~= ~!", []int{REQ, RNE, EOF}},
		{"arithmetic", "+ - * / %", []int{PLUS, MINUS, STAR, SLASH, PERCENT, EOF}},
		{"parens and brackets", "( ) [ ]", []int{LPAREN, RPAREN, LBRACKET, RBRACKET, EOF}},
		{"comma", ",", []int{COMMA, EOF}},
		{"double equals", "==", []int{EQ, EOF}},
		{"not equals bang", "!=", []int{NE, EOF}},
		{"not equals diamond", "<>", []int{NE, EOF}},
		{"dotted identifier", "spec.pages", []int{IDENTIFIER, EOF}},
		{"identifier with slash", "path/to/field", []int{IDENTIFIER, EOF}},
		{"identifier with brackets", "items[0]", []int{IDENTIFIER, EOF}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.input)
			err := lexer.Tokenize()
			if err != nil {
				t.Fatalf("Tokenize error: %v", err)
			}

			if len(lexer.tokens) != len(tt.expected) {
				t.Fatalf("expected %d tokens, got %d", len(tt.expected), len(lexer.tokens))
			}

			for i, expectedType := range tt.expected {
				if lexer.tokens[i].Type != expectedType {
					t.Errorf("token %d: expected type %d, got %d (value: %q)",
						i, expectedType, lexer.tokens[i].Type, lexer.tokens[i].Value)
				}
			}
		})
	}
}

func TestLexerKeywords(t *testing.T) {
	keywords := []string{"and", "or", "not", "like", "ilike", "between", "in", "is", "null", "true", "false", "len", "any", "all", "sum"}

	for _, kw := range keywords {
		t.Run(kw, func(t *testing.T) {
			lexer := NewLexer(kw)
			err := lexer.Tokenize()
			if err != nil {
				t.Fatalf("Tokenize error: %v", err)
			}

			if len(lexer.tokens) != 2 {
				t.Fatalf("expected 2 tokens (keyword + EOF), got %d", len(lexer.tokens))
			}

			if lexer.tokens[0].Type == IDENTIFIER {
				t.Errorf("keyword %q was tokenized as IDENTIFIER instead of a keyword token", kw)
			}
		})
	}
}

func TestLexerUnicode(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"unicode identifier", "名前 = 'test'"},
		{"unicode string", "name = '日本語'"},
		{"emoji in string", "name = '🎉hello'"},
		{"accented identifier", "café = 'test'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.input)
			err := lexer.Tokenize()
			if err != nil {
				t.Fatalf("Tokenize error: %v", err)
			}
		})
	}
}

func TestLexerErrors(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"lone bang", "!"},
		{"lone tilde", "~"},
		{"unterminated string", "'hello"},
		{"unexpected character", "@"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.input)
			err := lexer.Tokenize()
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}

func TestLexerStringEscapes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"newline", `'hello\nworld'`, "hello\nworld"},
		{"tab", `'hello\tworld'`, "hello\tworld"},
		{"escaped quote", `'it\'s'`, "it's"},
		{"backslash", `'path\\to'`, "path\\to"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lexer := NewLexer(tt.input)
			err := lexer.Tokenize()
			if err != nil {
				t.Fatalf("Tokenize error: %v", err)
			}

			if lexer.tokens[0].Value != tt.expected {
				t.Errorf("expected value %q, got %q", tt.expected, lexer.tokens[0].Value)
			}
		})
	}
}

func TestParseSizeValue(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"100", 100},
		{"3.14", 3.14},
		{"5k", 5000},
		{"5K", 5000},
		{"2M", 2000000},
		{"1G", 1000000000},
		{"1T", 1000000000000},
		{"1P", 1000000000000000},
		{"5Ki", 5120},
		{"2Mi", 2097152},
		{"1Gi", 1073741824},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := parseSizeValue(tt.input)
			if err != nil {
				t.Fatalf("parseSizeValue error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{"empty input", "", true},
		{"whitespace only", "   ", true},
		{"simple expression", "a = 1", false},
		{"deeply nested parens", "((((a = 1))))", false},
		{"trailing comma in array", "[1, 2, 3,]", false},
		{"empty array", "a in []", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.input)
			if tt.expectErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestParseErrorPositions(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectedPos int
	}{
		{"unexpected token", "a = = b", 4},
		{"unterminated string", "'hello", 0},
		{"unicode prefix", "名前 = = b", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.input)
			if err == nil {
				t.Fatal("expected error, got nil")
			}

			parseErr, ok := err.(*ParseError)
			if !ok {
				t.Fatalf("expected *ParseError, got %T", err)
			}
			if parseErr.Position != tt.expectedPos {
				t.Errorf("expected position %d, got %d", tt.expectedPos, parseErr.Position)
			}
		})
	}
}
