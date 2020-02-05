// Copyright 2018 Yaacov Zamir <kobi.zamir@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tsl

import "fmt"

// UnexpectedLiteralError is raised when an unexpected literal is found.
type UnexpectedLiteralError struct {
	ExpectedType string      // the expected literal type.
	Literal      interface{} // the literal found.
}

func (e UnexpectedLiteralError) Error() string {
	// if no literal is given then it's a missing literal error.
	if e.Literal == "" {
		return fmt.Sprintf("unexpected missing literal")
	}

	// If no type is given it's an unexpected literal error.
	if e.ExpectedType == "" {
		return fmt.Sprintf("unexpected literal: %v", e.Literal)
	}

	// o/w it's an unexpected type of literal error.
	return fmt.Sprintf("expected a %s literal, found: %v", e.ExpectedType, e.Literal)
}

// StackError is raised when the parser stack has unexpected size.
type StackError struct{}

func (e StackError) Error() string {
	return fmt.Sprintf("unexpected operator stack")
}

// ParseError is raised on parser error.
type ParseError struct {
	line   int
	column int
	msg    string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse error [%d:%d]: %s", e.line, e.column, e.msg)
}
