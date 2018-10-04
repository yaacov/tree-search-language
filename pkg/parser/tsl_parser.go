// Code generated from TSL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TSL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 185,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 45, 10, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 53, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 60, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 66, 10, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 75, 10, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 82, 10, 3, 12, 3, 14, 3, 85, 11, 3, 5, 3, 87, 10, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 97, 10, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 105, 10, 3, 12, 3, 14, 3, 108, 11, 3,
	3, 4, 3, 4, 5, 4, 112, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 125, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 130,
	10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 140,
	10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 146, 10, 10, 12, 10, 14, 10,
	149, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 157, 10,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 167,
	10, 11, 12, 11, 14, 11, 170, 11, 11, 3, 12, 5, 12, 173, 10, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 181, 10, 14, 3, 15, 3, 15, 3,
	15, 2, 5, 4, 18, 20, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 2, 7, 3, 2, 6, 9, 3, 2, 10, 12, 3, 2, 13, 17, 3, 2, 18, 19, 3, 2, 16,
	17, 2, 197, 2, 30, 3, 2, 2, 2, 4, 96, 3, 2, 2, 2, 6, 111, 3, 2, 2, 2, 8,
	113, 3, 2, 2, 2, 10, 115, 3, 2, 2, 2, 12, 117, 3, 2, 2, 2, 14, 119, 3,
	2, 2, 2, 16, 129, 3, 2, 2, 2, 18, 139, 3, 2, 2, 2, 20, 156, 3, 2, 2, 2,
	22, 172, 3, 2, 2, 2, 24, 176, 3, 2, 2, 2, 26, 180, 3, 2, 2, 2, 28, 182,
	3, 2, 2, 2, 30, 31, 5, 4, 3, 2, 31, 32, 7, 2, 2, 3, 32, 3, 3, 2, 2, 2,
	33, 34, 8, 3, 1, 2, 34, 35, 5, 20, 11, 2, 35, 36, 5, 6, 4, 2, 36, 37, 5,
	26, 14, 2, 37, 97, 3, 2, 2, 2, 38, 39, 5, 20, 11, 2, 39, 40, 5, 10, 6,
	2, 40, 41, 5, 24, 13, 2, 41, 97, 3, 2, 2, 2, 42, 44, 5, 20, 11, 2, 43,
	45, 5, 28, 15, 2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 46, 3, 2,
	2, 2, 46, 47, 7, 21, 2, 2, 47, 48, 5, 24, 13, 2, 48, 97, 3, 2, 2, 2, 49,
	50, 5, 20, 11, 2, 50, 52, 7, 26, 2, 2, 51, 53, 5, 28, 15, 2, 52, 51, 3,
	2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 7, 27, 2, 2, 55,
	97, 3, 2, 2, 2, 56, 57, 5, 20, 11, 2, 57, 59, 7, 26, 2, 2, 58, 60, 5, 28,
	15, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61,
	62, 5, 26, 14, 2, 62, 97, 3, 2, 2, 2, 63, 65, 5, 20, 11, 2, 64, 66, 5,
	28, 15, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2,
	67, 68, 7, 24, 2, 2, 68, 69, 5, 26, 14, 2, 69, 70, 7, 22, 2, 2, 70, 71,
	5, 26, 14, 2, 71, 97, 3, 2, 2, 2, 72, 74, 5, 20, 11, 2, 73, 75, 5, 28,
	15, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76,
	77, 7, 25, 2, 2, 77, 86, 7, 3, 2, 2, 78, 83, 5, 26, 14, 2, 79, 80, 7, 4,
	2, 2, 80, 82, 5, 26, 14, 2, 81, 79, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83,
	81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2,
	2, 86, 78, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89,
	7, 5, 2, 2, 89, 97, 3, 2, 2, 2, 90, 91, 7, 28, 2, 2, 91, 97, 5, 4, 3, 6,
	92, 93, 7, 3, 2, 2, 93, 94, 5, 4, 3, 2, 94, 95, 7, 5, 2, 2, 95, 97, 3,
	2, 2, 2, 96, 33, 3, 2, 2, 2, 96, 38, 3, 2, 2, 2, 96, 42, 3, 2, 2, 2, 96,
	49, 3, 2, 2, 2, 96, 56, 3, 2, 2, 2, 96, 63, 3, 2, 2, 2, 96, 72, 3, 2, 2,
	2, 96, 90, 3, 2, 2, 2, 96, 92, 3, 2, 2, 2, 97, 106, 3, 2, 2, 2, 98, 99,
	12, 5, 2, 2, 99, 100, 7, 22, 2, 2, 100, 105, 5, 4, 3, 6, 101, 102, 12,
	4, 2, 2, 102, 103, 7, 23, 2, 2, 103, 105, 5, 4, 3, 5, 104, 98, 3, 2, 2,
	2, 104, 101, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106,
	107, 3, 2, 2, 2, 107, 5, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 112, 9,
	2, 2, 2, 110, 112, 9, 3, 2, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2,
	2, 112, 7, 3, 2, 2, 2, 113, 114, 9, 4, 2, 2, 114, 9, 3, 2, 2, 2, 115, 116,
	9, 5, 2, 2, 116, 11, 3, 2, 2, 2, 117, 118, 7, 29, 2, 2, 118, 13, 3, 2,
	2, 2, 119, 120, 7, 29, 2, 2, 120, 15, 3, 2, 2, 2, 121, 122, 5, 12, 7, 2,
	122, 123, 7, 20, 2, 2, 123, 125, 3, 2, 2, 2, 124, 121, 3, 2, 2, 2, 124,
	125, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 5, 14, 8, 2, 127, 128,
	7, 20, 2, 2, 128, 130, 3, 2, 2, 2, 129, 124, 3, 2, 2, 2, 129, 130, 3, 2,
	2, 2, 130, 131, 3, 2, 2, 2, 131, 132, 7, 29, 2, 2, 132, 17, 3, 2, 2, 2,
	133, 134, 8, 10, 1, 2, 134, 140, 5, 22, 12, 2, 135, 136, 7, 3, 2, 2, 136,
	137, 5, 18, 10, 2, 137, 138, 7, 5, 2, 2, 138, 140, 3, 2, 2, 2, 139, 133,
	3, 2, 2, 2, 139, 135, 3, 2, 2, 2, 140, 147, 3, 2, 2, 2, 141, 142, 12, 4,
	2, 2, 142, 143, 5, 8, 5, 2, 143, 144, 5, 18, 10, 5, 144, 146, 3, 2, 2,
	2, 145, 141, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147,
	148, 3, 2, 2, 2, 148, 19, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150, 151, 8,
	11, 1, 2, 151, 157, 5, 16, 9, 2, 152, 153, 7, 3, 2, 2, 153, 154, 5, 20,
	11, 2, 154, 155, 7, 5, 2, 2, 155, 157, 3, 2, 2, 2, 156, 150, 3, 2, 2, 2,
	156, 152, 3, 2, 2, 2, 157, 168, 3, 2, 2, 2, 158, 159, 12, 5, 2, 2, 159,
	160, 5, 8, 5, 2, 160, 161, 5, 20, 11, 6, 161, 167, 3, 2, 2, 2, 162, 163,
	12, 4, 2, 2, 163, 164, 5, 8, 5, 2, 164, 165, 5, 18, 10, 2, 165, 167, 3,
	2, 2, 2, 166, 158, 3, 2, 2, 2, 166, 162, 3, 2, 2, 2, 167, 170, 3, 2, 2,
	2, 168, 166, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 21, 3, 2, 2, 2, 170,
	168, 3, 2, 2, 2, 171, 173, 9, 6, 2, 2, 172, 171, 3, 2, 2, 2, 172, 173,
	3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 7, 30, 2, 2, 175, 23, 3, 2,
	2, 2, 176, 177, 7, 31, 2, 2, 177, 25, 3, 2, 2, 2, 178, 181, 5, 22, 12,
	2, 179, 181, 5, 24, 13, 2, 180, 178, 3, 2, 2, 2, 180, 179, 3, 2, 2, 2,
	181, 27, 3, 2, 2, 2, 182, 183, 7, 28, 2, 2, 183, 29, 3, 2, 2, 2, 22, 44,
	52, 59, 65, 74, 83, 86, 96, 104, 106, 111, 124, 129, 139, 147, 156, 166,
	168, 172, 180,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'<>'",
	"'*'", "'/'", "'%'", "'+'", "'-'", "'~='", "'~!'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "K_LIKE", "K_AND", "K_OR", "K_BETWEEN", "K_IN", "K_IS", "K_NULL", "K_NOT",
	"IDENTIFIER", "NUMERIC_LITERAL", "STRING_LITERAL", "SPACES",
}

var ruleNames = []string{
	"start", "expr", "literalOp", "numericOp", "stringOp", "databaseName",
	"tableName", "columnName", "numericExp", "columnOp", "signedNumber", "stringValue",
	"literalValue", "keyNot",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TSLParser struct {
	*antlr.BaseParser
}

func NewTSLParser(input antlr.TokenStream) *TSLParser {
	this := new(TSLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TSL.g4"

	return this
}

// TSLParser tokens.
const (
	TSLParserEOF             = antlr.TokenEOF
	TSLParserT__0            = 1
	TSLParserT__1            = 2
	TSLParserT__2            = 3
	TSLParserT__3            = 4
	TSLParserT__4            = 5
	TSLParserT__5            = 6
	TSLParserT__6            = 7
	TSLParserT__7            = 8
	TSLParserT__8            = 9
	TSLParserT__9            = 10
	TSLParserT__10           = 11
	TSLParserT__11           = 12
	TSLParserT__12           = 13
	TSLParserT__13           = 14
	TSLParserT__14           = 15
	TSLParserT__15           = 16
	TSLParserT__16           = 17
	TSLParserT__17           = 18
	TSLParserK_LIKE          = 19
	TSLParserK_AND           = 20
	TSLParserK_OR            = 21
	TSLParserK_BETWEEN       = 22
	TSLParserK_IN            = 23
	TSLParserK_IS            = 24
	TSLParserK_NULL          = 25
	TSLParserK_NOT           = 26
	TSLParserIDENTIFIER      = 27
	TSLParserNUMERIC_LITERAL = 28
	TSLParserSTRING_LITERAL  = 29
	TSLParserSPACES          = 30
)

// TSLParser rules.
const (
	TSLParserRULE_start        = 0
	TSLParserRULE_expr         = 1
	TSLParserRULE_literalOp    = 2
	TSLParserRULE_numericOp    = 3
	TSLParserRULE_stringOp     = 4
	TSLParserRULE_databaseName = 5
	TSLParserRULE_tableName    = 6
	TSLParserRULE_columnName   = 7
	TSLParserRULE_numericExp   = 8
	TSLParserRULE_columnOp     = 9
	TSLParserRULE_signedNumber = 10
	TSLParserRULE_stringValue  = 11
	TSLParserRULE_literalValue = 12
	TSLParserRULE_keyNot       = 13
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(TSLParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *TSLParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TSLParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.expr(0)
	}
	{
		p.SetState(29)
		p.Match(TSLParserEOF)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParContext struct {
	*ExprContext
}

func NewParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParContext {
	var p = new(ParContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterPar(s)
	}
}

func (s *ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitPar(s)
	}
}

type NotContext struct {
	*ExprContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(TSLParserK_NOT, 0)
}

func (s *NotContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitNot(s)
	}
}

type LikeContext struct {
	*ExprContext
}

func NewLikeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikeContext {
	var p = new(LikeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *LikeContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(TSLParserK_LIKE, 0)
}

func (s *LikeContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *LikeContext) KeyNot() IKeyNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyNotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyNotContext)
}

func (s *LikeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterLike(s)
	}
}

func (s *LikeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitLike(s)
	}
}

type OrContext struct {
	*ExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) K_OR() antlr.TerminalNode {
	return s.GetToken(TSLParserK_OR, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitOr(s)
	}
}

type InContext struct {
	*ExprContext
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *InContext) K_IN() antlr.TerminalNode {
	return s.GetToken(TSLParserK_IN, 0)
}

func (s *InContext) KeyNot() IKeyNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyNotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyNotContext)
}

func (s *InContext) AllLiteralValue() []ILiteralValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem())
	var tst = make([]ILiteralValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralValueContext)
		}
	}

	return tst
}

func (s *InContext) LiteralValue(i int) ILiteralValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitIn(s)
	}
}

type IsLiteralContext struct {
	*ExprContext
}

func NewIsLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsLiteralContext {
	var p = new(IsLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IsLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsLiteralContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *IsLiteralContext) K_IS() antlr.TerminalNode {
	return s.GetToken(TSLParserK_IS, 0)
}

func (s *IsLiteralContext) LiteralValue() ILiteralValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *IsLiteralContext) KeyNot() IKeyNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyNotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyNotContext)
}

func (s *IsLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterIsLiteral(s)
	}
}

func (s *IsLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitIsLiteral(s)
	}
}

type AndContext struct {
	*ExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) K_AND() antlr.TerminalNode {
	return s.GetToken(TSLParserK_AND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitAnd(s)
	}
}

type BetweenContext struct {
	*ExprContext
}

func NewBetweenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenContext {
	var p = new(BetweenContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BetweenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *BetweenContext) K_BETWEEN() antlr.TerminalNode {
	return s.GetToken(TSLParserK_BETWEEN, 0)
}

func (s *BetweenContext) AllLiteralValue() []ILiteralValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem())
	var tst = make([]ILiteralValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralValueContext)
		}
	}

	return tst
}

func (s *BetweenContext) LiteralValue(i int) ILiteralValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *BetweenContext) K_AND() antlr.TerminalNode {
	return s.GetToken(TSLParserK_AND, 0)
}

func (s *BetweenContext) KeyNot() IKeyNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyNotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyNotContext)
}

func (s *BetweenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterBetween(s)
	}
}

func (s *BetweenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitBetween(s)
	}
}

type StringOpsContext struct {
	*ExprContext
}

func NewStringOpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringOpsContext {
	var p = new(StringOpsContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringOpsContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *StringOpsContext) StringOp() IStringOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringOpContext)
}

func (s *StringOpsContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *StringOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterStringOps(s)
	}
}

func (s *StringOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitStringOps(s)
	}
}

type IsNullContext struct {
	*ExprContext
}

func NewIsNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullContext {
	var p = new(IsNullContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IsNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *IsNullContext) K_IS() antlr.TerminalNode {
	return s.GetToken(TSLParserK_IS, 0)
}

func (s *IsNullContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(TSLParserK_NULL, 0)
}

func (s *IsNullContext) KeyNot() IKeyNotContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyNotContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeyNotContext)
}

func (s *IsNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterIsNull(s)
	}
}

func (s *IsNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitIsNull(s)
	}
}

type LiteralOpsContext struct {
	*ExprContext
}

func NewLiteralOpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralOpsContext {
	var p = new(LiteralOpsContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOpsContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *LiteralOpsContext) LiteralOp() ILiteralOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOpContext)
}

func (s *LiteralOpsContext) LiteralValue() ILiteralValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *LiteralOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterLiteralOps(s)
	}
}

func (s *LiteralOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitLiteralOps(s)
	}
}

func (p *TSLParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *TSLParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, TSLParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralOpsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(32)
			p.columnOp(0)
		}
		{
			p.SetState(33)
			p.LiteralOp()
		}
		{
			p.SetState(34)
			p.LiteralValue()
		}

	case 2:
		localctx = NewStringOpsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.columnOp(0)
		}
		{
			p.SetState(37)
			p.StringOp()
		}
		{
			p.SetState(38)
			p.StringValue()
		}

	case 3:
		localctx = NewLikeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.columnOp(0)
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(41)
				p.KeyNot()
			}

		}
		{
			p.SetState(44)
			p.Match(TSLParserK_LIKE)
		}
		{
			p.SetState(45)
			p.StringValue()
		}

	case 4:
		localctx = NewIsNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.columnOp(0)
		}
		{
			p.SetState(48)
			p.Match(TSLParserK_IS)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(49)
				p.KeyNot()
			}

		}
		{
			p.SetState(52)
			p.Match(TSLParserK_NULL)
		}

	case 5:
		localctx = NewIsLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.columnOp(0)
		}
		{
			p.SetState(55)
			p.Match(TSLParserK_IS)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(56)
				p.KeyNot()
			}

		}
		{
			p.SetState(59)
			p.LiteralValue()
		}

	case 6:
		localctx = NewBetweenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.columnOp(0)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(62)
				p.KeyNot()
			}

		}
		{
			p.SetState(65)
			p.Match(TSLParserK_BETWEEN)
		}
		{
			p.SetState(66)
			p.LiteralValue()
		}
		{
			p.SetState(67)
			p.Match(TSLParserK_AND)
		}
		{
			p.SetState(68)
			p.LiteralValue()
		}

	case 7:
		localctx = NewInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.columnOp(0)
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(71)
				p.KeyNot()
			}

		}
		{
			p.SetState(74)
			p.Match(TSLParserK_IN)
		}

		{
			p.SetState(75)
			p.Match(TSLParserT__0)
		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TSLParserT__13)|(1<<TSLParserT__14)|(1<<TSLParserNUMERIC_LITERAL)|(1<<TSLParserSTRING_LITERAL))) != 0 {
			{
				p.SetState(76)
				p.LiteralValue()
			}
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TSLParserT__1 {
				{
					p.SetState(77)
					p.Match(TSLParserT__1)
				}
				{
					p.SetState(78)
					p.LiteralValue()
				}

				p.SetState(83)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(86)
			p.Match(TSLParserT__2)
		}

	case 8:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(TSLParserK_NOT)
		}
		{
			p.SetState(89)
			p.expr(4)
		}

	case 9:
		localctx = NewParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Match(TSLParserT__0)
		}
		{
			p.SetState(91)
			p.expr(0)
		}
		{
			p.SetState(92)
			p.Match(TSLParserT__2)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(97)
					p.Match(TSLParserK_AND)
				}
				{
					p.SetState(98)
					p.expr(4)
				}

			case 2:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(100)
					p.Match(TSLParserK_OR)
				}
				{
					p.SetState(101)
					p.expr(3)
				}

			}

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// ILiteralOpContext is an interface to support dynamic dispatch.
type ILiteralOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralOpContext differentiates from other interfaces.
	IsLiteralOpContext()
}

type LiteralOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralOpContext() *LiteralOpContext {
	var p = new(LiteralOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_literalOp
	return p
}

func (*LiteralOpContext) IsLiteralOpContext() {}

func NewLiteralOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralOpContext {
	var p = new(LiteralOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_literalOp

	return p
}

func (s *LiteralOpContext) GetParser() antlr.Parser { return s.parser }
func (s *LiteralOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterLiteralOp(s)
	}
}

func (s *LiteralOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitLiteralOp(s)
	}
}

func (p *TSLParser) LiteralOp() (localctx ILiteralOpContext) {
	localctx = NewLiteralOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TSLParserRULE_literalOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserT__3, TSLParserT__4, TSLParserT__5, TSLParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TSLParserT__3)|(1<<TSLParserT__4)|(1<<TSLParserT__5)|(1<<TSLParserT__6))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case TSLParserT__7, TSLParserT__8, TSLParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TSLParserT__7)|(1<<TSLParserT__8)|(1<<TSLParserT__9))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumericOpContext is an interface to support dynamic dispatch.
type INumericOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericOpContext differentiates from other interfaces.
	IsNumericOpContext()
}

type NumericOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericOpContext() *NumericOpContext {
	var p = new(NumericOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_numericOp
	return p
}

func (*NumericOpContext) IsNumericOpContext() {}

func NewNumericOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericOpContext {
	var p = new(NumericOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_numericOp

	return p
}

func (s *NumericOpContext) GetParser() antlr.Parser { return s.parser }
func (s *NumericOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterNumericOp(s)
	}
}

func (s *NumericOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitNumericOp(s)
	}
}

func (p *TSLParser) NumericOp() (localctx INumericOpContext) {
	localctx = NewNumericOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TSLParserRULE_numericOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TSLParserT__10)|(1<<TSLParserT__11)|(1<<TSLParserT__12)|(1<<TSLParserT__13)|(1<<TSLParserT__14))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringOpContext is an interface to support dynamic dispatch.
type IStringOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringOpContext differentiates from other interfaces.
	IsStringOpContext()
}

type StringOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringOpContext() *StringOpContext {
	var p = new(StringOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_stringOp
	return p
}

func (*StringOpContext) IsStringOpContext() {}

func NewStringOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringOpContext {
	var p = new(StringOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_stringOp

	return p
}

func (s *StringOpContext) GetParser() antlr.Parser { return s.parser }
func (s *StringOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterStringOp(s)
	}
}

func (s *StringOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitStringOp(s)
	}
}

func (p *TSLParser) StringOp() (localctx IStringOpContext) {
	localctx = NewStringOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TSLParserRULE_stringOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TSLParserT__15 || _la == TSLParserT__16) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDatabaseNameContext is an interface to support dynamic dispatch.
type IDatabaseNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatabaseNameContext differentiates from other interfaces.
	IsDatabaseNameContext()
}

type DatabaseNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabaseNameContext() *DatabaseNameContext {
	var p = new(DatabaseNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_databaseName
	return p
}

func (*DatabaseNameContext) IsDatabaseNameContext() {}

func NewDatabaseNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabaseNameContext {
	var p = new(DatabaseNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_databaseName

	return p
}

func (s *DatabaseNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabaseNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TSLParserIDENTIFIER, 0)
}

func (s *DatabaseNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabaseNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabaseNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterDatabaseName(s)
	}
}

func (s *DatabaseNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitDatabaseName(s)
	}
}

func (p *TSLParser) DatabaseName() (localctx IDatabaseNameContext) {
	localctx = NewDatabaseNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TSLParserRULE_databaseName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(TSLParserIDENTIFIER)
	}

	return localctx
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_tableName
	return p
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TSLParserIDENTIFIER, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *TSLParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TSLParserRULE_tableName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(TSLParserIDENTIFIER)
	}

	return localctx
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TSLParserIDENTIFIER, 0)
}

func (s *ColumnNameContext) TableName() ITableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *ColumnNameContext) DatabaseName() IDatabaseNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatabaseNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatabaseNameContext)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (p *TSLParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TSLParserRULE_columnName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(119)
				p.DatabaseName()
			}
			{
				p.SetState(120)
				p.Match(TSLParserT__17)
			}

		}
		{
			p.SetState(124)
			p.TableName()
		}
		{
			p.SetState(125)
			p.Match(TSLParserT__17)
		}

	}
	{
		p.SetState(129)
		p.Match(TSLParserIDENTIFIER)
	}

	return localctx
}

// INumericExpContext is an interface to support dynamic dispatch.
type INumericExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericExpContext differentiates from other interfaces.
	IsNumericExpContext()
}

type NumericExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericExpContext() *NumericExpContext {
	var p = new(NumericExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_numericExp
	return p
}

func (*NumericExpContext) IsNumericExpContext() {}

func NewNumericExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericExpContext {
	var p = new(NumericExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_numericExp

	return p
}

func (s *NumericExpContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericExpContext) SignedNumber() ISignedNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedNumberContext)
}

func (s *NumericExpContext) AllNumericExp() []INumericExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericExpContext)(nil)).Elem())
	var tst = make([]INumericExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericExpContext)
		}
	}

	return tst
}

func (s *NumericExpContext) NumericExp(i int) INumericExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericExpContext)
}

func (s *NumericExpContext) NumericOp() INumericOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericOpContext)
}

func (s *NumericExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterNumericExp(s)
	}
}

func (s *NumericExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitNumericExp(s)
	}
}

func (p *TSLParser) NumericExp() (localctx INumericExpContext) {
	return p.numericExp(0)
}

func (p *TSLParser) numericExp(_p int) (localctx INumericExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewNumericExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumericExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, TSLParserRULE_numericExp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserT__13, TSLParserT__14, TSLParserNUMERIC_LITERAL:
		{
			p.SetState(132)
			p.SignedNumber()
		}

	case TSLParserT__0:
		{
			p.SetState(133)
			p.Match(TSLParserT__0)
		}
		{
			p.SetState(134)
			p.numericExp(0)
		}
		{
			p.SetState(135)
			p.Match(TSLParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewNumericExpContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_numericExp)
			p.SetState(139)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(140)
				p.NumericOp()
			}
			{
				p.SetState(141)
				p.numericExp(3)
			}

		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IColumnOpContext is an interface to support dynamic dispatch.
type IColumnOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnOpContext differentiates from other interfaces.
	IsColumnOpContext()
}

type ColumnOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnOpContext() *ColumnOpContext {
	var p = new(ColumnOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_columnOp
	return p
}

func (*ColumnOpContext) IsColumnOpContext() {}

func NewColumnOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnOpContext {
	var p = new(ColumnOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_columnOp

	return p
}

func (s *ColumnOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnOpContext) CopyFrom(ctx *ColumnOpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ColumnOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ColumnContext struct {
	*ColumnOpContext
}

func NewColumnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnContext {
	var p = new(ColumnContext)

	p.ColumnOpContext = NewEmptyColumnOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ColumnOpContext))

	return p
}

func (s *ColumnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterColumn(s)
	}
}

func (s *ColumnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitColumn(s)
	}
}

type ColumnNameOpContext struct {
	*ColumnOpContext
}

func NewColumnNameOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnNameOpContext {
	var p = new(ColumnNameOpContext)

	p.ColumnOpContext = NewEmptyColumnOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ColumnOpContext))

	return p
}

func (s *ColumnNameOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameOpContext) AllColumnOp() []IColumnOpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IColumnOpContext)(nil)).Elem())
	var tst = make([]IColumnOpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IColumnOpContext)
		}
	}

	return tst
}

func (s *ColumnNameOpContext) ColumnOp(i int) IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *ColumnNameOpContext) NumericOp() INumericOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericOpContext)
}

func (s *ColumnNameOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterColumnNameOp(s)
	}
}

func (s *ColumnNameOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitColumnNameOp(s)
	}
}

type ColumnNameNumericOpContext struct {
	*ColumnOpContext
}

func NewColumnNameNumericOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnNameNumericOpContext {
	var p = new(ColumnNameNumericOpContext)

	p.ColumnOpContext = NewEmptyColumnOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ColumnOpContext))

	return p
}

func (s *ColumnNameNumericOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameNumericOpContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *ColumnNameNumericOpContext) NumericOp() INumericOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericOpContext)
}

func (s *ColumnNameNumericOpContext) NumericExp() INumericExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericExpContext)
}

func (s *ColumnNameNumericOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterColumnNameNumericOp(s)
	}
}

func (s *ColumnNameNumericOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitColumnNameNumericOp(s)
	}
}

type ColumnNameParContext struct {
	*ColumnOpContext
}

func NewColumnNameParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnNameParContext {
	var p = new(ColumnNameParContext)

	p.ColumnOpContext = NewEmptyColumnOpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ColumnOpContext))

	return p
}

func (s *ColumnNameParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameParContext) ColumnOp() IColumnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnOpContext)
}

func (s *ColumnNameParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterColumnNamePar(s)
	}
}

func (s *ColumnNameParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitColumnNamePar(s)
	}
}

func (p *TSLParser) ColumnOp() (localctx IColumnOpContext) {
	return p.columnOp(0)
}

func (p *TSLParser) columnOp(_p int) (localctx IColumnOpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewColumnOpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IColumnOpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, TSLParserRULE_columnOp, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserIDENTIFIER:
		localctx = NewColumnContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(149)
			p.ColumnName()
		}

	case TSLParserT__0:
		localctx = NewColumnNameParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(150)
			p.Match(TSLParserT__0)
		}
		{
			p.SetState(151)
			p.columnOp(0)
		}
		{
			p.SetState(152)
			p.Match(TSLParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewColumnNameOpContext(p, NewColumnOpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_columnOp)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(157)
					p.NumericOp()
				}
				{
					p.SetState(158)
					p.columnOp(4)
				}

			case 2:
				localctx = NewColumnNameNumericOpContext(p, NewColumnOpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_columnOp)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(161)
					p.NumericOp()
				}
				{
					p.SetState(162)
					p.numericExp(0)
				}

			}

		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// ISignedNumberContext is an interface to support dynamic dispatch.
type ISignedNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSignedNumberContext differentiates from other interfaces.
	IsSignedNumberContext()
}

type SignedNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignedNumberContext() *SignedNumberContext {
	var p = new(SignedNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_signedNumber
	return p
}

func (*SignedNumberContext) IsSignedNumberContext() {}

func NewSignedNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignedNumberContext {
	var p = new(SignedNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_signedNumber

	return p
}

func (s *SignedNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *SignedNumberContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(TSLParserNUMERIC_LITERAL, 0)
}

func (s *SignedNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignedNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignedNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterSignedNumber(s)
	}
}

func (s *SignedNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitSignedNumber(s)
	}
}

func (p *TSLParser) SignedNumber() (localctx ISignedNumberContext) {
	localctx = NewSignedNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TSLParserRULE_signedNumber)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TSLParserT__13 || _la == TSLParserT__14 {
		{
			p.SetState(169)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TSLParserT__13 || _la == TSLParserT__14) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(172)
		p.Match(TSLParserNUMERIC_LITERAL)
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(TSLParserSTRING_LITERAL, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *TSLParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TSLParserRULE_stringValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(TSLParserSTRING_LITERAL)
	}

	return localctx
}

// ILiteralValueContext is an interface to support dynamic dispatch.
type ILiteralValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralValueContext differentiates from other interfaces.
	IsLiteralValueContext()
}

type LiteralValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralValueContext() *LiteralValueContext {
	var p = new(LiteralValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_literalValue
	return p
}

func (*LiteralValueContext) IsLiteralValueContext() {}

func NewLiteralValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralValueContext {
	var p = new(LiteralValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_literalValue

	return p
}

func (s *LiteralValueContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralValueContext) SignedNumber() ISignedNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedNumberContext)
}

func (s *LiteralValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterLiteralValue(s)
	}
}

func (s *LiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitLiteralValue(s)
	}
}

func (p *TSLParser) LiteralValue() (localctx ILiteralValueContext) {
	localctx = NewLiteralValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TSLParserRULE_literalValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserT__13, TSLParserT__14, TSLParserNUMERIC_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.SignedNumber()
		}

	case TSLParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.StringValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKeyNotContext is an interface to support dynamic dispatch.
type IKeyNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyNotContext differentiates from other interfaces.
	IsKeyNotContext()
}

type KeyNotContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyNotContext() *KeyNotContext {
	var p = new(KeyNotContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_keyNot
	return p
}

func (*KeyNotContext) IsKeyNotContext() {}

func NewKeyNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyNotContext {
	var p = new(KeyNotContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_keyNot

	return p
}

func (s *KeyNotContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyNotContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(TSLParserK_NOT, 0)
}

func (s *KeyNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterKeyNot(s)
	}
}

func (s *KeyNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitKeyNot(s)
	}
}

func (p *TSLParser) KeyNot() (localctx IKeyNotContext) {
	localctx = NewKeyNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TSLParserRULE_keyNot)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(TSLParserK_NOT)
	}

	return localctx
}

func (p *TSLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 8:
		var t *NumericExpContext = nil
		if localctx != nil {
			t = localctx.(*NumericExpContext)
		}
		return p.NumericExp_Sempred(t, predIndex)

	case 9:
		var t *ColumnOpContext = nil
		if localctx != nil {
			t = localctx.(*ColumnOpContext)
		}
		return p.ColumnOp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TSLParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TSLParser) NumericExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *TSLParser) ColumnOp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
