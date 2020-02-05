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

	"github.com/yaacov/tree-search-language/v4/pkg/tsl"
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
		newNode, err := handleIdent(n, eval)
		if err != nil {
			return false, err
		}
		return Walk(newNode, eval)
	}

	// Implement tree semantics.
	switch n.Func {
	case tsl.EqOp, tsl.NotEqOp, tsl.LtOp, tsl.LteOp, tsl.GtOp, tsl.GteOp, tsl.RegexOp, tsl.NotRegexOp,
		tsl.BetweenOp, tsl.NotBetweenOp, tsl.NotInOp, tsl.InOp, tsl.LikeOp, tsl.ILikeOp:
		r := n.Right.(tsl.Node)

		switch l.Func {
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

func handleIdent(n tsl.Node, eval EvalFunc) (tsl.Node, error) {
	l := n.Left.(tsl.Node)

	_v, _ := eval(l.Left.(string))
	switch v := _v.(type) {
	case string:
		n.Left = tsl.Node{
			Func: tsl.StringOp,
			Left: v,
		}
	case nil:
		n.Left = tsl.Node{
			Func: tsl.NullOp,
			Left: nil,
		}
	case bool:
		val := "false"
		if v {
			val = "true"
		}
		n.Left = tsl.Node{
			Func: tsl.StringOp,
			Left: val,
		}
	case float32:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case float64:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: v,
		}
	case int32:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case int64:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case uint32:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case uint64:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case int:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	case uint:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(v),
		}
	default:
		return n, tsl.UnexpectedLiteralError{Literal: fmt.Sprintf("%s[%v]", l.Left.(string), v)}
	}

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

func handleStringArrayOp(n tsl.Node, eval EvalFunc) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left := l.Left.(string)
	right := r.Right.([]tsl.Node)

	switch n.Func {
	case tsl.BetweenOp:
		begin := right[0].Left.(string)
		end := right[1].Left.(string)
		return left >= begin && left < end, nil
	case tsl.NotBetweenOp:
		begin := right[0].Left.(string)
		end := right[1].Left.(string)
		return left < begin || left >= end, nil
	case tsl.InOp:
		b := false
		for _, node := range right {
			b = b || left == node.Left.(string)
		}
		return b, nil
	case tsl.NotInOp:
		b := true
		for _, node := range right {
			b = b && left != node.Left.(string)
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
		begin := right[0].Left.(float64)
		end := right[1].Left.(float64)
		return left >= begin && left < end, nil
	case tsl.NotBetweenOp:
		begin := right[0].Left.(float64)
		end := right[1].Left.(float64)
		return left < begin || left >= end, nil
	case tsl.InOp:
		b := false
		for _, node := range right {
			b = b || left == node.Left.(float64)
		}
		return b, nil
	case tsl.NotInOp:
		b := true
		for _, node := range right {
			b = b && left != node.Left.(float64)
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
