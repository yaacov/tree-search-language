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
	expectedType string // the expected literal type.
	literal      string // the literal found.
}

func (e UnexpectedLiteralError) Error() string {
	return fmt.Sprintf("expected a %s literal, found %s", e.expectedType, e.literal)
}

// StackError is raised when the parser stack has unexpected size.
type StackError struct{}

func (e StackError) Error() string {
	return fmt.Sprintf("unexpected operator stack")
}

// UnexpectedOpError is raised when an unexpected operand is found.
type UnexpectedOpError struct {
	Op string // the operand found.
}

func (e UnexpectedOpError) Error() string {
	return fmt.Sprintf("unexpected operand: %s", e.Op)
}

// UnexpectedArgError is raised when an unexpected argument is found.
type UnexpectedArgError struct {
	Arg interface{} // the argument found.
}

func (e UnexpectedArgError) Error() string {
	return fmt.Sprintf("unexpected argument: %v", e.Arg)
}

// MissingArgError is raised when argument is missing.
type MissingArgError struct{}

func (e MissingArgError) Error() string {
	return fmt.Sprintf("unexpected missing")
}
