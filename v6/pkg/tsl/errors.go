package tsl

import "fmt"

// SyntaxError represents a TSL parsing error with context
type SyntaxError struct {
	Message  string
	Position int
	Context  string
	Input    string
}

// Error implements the error interface
func (e *SyntaxError) Error() string {
	// Create error pointer string
	pointer := ""
	for i := 0; i < e.Position; i++ {
		pointer += " "
	}
	pointer += "^"

	return fmt.Sprintf("%s at position %d:\n%s\n%s",
		e.Message, e.Position, e.Input, pointer)
}

// UnexpectedLiteralError is returned when encountering an unexpected literal or operator
type UnexpectedLiteralError struct {
	Literal interface{}
}

func (e UnexpectedLiteralError) Error() string {
	return fmt.Sprintf("unexpected literal: %v", e.Literal)
}
