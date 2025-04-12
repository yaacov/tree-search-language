package tsl

// Operator represents the type definitions.
// These values must match the yytokentype enum tsl_parser.tab.h
type Operator int

const (
	OpLike    Operator = 258 // K_LIKE (SQL LIKE)
	OpILike   Operator = 259 // K_ILIKE (SQL ILIKE)
	OpAnd     Operator = 260 // K_AND
	OpOr      Operator = 261 // K_OR
	OpBetween Operator = 262 // K_BETWEEN
	OpIn      Operator = 263 // K_IN
	OpIs      Operator = 264 // K_IS (Only for NULL and NOT NULL)
	OpNot     Operator = 266 // K_NOT
	OpLen     Operator = 269 // K_LEN
	OpAny     Operator = 270 // K_ANY
	OpAll     Operator = 271 // K_ALL
	OpLE      Operator = 290 // LE
	OpLT      Operator = 289 // LT
	OpGE      Operator = 292 // GE
	OpGT      Operator = 291 // GT
	OpNE      Operator = 288 // NE
	OpEQ      Operator = 287 // EQ
	OpREQ     Operator = 293 // REQ (Regular expression equals)
	OpRNE     Operator = 294 // RNE (Regular expression not equals)
	OpPlus    Operator = 277 // PLUS
	OpMinus   Operator = 278 // MINUS
	OpStar    Operator = 279 // STAR (MULTIPLY)
	OpSlash   Operator = 280 // SLASH (DIVIDE)
	OpPercent Operator = 281 // PERCENT (MODULUS)
	OpUMinus  Operator = 295 // UMINUS
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

	// Unary Operators
	case OpUMinus:
		return "NEG"
	case OpLen:
		return "LEN"
	case OpAny:
		return "ANY"
	case OpAll:
		return "ALL"

	default:
		return "UNKNOWN"
	}
}
