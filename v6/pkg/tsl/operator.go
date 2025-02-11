package tsl

// Operator represents the type definitions.
// These values must match the yytokentype enum tsl_parser.tab.h
type Operator int

const (
	OpLike    Operator = 258 // K_LIKE
	OpILike   Operator = 259 // K_ILIKE
	OpAnd     Operator = 260 // K_AND
	OpOr      Operator = 261 // K_OR
	OpBetween Operator = 262 // K_BETWEEN
	OpIn      Operator = 263 // K_IN
	OpIs      Operator = 264 // K_IS
	OpNot     Operator = 266 // K_NOT
	OpLE      Operator = 269 // LE
	OpLT      Operator = 270 // LT
	OpGE      Operator = 271 // GE
	OpGT      Operator = 272 // GT
	OpNE      Operator = 273 // NE
	OpEQ      Operator = 274 // EQ
	OpREQ     Operator = 275 // REQ (Regular expression equals)
	OpRNE     Operator = 276 // RNE (Regular expression not equals)
	OpPlus    Operator = 282 // PLUS
	OpMinus   Operator = 283 // MINUS
	OpStar    Operator = 284 // STAR
	OpSlash   Operator = 285 // SLASH
	OpPercent Operator = 286 // PERCENT
	OpUMinus  Operator = 292 // UMINUS
)

// String returns the string representation of an OperatorType
func (op Operator) String() string {
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
