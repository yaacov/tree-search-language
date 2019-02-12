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

package main

import (
	"fmt"
	"regexp"

	"github.com/yaacov/tsl/pkg/tsl"
)

func handleIdent(n tsl.Node, book Book) tsl.Node {
	l := n.Left.(tsl.Node)

	switch book[l.Left.(string)].(type) {
	case string:
		n.Left = tsl.Node{
			Func: tsl.StringOp,
			Left: book[l.Left.(string)].(string),
		}
	case float64:
		n.Left = tsl.Node{
			Func: tsl.StringOp,
			Left: book[l.Left.(string)].(float64),
		}
	case int:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(book[l.Left.(string)].(int)),
		}
	case uint:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: float64(book[l.Left.(string)].(uint)),
		}
	default:
		n.Left = tsl.Node{
			Func: tsl.NumberOp,
			Left: nil,
		}
	}

	return n
}

func handleNullOp(n tsl.Node, book Book) (bool, error) {
	var left interface{}

	l, ok := n.Left.(tsl.Node)
	if ok {
		left = l.Left
	} else {
		left = nil
	}

	switch n.Func {
	case tsl.EqOp:
		// if here left must be null and right must be not null
		return false, nil
	case tsl.NotEqOp:
		// if here left must be null and right must be not null
		return true, nil
	case tsl.IsNilOp:
		return left == nil, nil
	case tsl.IsNotNilOp:
		return left != nil, nil
	}

	return false, fmt.Errorf("not supported null operator")
}

func handleStringOp(n tsl.Node, book Book) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left, ok := l.Left.(string)
	if !ok {
		return handleNullOp(n, book)
	}
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
	case tsl.RegexOp:
		var valid = regexp.MustCompile(right)
		return valid.MatchString(left), nil
	case tsl.NotRegexOp:
		var valid = regexp.MustCompile(right)
		return !valid.MatchString(left), nil
	}

	return false, fmt.Errorf("not supported string operator")
}

func handleNumberOp(n tsl.Node, book Book) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left, ok := l.Left.(float64)
	if !ok {
		return handleNullOp(n, book)
	}
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

	return false, fmt.Errorf("not supported number operator")
}

func handleStringArrayOp(n tsl.Node, book Book) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left, ok := l.Left.(string)
	if !ok {
		return false, fmt.Errorf("not supported string array null operator")
	}
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

	return false, fmt.Errorf("not supported string array operator")
}

func handleNumberArrayOp(n tsl.Node, book Book) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	left, ok := l.Left.(float64)
	if !ok {
		return false, fmt.Errorf("not supported number array null operator")
	}
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

	return false, fmt.Errorf("not supported number array operator")
}

func handleLogicalOp(n tsl.Node, book Book) (bool, error) {
	l := n.Left.(tsl.Node)
	r := n.Right.(tsl.Node)

	right, err := Walk(r, book)
	if err != nil {
		return false, err
	}
	left, err := Walk(l, book)
	if err != nil {
		return false, err
	}

	switch n.Func {
	case tsl.AndOp:
		return right && left, nil
	case tsl.OrOp:
		return right || left, nil
	}

	return false, fmt.Errorf("not supported logical operator")
}

// Walk implements sql semantics.
func Walk(n tsl.Node, book Book) (bool, error) {
	// Walk tree.
	switch n.Func {
	case tsl.EqOp, tsl.NotEqOp, tsl.LtOp, tsl.LteOp, tsl.GtOp, tsl.GteOp, tsl.RegexOp, tsl.NotRegexOp,
		tsl.BetweenOp, tsl.NotBetweenOp, tsl.NotInOp, tsl.InOp:
		l := n.Left.(tsl.Node)
		r := n.Right.(tsl.Node)

		switch l.Func {
		case tsl.IdentOp:
			return Walk(handleIdent(n, book), book)
		case tsl.StringOp:
			if r.Func == tsl.StringOp {
				return handleStringOp(n, book)
			}
			if r.Func == tsl.ArrayOp {
				return handleStringArrayOp(n, book)
			}
		case tsl.NumberOp:
			if r.Func == tsl.NumberOp {
				return handleNumberOp(n, book)
			}
			if r.Func == tsl.ArrayOp {
				return handleNumberArrayOp(n, book)
			}
		}
	case tsl.IsNotNilOp, tsl.IsNilOp:
		l := n.Left.(tsl.Node)

		switch l.Func {
		case tsl.IdentOp:
			return Walk(handleIdent(n, book), book)
		default:
			return handleNullOp(n, book)
		}
	case tsl.AndOp, tsl.OrOp:
		return handleLogicalOp(n, book)
	}

	return false, fmt.Errorf("unimplemented syntax")
}
