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

// TLS operators.
const (
	IdentOp      = "$ident"  // Empty operator for itentifiers
	ArrayOp      = "$array"  // Empty operator for arrays
	StringOp     = "$string" // Empty operator for strings
	NumberOp     = "$number" // Empty operator for numbers
	NullOp       = "$null"   // Empty operator for nulls
	LtOp         = "$lt"
	LteOp        = "$lte"
	GtOp         = "$gt"
	GteOp        = "$gte"
	EqOp         = "$eq"
	NotEqOp      = "$ne"
	RegexOp      = "$regex"
	NotRegexOp   = "$nregex"
	LikeOp       = "$like"
	ILikeOp      = "$ilike"
	InOp         = "$in"
	NotInOp      = "$nin"
	BetweenOp    = "$between"
	NotBetweenOp = "$nbetween"
	NotOp        = "$not"
	AndOp        = "$and"
	OrOp         = "$or"
	IsNilOp      = "$nexists"
	IsNotNilOp   = "$exists"
	AddOp        = "$add"
	SubtractOp   = "$subtract"
	MultiplyOp   = "$multiply"
	DivideOp     = "$divide"
	ModuloOp     = "$modulo"
)

// opDic maps SQL'ish operators to TLS operators.
var opDic = map[string]string{
	"<":  LtOp,
	"<=": LteOp,
	">":  GtOp,
	">=": GteOp,
	"=":  EqOp,
	"!=": NotEqOp,
	"<>": NotEqOp,
	"~=": RegexOp,
	"~!": NotRegexOp,
	"+":  AddOp,
	"-":  SubtractOp,
	"*":  MultiplyOp,
	"/":  DivideOp,
	"%":  ModuloOp,
}
