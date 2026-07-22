package parser

import "testing"

func BenchmarkParseSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := Parse("name = 'alice'"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParseComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := Parse("(age > 25 and city = 'rome') or (status in ['active', 'pending'] and price between 10 and 100)"); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLexerSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewLexer("name = 'alice'")
		if err := l.Tokenize(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLexerComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := NewLexer("(age > 25 and city = 'rome') or (status in ['active', 'pending'] and price between 10 and 100)")
		if err := l.Tokenize(); err != nil {
			b.Fatal(err)
		}
	}
}
