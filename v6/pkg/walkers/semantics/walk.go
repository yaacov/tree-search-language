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

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

// EvalFunc is a key evaluation function type
type EvalFunc = func(string) (interface{}, bool)

// Walk traverses the TSL tree and implements search semantics.
//
// Users can call the Walk method to check if a document compiles to `true` or `false`
// when applied to a tsl tree.
//
// Example:
//
//	record := map[string]interface{} {
//		"title":       "A good book",
//		"author":      "Joe",
//		"spec.pages":  14,
//		"spec.rating": 5,
//		"created_at":  time.Now(),
//		"is_active":   true,
//	}
//
//	// evalFactory creates an evaluation function for a data record.
//	func evalFactory(r map[string]interface{}) semantics.EvalFunc {
//		// Returns:
//		// A function (semantics.EvalFunc) that gets a `key` for a record and returns
//		// the value of the document for that key.
//		// If no value can be found for this `key` in our record, it will return
//		// ok = false, if value is found it will return ok = true.
//		return func(k string) (interface{}, bool) {
//			v, ok := r[k]
//			return v, ok
//		}
//	}
//
//	// Check if our record complies with our tsl tree.
//	//
//	// For example:
//	//   if our tsl tree represents the phrase "author = 'Joe' and created_at > '2023-01-01'"
//	//   we will get the boolean value `true` for our record.
//	//
//	//   if our tsl tree represents the phrase "spec.pages > 50"
//	//   we will get the boolean value `false` for our record.
//	eval := evalFactory(record)
//	compliance, err = semantics.Walk(tree, eval)
func Walk(n *tsl.TSLNode, eval EvalFunc) (interface{}, error) {
	if n == nil {
		return nil, nil
	}

	// If this is an identifier, evaluate it
	if n.Type() == tsl.KindIdentifier {
		return evalIdentNode(n, eval)
	}

	// For binary expressions, handle each side
	if n.Type() == tsl.KindBinaryExpr {
		exprOp := n.Value().(tsl.TSLExpressionOp)
		return evalBinaryExpr(exprOp, eval)
	}

	// For unary expressions
	if n.Type() == tsl.KindUnaryExpr {
		exprOp := n.Value().(tsl.TSLExpressionOp)
		return evalUnaryExpr(exprOp, eval)
	}

	// For literal values, return them directly
	return n.Value(), nil
}

func evalBinaryExpr(expr tsl.TSLExpressionOp, eval EvalFunc) (interface{}, error) {
	leftVal, err := Walk(expr.Left, eval)
	if err != nil {
		return nil, err
	}
	rightVal, err := Walk(expr.Right, eval)
	if err != nil {
		return nil, err
	}

	// Handle arithmetic operations
	if leftNum, okLeft := leftVal.(float64); okLeft {
		if rightNum, okRight := rightVal.(float64); okRight {
			switch expr.Operator {
			case tsl.OpPlus:
				return evaluateArithmetic(leftNum, rightNum, expr.Operator)
			case tsl.OpMinus:
				return evaluateArithmetic(leftNum, rightNum, expr.Operator)
			case tsl.OpStar:
				return evaluateArithmetic(leftNum, rightNum, expr.Operator)
			case tsl.OpSlash:
				return evaluateArithmetic(leftNum, rightNum, expr.Operator)
			case tsl.OpPercent:
				return evaluateArithmetic(leftNum, rightNum, expr.Operator)
			}
		}
	}

	// Handle other operators
	switch expr.Operator {
	case tsl.OpAnd, tsl.OpOr:
		return evalLogicalOp(leftVal, rightVal, expr.Operator)
	case tsl.OpEQ, tsl.OpNE, tsl.OpLT, tsl.OpLE, tsl.OpGT, tsl.OpGE:
		return evalComparisonOp(leftVal, rightVal, expr.Operator)
	case tsl.OpLike, tsl.OpILike, tsl.OpREQ, tsl.OpRNE:
		return evalStringMatchOp(leftVal, rightVal, expr.Operator)
	case tsl.OpIn, tsl.OpBetween:
		return evalArrayOp(leftVal, rightVal, expr.Operator, eval)
	case tsl.OpIs:
		return evalIsOp(leftVal, rightVal, expr.Operator)
	}

	return nil, tsl.UnexpectedOperatorError{Operator: expr.Operator}
}

func evalUnaryExpr(expr tsl.TSLExpressionOp, eval EvalFunc) (interface{}, error) {
	if expr.Operator != tsl.OpNot {
		return nil, tsl.UnexpectedOperatorError{Operator: expr.Operator}
	}

	rightVal, err := Walk(expr.Right, eval)
	if err != nil {
		return nil, err
	}

	rightBool, ok := rightVal.(bool)
	if !ok {
		return nil, tsl.TypeMismatchError{Expected: "boolean", Got: fmt.Sprintf("%T", rightVal)}
	}

	return !rightBool, nil
}

func evalComparisonOp(left, right interface{}, op tsl.Operator) (interface{}, error) {
	switch l := left.(type) {
	case string:
		r, ok := right.(string)
		if !ok {
			return nil, tsl.TypeMismatchError{Expected: "string", Got: fmt.Sprintf("%T", right)}
		}
		result, err := compareStrings(l, r, op)
		return result, err

	case float64:
		r, ok := right.(float64)
		if !ok {
			return nil, tsl.TypeMismatchError{Expected: "number", Got: fmt.Sprintf("%T", right)}
		}
		result, err := compareNumbers(l, r, op)
		return result, err

	case time.Time:
		r, ok := right.(time.Time)
		if !ok {
			return nil, tsl.TypeMismatchError{Expected: "time", Got: fmt.Sprintf("%T", right)}
		}
		result, err := compareDates(l, r, op)
		return result, err

	case bool:
		r, ok := right.(bool)
		if !ok {
			return nil, tsl.TypeMismatchError{Expected: "boolean", Got: fmt.Sprintf("%T", right)}
		}
		result, err := compareBooleans(l, r, op)
		return result, err
	}

	return nil, tsl.TypeMismatchError{Expected: "comparable type", Got: fmt.Sprintf("%T", left)}
}

func evalLogicalOp(left, right interface{}, op tsl.Operator) (interface{}, error) {
	leftBool, ok := left.(bool)
	if !ok {
		return nil, tsl.TypeMismatchError{Expected: "boolean", Got: fmt.Sprintf("%T", left)}
	}
	rightBool, ok := right.(bool)
	if !ok {
		return nil, tsl.TypeMismatchError{Expected: "boolean", Got: fmt.Sprintf("%T", right)}
	}

	switch op {
	case tsl.OpAnd:
		return leftBool && rightBool, nil
	case tsl.OpOr:
		return leftBool || rightBool, nil
	}

	return nil, tsl.UnexpectedOperatorError{Operator: op}
}

func evalArrayOp(leftVal, rightVal interface{}, op tsl.Operator, eval EvalFunc) (interface{}, error) {
	if op == tsl.OpBetween {
		array, ok := rightVal.(tsl.TSLArrayLiteral)
		if !ok || len(array.Values) != 2 {
			return nil, tsl.BetweenOperatorError{Message: "requires exactly 2 values"}
		}

		begin, err := Walk(array.Values[0], eval)
		if err != nil {
			return nil, err
		}
		end, err := Walk(array.Values[1], eval)
		if err != nil {
			return nil, err
		}

		switch leftVal.(type) {
		case string:
			return leftVal.(string) >= begin.(string) && leftVal.(string) < end.(string), nil
		case float64:
			return leftVal.(float64) >= begin.(float64) && leftVal.(float64) < end.(float64), nil
		case time.Time:
			leftTime := leftVal.(time.Time)
			return !leftTime.Before(begin.(time.Time)) && leftTime.Before(end.(time.Time)), nil
		}
		return nil, tsl.TypeMismatchError{Expected: "comparable type", Got: fmt.Sprintf("%T", leftVal)}
	}

	// Handle IN operator
	array, ok := rightVal.(tsl.TSLArrayLiteral)
	if !ok {
		return nil, tsl.TypeMismatchError{Expected: "array", Got: fmt.Sprintf("%T", rightVal)}
	}

	for _, elem := range array.Values {
		elemVal, err := Walk(elem, eval)
		if err != nil {
			return nil, err
		}
		if elemVal == leftVal {
			return true, nil
		}
	}
	return false, nil
}

func evalIdentNode(n *tsl.TSLNode, eval EvalFunc) (interface{}, error) {
	if n == nil {
		return nil, nil
	}

	// If not an identifier, just return the node
	if n.Type() != tsl.KindIdentifier {
		return Walk(n, eval)
	}

	// Get value using eval function
	value, exists := eval(n.Value().(string))
	if !exists {
		return nil, nil
	}

	return value, nil
}

func evalIsOp(leftVal, rightVal interface{}, op tsl.Operator) (interface{}, error) {
	if op != tsl.OpIs {
		return nil, tsl.UnexpectedOperatorError{Operator: op}
	}

	// Handle IS NULL check
	if rightVal == nil || rightVal == "NULL" {
		return leftVal == nil, nil
	}

	// Both values should be non-nil for other comparisons
	return leftVal == rightVal, nil
}

// Comparison helper functions
func compareStrings(left, right string, op tsl.Operator) (bool, error) {
	switch op {
	case tsl.OpEQ:
		return left == right, nil
	case tsl.OpNE:
		return left != right, nil
	case tsl.OpLT:
		return left < right, nil
	case tsl.OpLE:
		return left <= right, nil
	case tsl.OpGT:
		return left > right, nil
	case tsl.OpGE:
		return left >= right, nil
	}
	return false, tsl.UnexpectedOperatorError{Operator: op}
}

func compareNumbers(left, right float64, op tsl.Operator) (bool, error) {
	switch op {
	case tsl.OpEQ:
		return left == right, nil
	case tsl.OpNE:
		return left != right, nil
	case tsl.OpLT:
		return left < right, nil
	case tsl.OpLE:
		return left <= right, nil
	case tsl.OpGT:
		return left > right, nil
	case tsl.OpGE:
		return left >= right, nil
	}
	return false, tsl.UnexpectedOperatorError{Operator: op}
}

func evaluateArithmetic(left, right float64, op tsl.Operator) (float64, error) {
	switch op {
	case tsl.OpPlus:
		return left + right, nil
	case tsl.OpMinus:
		return left - right, nil
	case tsl.OpStar:
		return left * right, nil
	case tsl.OpSlash:
		if right == 0 {
			return 0, tsl.DivisionByZeroError{Operation: "division"}
		}
		return left / right, nil
	case tsl.OpPercent:
		if right == 0 {
			return 0, tsl.DivisionByZeroError{Operation: "modulo"}
		}
		return float64(int64(left) % int64(right)), nil
	}
	return 0, tsl.UnexpectedOperatorError{Operator: op}
}

func compareDates(left, right time.Time, op tsl.Operator) (bool, error) {
	switch op {
	case tsl.OpEQ:
		return left.Equal(right), nil
	case tsl.OpNE:
		return !left.Equal(right), nil
	case tsl.OpLT:
		return left.Before(right), nil
	case tsl.OpLE:
		return left.Before(right) || left.Equal(right), nil
	case tsl.OpGT:
		return right.Before(left), nil
	case tsl.OpGE:
		return right.Before(left) || left.Equal(right), nil
	}
	return false, tsl.UnexpectedOperatorError{Operator: op}
}

func compareBooleans(left, right bool, op tsl.Operator) (bool, error) {
	switch op {
	case tsl.OpEQ:
		return left == right, nil
	case tsl.OpNE:
		return left != right, nil
	}
	return false, tsl.UnexpectedOperatorError{Operator: op}
}

func evalStringMatchOp(leftVal, rightVal interface{}, op tsl.Operator) (interface{}, error) {
	leftStr, ok1 := leftVal.(string)
	rightStr, ok2 := rightVal.(string)
	if !ok1 || !ok2 {
		return nil, tsl.TypeMismatchError{Expected: "string", Got: fmt.Sprintf("%T,%T", leftVal, rightVal)}
	}

	// Case-insensitive comparison for ILIKE
	if op == tsl.OpILike {
		leftStr = strings.ToLower(leftStr)
		rightStr = strings.ToLower(rightStr)
	}

	switch op {
	case tsl.OpLike, tsl.OpILike:
		// Convert SQL LIKE pattern to regex
		pattern := strings.ReplaceAll(rightStr, "%", ".*")
		pattern = strings.ReplaceAll(pattern, "_", ".")
		pattern = "^" + pattern + "$"

		matched, err := regexp.MatchString(pattern, leftStr)
		if err != nil {
			return nil, err
		}
		return matched, nil

	case tsl.OpREQ:
		matched, err := regexp.MatchString(rightStr, leftStr)
		if err != nil {
			return nil, err
		}
		return matched, nil

	case tsl.OpRNE:
		matched, err := regexp.MatchString(rightStr, leftStr)
		if err != nil {
			return nil, err
		}
		return !matched, nil
	}

	return nil, tsl.UnexpectedOperatorError{Operator: op}
}

// Walk performs semantic validation on the TSL tree and returns a bool indicating validity
func validateTree(n *tsl.TSLNode) error {
	switch n.Type() {
	case tsl.KindBinaryExpr:
		return validateBinaryExpr(n)
	case tsl.KindUnaryExpr:
		return validateUnaryExpr(n)
	case tsl.KindIdentifier, tsl.KindNumericLiteral, tsl.KindStringLiteral,
		tsl.KindBooleanLiteral, tsl.KindDateLiteral, tsl.KindTimestampLiteral,
		tsl.KindArrayLiteral, tsl.KindNullLiteral:
		return nil
	default:
		return tsl.UnexpectedTypeError{Type: n.Type()}
	}
}

func validateBinaryExpr(n *tsl.TSLNode) error {
	op := n.Value().(tsl.TSLExpressionOp)

	// For arithmetic operations, validate recursively
	if op.Operator == tsl.OpPlus || op.Operator == tsl.OpMinus ||
		op.Operator == tsl.OpStar || op.Operator == tsl.OpSlash ||
		op.Operator == tsl.OpPercent {
		if err := validateTree(op.Left); err != nil {
			return err
		}
		if err := validateTree(op.Right); err != nil {
			return err
		}
		return validateArithmeticOperands(op.Left, op.Right)
	}

	// Validate left and right nodes recursively
	if err := validateTree(op.Left); err != nil {
		return err
	}
	if err := validateTree(op.Right); err != nil {
		return err
	}

	// Validate operator
	switch op.Operator {
	case tsl.OpAnd, tsl.OpOr:
		return validateLogicalOperands(op.Left, op.Right)
	case tsl.OpPlus, tsl.OpMinus, tsl.OpStar, tsl.OpSlash, tsl.OpPercent:
		return validateArithmeticOperands(op.Left, op.Right)
	case tsl.OpEQ, tsl.OpNE, tsl.OpLT, tsl.OpLE, tsl.OpGT, tsl.OpGE:
		return validateComparisonOperands(op.Left, op.Right)
	case tsl.OpLike, tsl.OpILike, tsl.OpREQ, tsl.OpRNE:
		return validateStringOperands(op.Left, op.Right)
	case tsl.OpIn, tsl.OpBetween:
		return validateArrayOperands(op.Left, op.Right)
	case tsl.OpIs:
		return validateNullOperands(op.Left, op.Right)
	default:
		return tsl.UnexpectedOperatorError{Operator: op.Operator}
	}
}

func validateUnaryExpr(n *tsl.TSLNode) error {
	op := n.Value().(tsl.TSLExpressionOp)

	// Validate the operand recursively
	if err := validateTree(op.Right); err != nil {
		return err
	}

	if op.Operator != tsl.OpNot {
		return tsl.UnexpectedOperatorError{Operator: op.Operator}
	}

	return nil
}

func validateLogicalOperands(_, _ *tsl.TSLNode) error {
	// For now, allow any expression in logical operations
	return nil
}

func validateArithmeticOperands(_, _ *tsl.TSLNode) error {
	// Allow any numeric expression
	return nil
}

func validateComparisonOperands(left, _ *tsl.TSLNode) error {
	// Allow comparison between identifiers and literals
	if left.Type() != tsl.KindIdentifier {
		return tsl.TypeMismatchError{Expected: "identifier", Got: left.Type()}
	}
	return nil
}

func validateStringOperands(left, right *tsl.TSLNode) error {
	if left.Type() != tsl.KindIdentifier {
		return tsl.TypeMismatchError{Expected: "identifier", Got: left.Type()}
	}
	if right.Type() != tsl.KindStringLiteral {
		return tsl.TypeMismatchError{Expected: "string", Got: right.Type()}
	}
	return nil
}

func validateArrayOperands(left, right *tsl.TSLNode) error {
	if left.Type() != tsl.KindIdentifier {
		return tsl.TypeMismatchError{Expected: "identifier", Got: left.Type()}
	}
	if right.Type() != tsl.KindArrayLiteral {
		return tsl.TypeMismatchError{Expected: "array", Got: right.Type()}
	}
	return nil
}

func validateNullOperands(left, right *tsl.TSLNode) error {
	if left.Type() != tsl.KindIdentifier {
		return tsl.TypeMismatchError{Expected: "identifier", Got: left.Type()}
	}
	if right.Type() != tsl.KindNullLiteral {
		return tsl.TypeMismatchError{Expected: "NULL", Got: right.Type()}
	}
	return nil
}
