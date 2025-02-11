package semantics

import "fmt"

// UnexpectedTypeError is returned when encountering an unexpected type in the AST
type UnexpectedTypeError struct {
	Type interface{}
}

func (e UnexpectedTypeError) Error() string {
	return fmt.Sprintf("unexpected type: %v", e.Type)
}

// UnexpectedOperatorError is returned when encountering an unexpected operator
type UnexpectedOperatorError struct {
	Operator interface{}
}

func (e UnexpectedOperatorError) Error() string {
	return fmt.Sprintf("unexpected operator: %v", e.Operator)
}

// TypeMismatchError is returned when operand types don't match the operator
type TypeMismatchError struct {
	Expected string
	Got      interface{}
}

func (e TypeMismatchError) Error() string {
	return fmt.Sprintf("type mismatch: expected %s, got %v", e.Expected, e.Got)
}
