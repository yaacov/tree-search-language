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
	"github.com/yaacov/tsl/pkg/parser"
)

// Node is a Tree search node
type Node struct {
	Func  string      `json:"func"`
	Left  interface{} `json:"left,omitempty"`
	Right interface{} `json:"right,omitempty"`
}

// Listener is a Tree search listener
type Listener struct {
	*parser.BaseTSLListener

	Stack []Node
	Err   error
}

// hasLiteralValues is a type that contain literal values
type hasLiteralValues interface {
	AllLiteralValue() []parser.ILiteralValueContext
	LiteralValue(int) parser.ILiteralValueContext
}
