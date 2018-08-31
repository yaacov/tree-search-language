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

import (
	"fmt"
	"strconv"
	"strings"
)

// ternaryOp return lh if conditional is true, rh o/w.
func ternaryOp(conditional bool, lh string, rh string) string {
	if conditional {
		return lh
	}

	return rh
}

// literalValueToArg clean literal value, and return an arg string or float.
func literalValueToArg(v string) (arg interface{}) {
	l := len(v)

	// Check string length.
	if l < 1 {
		return
	}

	// Check for ''.
	if v[0] == '\'' {
		v = v[1 : l-1]
		arg = strings.Replace(v, "''", "'", -1)
		return
	}

	// It's a float.
	arg, _ = strconv.ParseFloat(v, 64)

	return
}

// literalValuesToArgs collect literal values, and create args list.
func literalValuesToArgs(c hasLiteralValues) (args []interface{}) {
	// Get length of literal values list.
	l := len(c.AllLiteralValue())

	// Create the arg list.
	args = make([]interface{}, l)
	for i := 0; i < l; i++ {
		args[i] = literalValueToArg(c.LiteralValue(i).GetText())
	}

	return
}

func (l *Listener) push(i Node) {
	l.Stack = append(l.Stack, i)
}

func (l *Listener) pop() (n Node) {
	// Check that we have nodes in the stack.
	size := len(l.Stack)
	if size < 1 {
		l.Err = fmt.Errorf("operator stack is empty")
		return
	}

	// Pop the last value from the Stack.
	n, l.Stack = l.Stack[size-1], l.Stack[:size-1]

	return
}
