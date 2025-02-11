package tsl

// NodeKind represents the kind of node in the AST
// These values must match the ast_node_type enum in tsl_ast.h
type NodeKind int

const (
	KindNumericLiteral   NodeKind = 0 // AST_NUMBER
	KindStringLiteral    NodeKind = 1 // AST_STRING
	KindIdentifier       NodeKind = 2 // AST_IDENTIFIER
	KindBinaryExpr       NodeKind = 3 // AST_BINARY_OP
	KindUnaryExpr        NodeKind = 4 // AST_UNARY_OP
	KindDateLiteral      NodeKind = 5 // AST_DATE
	KindTimestampLiteral NodeKind = 6 // AST_RFC3339
	KindArrayLiteral     NodeKind = 7 // AST_ARRAY
	KindBooleanLiteral   NodeKind = 8 // AST_BOOL
	KindNullLiteral      NodeKind = 9 // AST_NULL
)

// String returns the string representation of a NodeKind
func (kind NodeKind) String() string {
	switch kind {
	case KindStringLiteral:
		return "STRING"
	case KindNumericLiteral:
		return "NUMBER"
	case KindBooleanLiteral:
		return "BOOLEAN"
	case KindNullLiteral:
		return "NULL"
	case KindDateLiteral:
		return "DATE"
	case KindTimestampLiteral:
		return "TIMESTAMP"
	case KindArrayLiteral:
		return "ARRAY"
	case KindIdentifier:
		return "IDENTIFIER"
	case KindBinaryExpr:
		return "BINARY_EXP"
	case KindUnaryExpr:
		return "UNARY_EXP"
	default:
		return "UNKNOWN"
	}
}
