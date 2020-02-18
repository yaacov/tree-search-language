// Code generated from ./TSL.g4 by ANTLR 4.7.1. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 32, 169,
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
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 126, 10, 10, 3, 11, 3, 11, 5, 11,
	130, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 138, 10,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 155, 10, 12, 12, 12, 14, 12,
	158, 11, 12, 3, 13, 5, 13, 161, 10, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 2, 4, 4, 22, 16, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 2, 7, 3, 2, 6, 9, 3, 2, 10, 12, 3, 2, 13, 14, 3, 2, 20,
	21, 3, 2, 18, 19, 2, 181, 2, 30, 3, 2, 2, 2, 4, 96, 3, 2, 2, 2, 6, 111,
	3, 2, 2, 2, 8, 113, 3, 2, 2, 2, 10, 115, 3, 2, 2, 2, 12, 117, 3, 2, 2,
	2, 14, 119, 3, 2, 2, 2, 16, 121, 3, 2, 2, 2, 18, 125, 3, 2, 2, 2, 20, 129,
	3, 2, 2, 2, 22, 137, 3, 2, 2, 2, 24, 160, 3, 2, 2, 2, 26, 164, 3, 2, 2,
	2, 28, 166, 3, 2, 2, 2, 30, 31, 5, 4, 3, 2, 31, 32, 7, 2, 2, 3, 32, 3,
	3, 2, 2, 2, 33, 34, 8, 3, 1, 2, 34, 35, 5, 22, 12, 2, 35, 36, 5, 6, 4,
	2, 36, 37, 5, 20, 11, 2, 37, 97, 3, 2, 2, 2, 38, 39, 5, 22, 12, 2, 39,
	40, 5, 8, 5, 2, 40, 41, 5, 20, 11, 2, 41, 97, 3, 2, 2, 2, 42, 44, 5, 22,
	12, 2, 43, 45, 5, 28, 15, 2, 44, 43, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45,
	46, 3, 2, 2, 2, 46, 47, 5, 10, 6, 2, 47, 48, 5, 18, 10, 2, 48, 97, 3, 2,
	2, 2, 49, 50, 5, 22, 12, 2, 50, 52, 7, 26, 2, 2, 51, 53, 5, 28, 15, 2,
	52, 51, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 7,
	27, 2, 2, 55, 97, 3, 2, 2, 2, 56, 57, 5, 22, 12, 2, 57, 59, 7, 26, 2, 2,
	58, 60, 5, 28, 15, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61, 3,
	2, 2, 2, 61, 62, 5, 18, 10, 2, 62, 97, 3, 2, 2, 2, 63, 65, 5, 22, 12, 2,
	64, 66, 5, 28, 15, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 3,
	2, 2, 2, 67, 68, 7, 24, 2, 2, 68, 69, 5, 18, 10, 2, 69, 70, 7, 22, 2, 2,
	70, 71, 5, 18, 10, 2, 71, 97, 3, 2, 2, 2, 72, 74, 5, 22, 12, 2, 73, 75,
	5, 28, 15, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 76, 3, 2, 2,
	2, 76, 77, 7, 25, 2, 2, 77, 86, 7, 3, 2, 2, 78, 83, 5, 18, 10, 2, 79, 80,
	7, 4, 2, 2, 80, 82, 5, 18, 10, 2, 81, 79, 3, 2, 2, 2, 82, 85, 3, 2, 2,
	2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83,
	3, 2, 2, 2, 86, 78, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2,
	88, 89, 7, 5, 2, 2, 89, 97, 3, 2, 2, 2, 90, 91, 7, 28, 2, 2, 91, 97, 5,
	4, 3, 6, 92, 93, 7, 3, 2, 2, 93, 94, 5, 4, 3, 2, 94, 95, 7, 5, 2, 2, 95,
	97, 3, 2, 2, 2, 96, 33, 3, 2, 2, 2, 96, 38, 3, 2, 2, 2, 96, 42, 3, 2, 2,
	2, 96, 49, 3, 2, 2, 2, 96, 56, 3, 2, 2, 2, 96, 63, 3, 2, 2, 2, 96, 72,
	3, 2, 2, 2, 96, 90, 3, 2, 2, 2, 96, 92, 3, 2, 2, 2, 97, 106, 3, 2, 2, 2,
	98, 99, 12, 5, 2, 2, 99, 100, 7, 22, 2, 2, 100, 105, 5, 4, 3, 6, 101, 102,
	12, 4, 2, 2, 102, 103, 7, 23, 2, 2, 103, 105, 5, 4, 3, 5, 104, 98, 3, 2,
	2, 2, 104, 101, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 5, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 112,
	9, 2, 2, 2, 110, 112, 9, 3, 2, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2,
	2, 2, 112, 7, 3, 2, 2, 2, 113, 114, 9, 4, 2, 2, 114, 9, 3, 2, 2, 2, 115,
	116, 9, 5, 2, 2, 116, 11, 3, 2, 2, 2, 117, 118, 7, 29, 2, 2, 118, 13, 3,
	2, 2, 2, 119, 120, 7, 29, 2, 2, 120, 15, 3, 2, 2, 2, 121, 122, 7, 29, 2,
	2, 122, 17, 3, 2, 2, 2, 123, 126, 5, 24, 13, 2, 124, 126, 5, 26, 14, 2,
	125, 123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 19, 3, 2, 2, 2, 127, 130,
	5, 18, 10, 2, 128, 130, 5, 22, 12, 2, 129, 127, 3, 2, 2, 2, 129, 128, 3,
	2, 2, 2, 130, 21, 3, 2, 2, 2, 131, 132, 8, 12, 1, 2, 132, 138, 5, 16, 9,
	2, 133, 134, 7, 3, 2, 2, 134, 135, 5, 22, 12, 2, 135, 136, 7, 5, 2, 2,
	136, 138, 3, 2, 2, 2, 137, 131, 3, 2, 2, 2, 137, 133, 3, 2, 2, 2, 138,
	156, 3, 2, 2, 2, 139, 140, 12, 8, 2, 2, 140, 141, 7, 15, 2, 2, 141, 155,
	5, 20, 11, 2, 142, 143, 12, 7, 2, 2, 143, 144, 7, 16, 2, 2, 144, 155, 5,
	20, 11, 2, 145, 146, 12, 6, 2, 2, 146, 147, 7, 17, 2, 2, 147, 155, 5, 20,
	11, 2, 148, 149, 12, 5, 2, 2, 149, 150, 7, 18, 2, 2, 150, 155, 5, 20, 11,
	2, 151, 152, 12, 4, 2, 2, 152, 153, 7, 19, 2, 2, 153, 155, 5, 20, 11, 2,
	154, 139, 3, 2, 2, 2, 154, 142, 3, 2, 2, 2, 154, 145, 3, 2, 2, 2, 154,
	148, 3, 2, 2, 2, 154, 151, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154,
	3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 23, 3, 2, 2, 2, 158, 156, 3, 2,
	2, 2, 159, 161, 9, 6, 2, 2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2,
	161, 162, 3, 2, 2, 2, 162, 163, 7, 30, 2, 2, 163, 25, 3, 2, 2, 2, 164,
	165, 7, 31, 2, 2, 165, 27, 3, 2, 2, 2, 166, 167, 7, 28, 2, 2, 167, 29,
	3, 2, 2, 2, 19, 44, 52, 59, 65, 74, 83, 86, 96, 104, 106, 111, 125, 129,
	137, 154, 156, 160,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'<>'",
	"'~='", "'~!'", "'*'", "'/'", "'%'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"K_LIKE", "K_ILIKE", "K_AND", "K_OR", "K_BETWEEN", "K_IN", "K_IS", "K_NULL",
	"K_NOT", "IDENTIFIER", "NUMERIC_LITERAL", "STRING_LITERAL", "SPACES",
}

var ruleNames = []string{
	"start", "expr", "literalOp", "stringOp", "likeOp", "databaseName", "tableName",
	"columnName", "literalValue", "literalOrMath", "mathExp", "signedNumber",
	"stringValue", "keyNot",
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
	TSLParserK_LIKE          = 18
	TSLParserK_ILIKE         = 19
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
	TSLParserRULE_start         = 0
	TSLParserRULE_expr          = 1
	TSLParserRULE_literalOp     = 2
	TSLParserRULE_stringOp      = 3
	TSLParserRULE_likeOp        = 4
	TSLParserRULE_databaseName  = 5
	TSLParserRULE_tableName     = 6
	TSLParserRULE_columnName    = 7
	TSLParserRULE_literalValue  = 8
	TSLParserRULE_literalOrMath = 9
	TSLParserRULE_mathExp       = 10
	TSLParserRULE_signedNumber  = 11
	TSLParserRULE_stringValue   = 12
	TSLParserRULE_keyNot        = 13
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

func (s *LikeContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *LikeContext) LikeOp() ILikeOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeOpContext)
}

func (s *LikeContext) LiteralValue() ILiteralValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
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

func (s *InContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
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

func (s *IsLiteralContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
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

func (s *BetweenContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
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

func (s *StringOpsContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *StringOpsContext) StringOp() IStringOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringOpContext)
}

func (s *StringOpsContext) LiteralOrMath() ILiteralOrMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOrMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOrMathContext)
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

func (s *IsNullContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
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

func (s *LiteralOpsContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *LiteralOpsContext) LiteralOp() ILiteralOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOpContext)
}

func (s *LiteralOpsContext) LiteralOrMath() ILiteralOrMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOrMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOrMathContext)
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
			p.mathExp(0)
		}
		{
			p.SetState(33)
			p.LiteralOp()
		}
		{
			p.SetState(34)
			p.LiteralOrMath()
		}

	case 2:
		localctx = NewStringOpsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(36)
			p.mathExp(0)
		}
		{
			p.SetState(37)
			p.StringOp()
		}
		{
			p.SetState(38)
			p.LiteralOrMath()
		}

	case 3:
		localctx = NewLikeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.mathExp(0)
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
			p.LikeOp()
		}
		{
			p.SetState(45)
			p.LiteralValue()
		}

	case 4:
		localctx = NewIsNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.mathExp(0)
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
			p.mathExp(0)
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
			p.mathExp(0)
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
			p.mathExp(0)
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

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TSLParserT__15)|(1<<TSLParserT__16)|(1<<TSLParserNUMERIC_LITERAL)|(1<<TSLParserSTRING_LITERAL))) != 0 {
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
	p.EnterRule(localctx, 6, TSLParserRULE_stringOp)
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

		if !(_la == TSLParserT__10 || _la == TSLParserT__11) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILikeOpContext is an interface to support dynamic dispatch.
type ILikeOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLikeOpContext differentiates from other interfaces.
	IsLikeOpContext()
}

type LikeOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLikeOpContext() *LikeOpContext {
	var p = new(LikeOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_likeOp
	return p
}

func (*LikeOpContext) IsLikeOpContext() {}

func NewLikeOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeOpContext {
	var p = new(LikeOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_likeOp

	return p
}

func (s *LikeOpContext) GetParser() antlr.Parser { return s.parser }

func (s *LikeOpContext) K_LIKE() antlr.TerminalNode {
	return s.GetToken(TSLParserK_LIKE, 0)
}

func (s *LikeOpContext) K_ILIKE() antlr.TerminalNode {
	return s.GetToken(TSLParserK_ILIKE, 0)
}

func (s *LikeOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterLikeOp(s)
	}
}

func (s *LikeOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitLikeOp(s)
	}
}

func (p *TSLParser) LikeOp() (localctx ILikeOpContext) {
	localctx = NewLikeOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TSLParserRULE_likeOp)
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

		if !(_la == TSLParserK_LIKE || _la == TSLParserK_ILIKE) {
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
	{
		p.SetState(119)
		p.Match(TSLParserIDENTIFIER)
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

func (s *LiteralValueContext) CopyFrom(ctx *LiteralValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringLiteralContext struct {
	*LiteralValueContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.LiteralValueContext = NewEmptyLiteralValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralValueContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type NumberLiteralContext struct {
	*LiteralValueContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.LiteralValueContext = NewEmptyLiteralValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralValueContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) SignedNumber() ISignedNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISignedNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISignedNumberContext)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (p *TSLParser) LiteralValue() (localctx ILiteralValueContext) {
	localctx = NewLiteralValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TSLParserRULE_literalValue)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserT__15, TSLParserT__16, TSLParserNUMERIC_LITERAL:
		localctx = NewNumberLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.SignedNumber()
		}

	case TSLParserSTRING_LITERAL:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.StringValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralOrMathContext is an interface to support dynamic dispatch.
type ILiteralOrMathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralOrMathContext differentiates from other interfaces.
	IsLiteralOrMathContext()
}

type LiteralOrMathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralOrMathContext() *LiteralOrMathContext {
	var p = new(LiteralOrMathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_literalOrMath
	return p
}

func (*LiteralOrMathContext) IsLiteralOrMathContext() {}

func NewLiteralOrMathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralOrMathContext {
	var p = new(LiteralOrMathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_literalOrMath

	return p
}

func (s *LiteralOrMathContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralOrMathContext) LiteralValue() ILiteralValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralValueContext)
}

func (s *LiteralOrMathContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *LiteralOrMathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralOrMathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralOrMathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterLiteralOrMath(s)
	}
}

func (s *LiteralOrMathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitLiteralOrMath(s)
	}
}

func (p *TSLParser) LiteralOrMath() (localctx ILiteralOrMathContext) {
	localctx = NewLiteralOrMathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TSLParserRULE_literalOrMath)

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

	switch p.GetTokenStream().LA(1) {
	case TSLParserT__15, TSLParserT__16, TSLParserNUMERIC_LITERAL, TSLParserSTRING_LITERAL:
		{
			p.SetState(125)
			p.LiteralValue()
		}

	case TSLParserT__0, TSLParserIDENTIFIER:
		{
			p.SetState(126)
			p.mathExp(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMathExpContext is an interface to support dynamic dispatch.
type IMathExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathExpContext differentiates from other interfaces.
	IsMathExpContext()
}

type MathExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExpContext() *MathExpContext {
	var p = new(MathExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TSLParserRULE_mathExp
	return p
}

func (*MathExpContext) IsMathExpContext() {}

func NewMathExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExpContext {
	var p = new(MathExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TSLParserRULE_mathExp

	return p
}

func (s *MathExpContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExpContext) CopyFrom(ctx *MathExpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MathExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MathParContext struct {
	*MathExpContext
}

func NewMathParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MathParContext {
	var p = new(MathParContext)

	p.MathExpContext = NewEmptyMathExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathExpContext))

	return p
}

func (s *MathParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathParContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *MathParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterMathPar(s)
	}
}

func (s *MathParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitMathPar(s)
	}
}

type ModOpsContext struct {
	*MathExpContext
}

func NewModOpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModOpsContext {
	var p = new(ModOpsContext)

	p.MathExpContext = NewEmptyMathExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathExpContext))

	return p
}

func (s *ModOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModOpsContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *ModOpsContext) LiteralOrMath() ILiteralOrMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOrMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOrMathContext)
}

func (s *ModOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterModOps(s)
	}
}

func (s *ModOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitModOps(s)
	}
}

type SubOpsContext struct {
	*MathExpContext
}

func NewSubOpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubOpsContext {
	var p = new(SubOpsContext)

	p.MathExpContext = NewEmptyMathExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathExpContext))

	return p
}

func (s *SubOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubOpsContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *SubOpsContext) LiteralOrMath() ILiteralOrMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOrMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOrMathContext)
}

func (s *SubOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterSubOps(s)
	}
}

func (s *SubOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitSubOps(s)
	}
}

type MulOpsContext struct {
	*MathExpContext
}

func NewMulOpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulOpsContext {
	var p = new(MulOpsContext)

	p.MathExpContext = NewEmptyMathExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathExpContext))

	return p
}

func (s *MulOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOpsContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *MulOpsContext) LiteralOrMath() ILiteralOrMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOrMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOrMathContext)
}

func (s *MulOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterMulOps(s)
	}
}

func (s *MulOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitMulOps(s)
	}
}

type DivOpsContext struct {
	*MathExpContext
}

func NewDivOpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivOpsContext {
	var p = new(DivOpsContext)

	p.MathExpContext = NewEmptyMathExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathExpContext))

	return p
}

func (s *DivOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivOpsContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *DivOpsContext) LiteralOrMath() ILiteralOrMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOrMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOrMathContext)
}

func (s *DivOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterDivOps(s)
	}
}

func (s *DivOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitDivOps(s)
	}
}

type ColumnIdentifierContext struct {
	*MathExpContext
}

func NewColumnIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnIdentifierContext {
	var p = new(ColumnIdentifierContext)

	p.MathExpContext = NewEmptyMathExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathExpContext))

	return p
}

func (s *ColumnIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnIdentifierContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterColumnIdentifier(s)
	}
}

func (s *ColumnIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitColumnIdentifier(s)
	}
}

type AddOpsContext struct {
	*MathExpContext
}

func NewAddOpsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddOpsContext {
	var p = new(AddOpsContext)

	p.MathExpContext = NewEmptyMathExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathExpContext))

	return p
}

func (s *AddOpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpsContext) MathExp() IMathExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpContext)
}

func (s *AddOpsContext) LiteralOrMath() ILiteralOrMathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralOrMathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralOrMathContext)
}

func (s *AddOpsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.EnterAddOps(s)
	}
}

func (s *AddOpsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TSLListener); ok {
		listenerT.ExitAddOps(s)
	}
}

func (p *TSLParser) MathExp() (localctx IMathExpContext) {
	return p.mathExp(0)
}

func (p *TSLParser) mathExp(_p int) (localctx IMathExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, TSLParserRULE_mathExp, _p)

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
	p.SetState(135)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserIDENTIFIER:
		localctx = NewColumnIdentifierContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(130)
			p.ColumnName()
		}

	case TSLParserT__0:
		localctx = NewMathParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(TSLParserT__0)
		}
		{
			p.SetState(132)
			p.mathExp(0)
		}
		{
			p.SetState(133)
			p.Match(TSLParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulOpsContext(p, NewMathExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_mathExp)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(138)
					p.Match(TSLParserT__12)
				}
				{
					p.SetState(139)
					p.LiteralOrMath()
				}

			case 2:
				localctx = NewDivOpsContext(p, NewMathExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_mathExp)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(141)
					p.Match(TSLParserT__13)
				}
				{
					p.SetState(142)
					p.LiteralOrMath()
				}

			case 3:
				localctx = NewModOpsContext(p, NewMathExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_mathExp)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(144)
					p.Match(TSLParserT__14)
				}
				{
					p.SetState(145)
					p.LiteralOrMath()
				}

			case 4:
				localctx = NewAddOpsContext(p, NewMathExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_mathExp)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(147)
					p.Match(TSLParserT__15)
				}
				{
					p.SetState(148)
					p.LiteralOrMath()
				}

			case 5:
				localctx = NewSubOpsContext(p, NewMathExpContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_mathExp)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(150)
					p.Match(TSLParserT__16)
				}
				{
					p.SetState(151)
					p.LiteralOrMath()
				}

			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, TSLParserRULE_signedNumber)
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
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TSLParserT__15 || _la == TSLParserT__16 {
		{
			p.SetState(157)
			_la = p.GetTokenStream().LA(1)

			if !(_la == TSLParserT__15 || _la == TSLParserT__16) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(160)
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
	p.EnterRule(localctx, 24, TSLParserRULE_stringValue)

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
		p.SetState(162)
		p.Match(TSLParserSTRING_LITERAL)
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
		p.SetState(164)
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

	case 10:
		var t *MathExpContext = nil
		if localctx != nil {
			t = localctx.(*MathExpContext)
		}
		return p.MathExp_Sempred(t, predIndex)

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

func (p *TSLParser) MathExp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
