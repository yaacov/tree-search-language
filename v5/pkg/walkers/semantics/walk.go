// Copyright 2019 Yaacov Zamir <kobi.zamir@gmail.com>
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

// Author: 2019 Nimrod Shneor <nimrodshn@gmail.com>
// Author: 2019 Yaacov Zamir <kobi.zamir@gmail.com>

// Package semantics implements TSL tree semantics.
package semantics

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/yaacov/tree-search-language/v5/pkg/tsl"
)

// EvalFunc is a key evaluation function type.
//
// The function gets a key (string), and returns the value (interface{}) stored in the data document for that key.
type EvalFunc = func(string) (interface{}, bool)

// Walk travel the TSL tree and implements search semantics.
//
// Users can call the Walk method to check if a document compiles to `true` or `false`
// when applied to a tsl tree.
//
// Example:
//  	record :=  map[string]string {
//  		"title":       "A good book",
//  		"author":      "Joe",
//  		"spec.pages":  14,
//  		"spec.rating": 5,
//  	}
//
//  	// evalFactory creates an evaluation function for a data record.
//  	func evalFactory(r map[string]string) semantics.EvalFunc {
//  		// Returns:
//  		// A function (semantics.EvalFunc) that gets a `key` for a record and returns
//  		// the value of the document for that key.
//  		// If no value can be found for this `key` in our record, it will return
//  		// ok = false, if value is found it will return ok = true.
//  		return func(k string) (interface{}, bool) {
//  			v, ok := r[k]
//  			return v, ok
//  		}
//  	}
//
//  	// Check if our record complie with our tsl tree.
//  	//
//  	// For example:
//  	//   if our tsl tree represents the tsl phrase "author = 'Joe'"
//  	//   we will get the boolean value `true` for our record.
//  	//
//  	//   if our tsl tree represents the tsl phrase "spec.pages > 50"
//  	//   we will get the boolean value `false` for our record.
//  	eval :=  evalFactory(record)
//  	compliance, err = semantics.Walk(tree, eval)
//
func Walk(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)

	// Check for identifiers.
	if l.Func == tsl.IdentOp {
		newNode, err := handleIdentLeft(n, eval)
		if err != nil {
			return false, err
		}
		return Walk(newNode, eval)
	}

	// Check for math ops.
	if l.Func == tsl.AddOp || l.Func == tsl.SubtractOp || l.Func == tsl.MultiplyOp || l.Func == tsl.DivideOp {
		newNode, err := handleMathOpLeft(n, eval)
		if err != nil {
			return false, err
		}
		return Walk(newNode, eval)
	}

	// If we have a right hand side, check it for identifiers and math operators.
	if n.Right != nil {
		r := n.Right.(tsl.Node)

		// Check for identifiers.
		if r.Func == tsl.IdentOp {
			newNode, err := handleIdentRight(n, eval)
			if err != nil {
				return false, err
			}
			return Walk(newNode, eval)
		}

		// Check for math ops.
		if r.Func == tsl.AddOp || r.Func == tsl.SubtractOp || r.Func == tsl.MultiplyOp || r.Func == tsl.DivideOp {
			newNode, err := handleMathOpRight(n, eval)
			if err != nil {
				return false, err
			}
			return Walk(newNode, eval)
		}
	}

	// Implement tree semantics.
	return runSemantics(n, eval)
}

func runSemantics(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)

	switch n.Func {
	case tsl.EqOp, tsl.NotEqOp, tsl.LtOp, tsl.LteOp, tsl.GtOp, tsl.GteOp, tsl.RegexOp, tsl.NotRegexOp,
		tsl.BetweenOp, tsl.NotBetweenOp, tsl.NotInOp, tsl.InOp, tsl.LikeOp, tsl.ILikeOp:

		r := n.Right.(tsl.Node)
		if r.Func == tsl.NullOp {
			// Any comparison operation on a null element is false.
			return false, nil
		}

		switch l.Func {
		case tsl.ArrayOp:
			// This is a case of "has-many-table.fld IN (...)"
			if r.Func == tsl.ArrayOp {
				return handleStringArraysOp(n, eval)
			}
		case tsl.StringOp:
			if r.Func == tsl.StringOp {
				return handleStringOp(n, eval)
			}
			if r.Func == tsl.ArrayOp {
				return handleStringArrayOp(n, eval)
			}
		case tsl.NumberOp:
			if r.Func == tsl.NumberOp {
				return handleNumberOp(n, eval)
			}
			if r.Func == tsl.ArrayOp {
				return handleNumberArrayOp(n, eval)
			}
		case tsl.DateOp:
			if r.Func == tsl.DateOp {
				return handleDateOp(n, eval)
			}
			if r.Func == tsl.ArrayOp {
				return handleDateArrayOp(n, eval)
			}
		case tsl.BooleanOp:
			if r.Func == tsl.BooleanOp {
				return handleBooleanOp(n, eval)
			}
			if r.Func == tsl.StringOp {
				return handleBooleanOp(n, eval)
			}
			if r.Func == tsl.NumberOp {
				return handleBooleanOp(n, eval)
			}
		case tsl.NullOp:
			// Any comparison operation on a null element is false.
			return false, nil
		}

		return false, tsl.UnexpectedLiteralError{Literal: fmt.Sprintf("%v", r.Left)}
	case tsl.IsNotNilOp:
		return l.Func != tsl.NullOp, nil
	case tsl.IsNilOp:
		return l.Func == tsl.NullOp, nil
	case tsl.AndOp, tsl.OrOp:
		return handleLogicalOp(n, eval)
	case tsl.NotOp:
		return handleNotOp(n, eval)
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleIdentLeft(n tsl.Node, eval EvalFunc) (tsl.Node, error) {
	left := n.Left.(tsl.Node)
	l, err := evalIdentNode(left, eval)
	if err != nil {
		return n, err
	}

	n.Left = l
	return n, nil
}

func handleIdentRight(n tsl.Node, eval EvalFunc) (tsl.Node, error) {
	right := n.Right.(tsl.Node)
	r, err := evalIdentNode(right, eval)
	if err != nil {
		return n, err
	}

	n.Right = r
	return n, nil
}

func evalIdentNode(node tsl.Node, eval EvalFunc) (tsl.Node, error) {
	_v, _ := eval(node.Left.(string))
	n := tsl.Node{}

	switch v := _v.(type) {
	case []string:
		n = tsl.Node{
			Func: tsl.ArrayOp,
			Left: v,
		}
	case string:
		n = tsl.Node{
			Func: tsl.StringOp,
			Left: v,
		}
	case nil:
		n = tsl.Node{
			Func: tsl.NullOp,
			Left: nil,
		}
	case bool:
		n = tsl.Node{
			Func: tsl.BooleanOp,
			Left: v,
		}
	case time.Time:
		n = tsl.Node{
			Func: tsl.DateOp,
			Left: v,
		}
	case float32:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case float64:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: v,
		}
	case int32:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case int64:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case uint32:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case uint64:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case int:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case uint:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	default:
		return n, tsl.UnexpectedLiteralError{Literal: fmt.Sprintf("%v[%v]", _v, v)}
	}

	return n, nil
}

func handleMathOp(n tsl.Node, eval EvalFunc) (tsl.Node, error) {
	var err error

	left := n.Left.(tsl.Node)
	right := n.Right.(tsl.Node)

	// Check for identifiers.
	if left.Func == tsl.IdentOp {
		left, err = evalIdentNode(left, eval)
		if err != nil {
			return n, err
		}
	}
	if right.Func == tsl.IdentOp {
		right, err = evalIdentNode(right, eval)
		if err != nil {
			return n, err
		}
	}

	// Check for null
	if left.Func == tsl.NullOp || right.Func == tsl.NullOp {
		n = tsl.Node{
			Func: tsl.NullOp,
			Left: nil,
		}
		return n, nil
	}

	// Sanity check
	if left.Func != tsl.NumberOp || right.Func != tsl.NumberOp {
		return n, tsl.UnexpectedLiteralError{Literal: fmt.Sprintf("%v %v", left.Left, right.Left)}
	}

	switch n.Func {
	case tsl.AddOp:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: left.Left.(float64) + right.Left.(float64),
		}
	case tsl.SubtractOp:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: left.Left.(float64) - right.Left.(float64),
		}
	case tsl.MultiplyOp:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: left.Left.(float64) * right.Left.(float64),
		}
	case tsl.DivideOp:
		n = tsl.Node{
			Func: tsl.NumberOp,
			Left: left.Left.(float64) / right.Left.(float64),
		}
	default:
		return n, tsl.UnexpectedLiteralError{Literal: fmt.Sprintf("%v %v", left.Left, right.Left)}
	}

	return n, nil
}

func handleMathOpLeft(n tsl.Node, eval EvalFunc) (tsl.Node, error) {
	var err error

	left, err := handleMathOp(n.Left.(tsl.Node), eval)
	if err != nil {
		return n, err
	}

	n.Left = left
	return n, nil
}

func handleMathOpRight(n tsl.Node, eval EvalFunc) (tsl.Node, error) {
	var err error

	right, err := handleMathOp(n.Right.(tsl.Node), eval)
	if err != nil {
		return n, err
	}

	n.Right = right
	return n, nil
}

func handleStringOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(string)
	right := r.Left.(string)

	switch n.Func {
	case tsl.EqOp:
		return left == right, nil
	case tsl.NotEqOp:
		return left != right, nil
	case tsl.LtOp:
		return left < right, nil
	case tsl.LteOp:
		return left <= right, nil
	case tsl.GtOp:
		return left > right, nil
	case tsl.GteOp:
		return left >= right, nil
	case tsl.LikeOp:
		valid, err := regexp.Compile(likeToRegExp(right))
		if err != nil {
			return false, tsl.UnexpectedLiteralError{Literal: right}
		}
		return valid.MatchString(left), nil
	case tsl.ILikeOp:
		valid, err := regexp.Compile(iLikeToRegExp(right))
		if err != nil {
			return false, tsl.UnexpectedLiteralError{Literal: right}
		}
		return valid.MatchString(left), nil
	case tsl.RegexOp:
		valid, err := regexp.Compile(right)
		if err != nil {
			return false, tsl.UnexpectedLiteralError{Literal: right}
		}
		return valid.MatchString(left), nil
	case tsl.NotRegexOp:
		valid, err := regexp.Compile(right)
		if err != nil {
			return false, tsl.UnexpectedLiteralError{Literal: right}
		}
		return !valid.MatchString(left), nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleNumberOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(float64)
	right := r.Left.(float64)

	switch n.Func {
	case tsl.EqOp:
		return left == right, nil
	case tsl.NotEqOp:
		return left != right, nil
	case tsl.LtOp:
		return left < right, nil
	case tsl.LteOp:
		return left <= right, nil
	case tsl.GtOp:
		return left > right, nil
	case tsl.GteOp:
		return left >= right, nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleDateOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(time.Time)
	right := r.Left.(time.Time)

	switch n.Func {
	case tsl.EqOp:
		return left == right, nil
	case tsl.NotEqOp:
		return left != right, nil
	case tsl.LtOp:
		return right.After(left), nil
	case tsl.LteOp:
		return !right.Before(left), nil
	case tsl.GtOp:
		return left.After(right), nil
	case tsl.GteOp:
		return !left.Before(right), nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleBooleanOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(bool)
	right := false

	switch r.Func {
	case tsl.StringOp:
		right = (len(r.Left.(string)) == 1 || len(r.Left.(string)) == 4) &&
			(r.Left.(string)[0] == 't' || r.Left.(string)[0] == 'T')
	case tsl.NumberOp:
		right = r.Left.(float64) != 0
	case tsl.BooleanOp:
		right = r.Left.(bool)
	default:
		return false, tsl.UnexpectedLiteralError{Literal: n.Func}
	}

	switch n.Func {
	case tsl.EqOp:
		return left == right, nil
	case tsl.NotEqOp:
		return left != right, nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

// contains tells whether a contains x performing linear search
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func handleStringArraysOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.([]string)
	right := r.Right.([]tsl.Node)

	switch n.Func {
	case tsl.InOp:
		b := false
		for _, node := range right {
			s, ok := node.Left.(string)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b || contains(left, s)
		}
		return b, nil
	case tsl.NotInOp:
		b := true
		for _, node := range right {
			s, ok := node.Left.(string)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b && !contains(left, s)
		}
		return b, nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleStringArrayOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(string)
	right := r.Right.([]tsl.Node)

	switch n.Func {
	case tsl.BetweenOp:
		begin, ok := right[0].Left.(string)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		end, ok := right[1].Left.(string)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		return left >= begin && left < end, nil
	case tsl.NotBetweenOp:
		begin, ok := right[0].Left.(string)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		end, ok := right[1].Left.(string)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		return left < begin || left >= end, nil
	case tsl.InOp:
		b := false
		for _, node := range right {
			s, ok := node.Left.(string)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b || left == s
		}
		return b, nil
	case tsl.NotInOp:
		b := true
		for _, node := range right {
			s, ok := node.Left.(string)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b && left != s
		}
		return b, nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleNumberArrayOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(float64)
	right := r.Right.([]tsl.Node)

	switch n.Func {
	case tsl.BetweenOp:
		begin, ok := right[0].Left.(float64)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		end, ok := right[1].Left.(float64)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		return left >= begin && left < end, nil
	case tsl.NotBetweenOp:
		begin, ok := right[0].Left.(float64)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		end, ok := right[1].Left.(float64)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		return left < begin || left >= end, nil
	case tsl.InOp:
		b := false
		for _, node := range right {
			f, ok := node.Left.(float64)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b || left == f
		}
		return b, nil
	case tsl.NotInOp:
		b := true
		for _, node := range right {
			f, ok := node.Left.(float64)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b && left != f
		}
		return b, nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleDateArrayOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(time.Time)
	right := r.Right.([]tsl.Node)

	switch n.Func {
	case tsl.BetweenOp:
		begin, ok := right[0].Left.(time.Time)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		end, ok := right[1].Left.(time.Time)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		return !left.Before(begin) && left.Before(end), nil
	case tsl.NotBetweenOp:
		begin, ok := right[0].Left.(time.Time)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		end, ok := right[1].Left.(time.Time)
		if !ok {
			return false, tsl.UnexpectedLiteralError{Literal: n.Func}
		}
		return left.Before(begin) || !left.Before(end), nil
	case tsl.InOp:
		b := false
		for _, node := range right {
			t, ok := node.Left.(time.Time)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b || left == t
		}
		return b, nil
	case tsl.NotInOp:
		b := true
		for _, node := range right {
			t, ok := node.Left.(time.Time)
			if !ok {
				return false, tsl.UnexpectedLiteralError{Literal: n.Func}
			}
			b = b && left != t
		}
		return b, nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleLogicalOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	right, err := Walk(r, eval)
	if err != nil {
		return false, err
	}
	left, err := Walk(l, eval)
	if err != nil {
		return false, err
	}

	switch n.Func {
	case tsl.AndOp:
		return right && left, nil
	case tsl.OrOp:
		return right || left, nil
	}

	return false, tsl.UnexpectedLiteralError{Literal: n.Func}
}

func handleNotOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)

	left, err := Walk(l, eval)
	if err != nil {
		return false, err
	}

	return !left, nil
}

func likeToRegExp(l string) string {
	e := regexp.QuoteMeta(l)
	e = strings.Replace(e, "%", ".*", -1)
	e = strings.Replace(e, "_", ".", -1)
	e = fmt.Sprintf("^%s$", e)

	return e
}

func iLikeToRegExp(l string) string {
	return "(?i)" + likeToRegExp(l)
}
