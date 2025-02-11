package tsl

// OperatorType represents the type definitions.
// These values must match the yytokentype enum tsl_parser.tab.h
type OperatorType int

const (
	OpLike    OperatorType = 258 // K_LIKE
	OpILike   OperatorType = 259 // K_ILIKE
	OpAnd     OperatorType = 260 // K_AND
	OpOr      OperatorType = 261 // K_OR
	OpBetween OperatorType = 262 // K_BETWEEN
	OpIn      OperatorType = 263 // K_IN
	OpIs      OperatorType = 264 // K_IS
	OpNot     OperatorType = 266 // K_NOT
	OpLE      OperatorType = 269 // LE
	OpLT      OperatorType = 270 // LT
	OpGE      OperatorType = 271 // GE
	OpGT      OperatorType = 272 // GT
	OpNE      OperatorType = 273 // NE
	OpEQ      OperatorType = 274 // EQ
	OpREQ     OperatorType = 275 // REQ
	OpRNE     OperatorType = 276 // RNE
	OpPlus    OperatorType = 282 // PLUS
	OpMinus   OperatorType = 283 // MINUS
	OpStar    OperatorType = 284 // STAR
	OpSlash   OperatorType = 285 // SLASH
	OpPercent OperatorType = 286 // PERCENT
	OpUMinus  OperatorType = 292 // UMINUS
)

// String returns the string representation of an OperatorType
func (op OperatorType) String() string {
	switch op {
	// Comparison Operators
	case OpEQ:
		return "EQ"
	case OpNE:
		return "NE"
	case OpLT:
		return "LT"
	case OpLE:
		return "LE"
	case OpGT:
		return "GT"
	case OpGE:
		return "GE"
	case OpREQ:
		return "REQ"
	case OpRNE:
		return "RNE"

	// Logical Operators
	case OpAnd:
		return "AND"
	case OpOr:
		return "OR"
	case OpNot:
		return "NOT"

	// String Operators
	case OpLike:
		return "LIKE"
	case OpILike:
		return "ILIKE"

	// Array Operators
	case OpIn:
		return "IN"
	case OpBetween:
		return "BETWEEN"

	// Null Operators
	case OpIs:
		return "IS"

	// Arithmetic Operators
	case OpPlus:
		return "ADD"
	case OpMinus:
		return "SUB"
	case OpStar:
		return "MUL"
	case OpSlash:
		return "DIV"
	case OpPercent:
		return "MOD"
	case OpUMinus:
		return "NEG"

	default:
		return "UNKNOWN"
	}
}
