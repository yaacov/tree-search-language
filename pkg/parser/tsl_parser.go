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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 29, 139,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 5, 3, 39, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 47, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 54, 10, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 60, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 69,
	10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 76, 10, 3, 12, 3, 14, 3, 79,
	11, 3, 5, 3, 81, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 91, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 99, 10, 3, 12,
	3, 14, 3, 102, 11, 3, 3, 4, 3, 4, 5, 4, 106, 10, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 117, 10, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 122, 10, 8, 3, 8, 3, 8, 3, 9, 5, 9, 127, 10, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 5, 11, 135, 10, 11, 3, 12, 3, 12, 3, 12, 2, 3, 4, 13,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 6, 3, 2, 6, 9, 3, 2, 10, 12,
	3, 2, 13, 14, 3, 2, 16, 17, 2, 149, 2, 24, 3, 2, 2, 2, 4, 90, 3, 2, 2,
	2, 6, 105, 3, 2, 2, 2, 8, 107, 3, 2, 2, 2, 10, 109, 3, 2, 2, 2, 12, 111,
	3, 2, 2, 2, 14, 121, 3, 2, 2, 2, 16, 126, 3, 2, 2, 2, 18, 130, 3, 2, 2,
	2, 20, 134, 3, 2, 2, 2, 22, 136, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26,
	7, 2, 2, 3, 26, 3, 3, 2, 2, 2, 27, 28, 8, 3, 1, 2, 28, 29, 5, 14, 8, 2,
	29, 30, 5, 6, 4, 2, 30, 31, 5, 20, 11, 2, 31, 91, 3, 2, 2, 2, 32, 33, 5,
	14, 8, 2, 33, 34, 5, 8, 5, 2, 34, 35, 5, 18, 10, 2, 35, 91, 3, 2, 2, 2,
	36, 38, 5, 14, 8, 2, 37, 39, 5, 22, 12, 2, 38, 37, 3, 2, 2, 2, 38, 39,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 7, 18, 2, 2, 41, 42, 5, 18, 10,
	2, 42, 91, 3, 2, 2, 2, 43, 44, 5, 14, 8, 2, 44, 46, 7, 23, 2, 2, 45, 47,
	5, 22, 12, 2, 46, 45, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 3, 2, 2,
	2, 48, 49, 7, 24, 2, 2, 49, 91, 3, 2, 2, 2, 50, 51, 5, 14, 8, 2, 51, 53,
	7, 23, 2, 2, 52, 54, 5, 22, 12, 2, 53, 52, 3, 2, 2, 2, 53, 54, 3, 2, 2,
	2, 54, 55, 3, 2, 2, 2, 55, 56, 5, 20, 11, 2, 56, 91, 3, 2, 2, 2, 57, 59,
	5, 14, 8, 2, 58, 60, 5, 22, 12, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2,
	2, 60, 61, 3, 2, 2, 2, 61, 62, 7, 21, 2, 2, 62, 63, 5, 20, 11, 2, 63, 64,
	7, 19, 2, 2, 64, 65, 5, 20, 11, 2, 65, 91, 3, 2, 2, 2, 66, 68, 5, 14, 8,
	2, 67, 69, 5, 22, 12, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70,
	3, 2, 2, 2, 70, 71, 7, 22, 2, 2, 71, 80, 7, 3, 2, 2, 72, 77, 5, 20, 11,
	2, 73, 74, 7, 4, 2, 2, 74, 76, 5, 20, 11, 2, 75, 73, 3, 2, 2, 2, 76, 79,
	3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2,
	79, 77, 3, 2, 2, 2, 80, 72, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 3,
	2, 2, 2, 82, 83, 7, 5, 2, 2, 83, 91, 3, 2, 2, 2, 84, 85, 7, 25, 2, 2, 85,
	91, 5, 4, 3, 6, 86, 87, 7, 3, 2, 2, 87, 88, 5, 4, 3, 2, 88, 89, 7, 5, 2,
	2, 89, 91, 3, 2, 2, 2, 90, 27, 3, 2, 2, 2, 90, 32, 3, 2, 2, 2, 90, 36,
	3, 2, 2, 2, 90, 43, 3, 2, 2, 2, 90, 50, 3, 2, 2, 2, 90, 57, 3, 2, 2, 2,
	90, 66, 3, 2, 2, 2, 90, 84, 3, 2, 2, 2, 90, 86, 3, 2, 2, 2, 91, 100, 3,
	2, 2, 2, 92, 93, 12, 5, 2, 2, 93, 94, 7, 19, 2, 2, 94, 99, 5, 4, 3, 6,
	95, 96, 12, 4, 2, 2, 96, 97, 7, 20, 2, 2, 97, 99, 5, 4, 3, 5, 98, 92, 3,
	2, 2, 2, 98, 95, 3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2,
	100, 101, 3, 2, 2, 2, 101, 5, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 106,
	9, 2, 2, 2, 104, 106, 9, 3, 2, 2, 105, 103, 3, 2, 2, 2, 105, 104, 3, 2,
	2, 2, 106, 7, 3, 2, 2, 2, 107, 108, 9, 4, 2, 2, 108, 9, 3, 2, 2, 2, 109,
	110, 7, 26, 2, 2, 110, 11, 3, 2, 2, 2, 111, 112, 7, 26, 2, 2, 112, 13,
	3, 2, 2, 2, 113, 114, 5, 10, 6, 2, 114, 115, 7, 15, 2, 2, 115, 117, 3,
	2, 2, 2, 116, 113, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 118, 3, 2, 2,
	2, 118, 119, 5, 12, 7, 2, 119, 120, 7, 15, 2, 2, 120, 122, 3, 2, 2, 2,
	121, 116, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	124, 7, 26, 2, 2, 124, 15, 3, 2, 2, 2, 125, 127, 9, 5, 2, 2, 126, 125,
	3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 7, 27,
	2, 2, 129, 17, 3, 2, 2, 2, 130, 131, 7, 28, 2, 2, 131, 19, 3, 2, 2, 2,
	132, 135, 5, 16, 9, 2, 133, 135, 5, 18, 10, 2, 134, 132, 3, 2, 2, 2, 134,
	133, 3, 2, 2, 2, 135, 21, 3, 2, 2, 2, 136, 137, 7, 25, 2, 2, 137, 23, 3,
	2, 2, 2, 17, 38, 46, 53, 59, 68, 77, 80, 90, 98, 100, 105, 116, 121, 126,
	134,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'<'", "'<='", "'>'", "'>='", "'='", "'!='", "'<>'",
	"'~='", "'~!'", "'.'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "K_LIKE",
	"K_AND", "K_OR", "K_BETWEEN", "K_IN", "K_IS", "K_NULL", "K_NOT", "IDENTIFIER",
	"NUMERIC_LITERAL", "STRING_LITERAL", "SPACES",
}

var ruleNames = []string{
	"start", "expr", "literalOp", "stringOp", "databaseName", "tableName",
	"columnName", "signedNumber", "stringValue", "literalValue", "keyNot",
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
	TSLParserK_LIKE          = 16
	TSLParserK_AND           = 17
	TSLParserK_OR            = 18
	TSLParserK_BETWEEN       = 19
	TSLParserK_IN            = 20
	TSLParserK_IS            = 21
	TSLParserK_NULL          = 22
	TSLParserK_NOT           = 23
	TSLParserIDENTIFIER      = 24
	TSLParserNUMERIC_LITERAL = 25
	TSLParserSTRING_LITERAL  = 26
	TSLParserSPACES          = 27
)

// TSLParser rules.
const (
	TSLParserRULE_start        = 0
	TSLParserRULE_expr         = 1
	TSLParserRULE_literalOp    = 2
	TSLParserRULE_stringOp     = 3
	TSLParserRULE_databaseName = 4
	TSLParserRULE_tableName    = 5
	TSLParserRULE_columnName   = 6
	TSLParserRULE_signedNumber = 7
	TSLParserRULE_stringValue  = 8
	TSLParserRULE_literalValue = 9
	TSLParserRULE_keyNot       = 10
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
		p.SetState(22)
		p.expr(0)
	}
	{
		p.SetState(23)
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

func (s *LikeContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
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

func (s *InContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
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

func (s *IsLiteralContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
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

func (s *BetweenContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
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

func (s *StringOpsContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
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

func (s *IsNullContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
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

func (s *LiteralOpsContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
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
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLiteralOpsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(26)
			p.ColumnName()
		}
		{
			p.SetState(27)
			p.LiteralOp()
		}
		{
			p.SetState(28)
			p.LiteralValue()
		}

	case 2:
		localctx = NewStringOpsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.ColumnName()
		}
		{
			p.SetState(31)
			p.StringOp()
		}
		{
			p.SetState(32)
			p.StringValue()
		}

	case 3:
		localctx = NewLikeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.ColumnName()
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(35)
				p.KeyNot()
			}

		}
		{
			p.SetState(38)
			p.Match(TSLParserK_LIKE)
		}
		{
			p.SetState(39)
			p.StringValue()
		}

	case 4:
		localctx = NewIsNullContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.ColumnName()
		}
		{
			p.SetState(42)
			p.Match(TSLParserK_IS)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(43)
				p.KeyNot()
			}

		}
		{
			p.SetState(46)
			p.Match(TSLParserK_NULL)
		}

	case 5:
		localctx = NewIsLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.ColumnName()
		}
		{
			p.SetState(49)
			p.Match(TSLParserK_IS)
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(50)
				p.KeyNot()
			}

		}
		{
			p.SetState(53)
			p.LiteralValue()
		}

	case 6:
		localctx = NewBetweenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(55)
			p.ColumnName()
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
			p.Match(TSLParserK_BETWEEN)
		}
		{
			p.SetState(60)
			p.LiteralValue()
		}
		{
			p.SetState(61)
			p.Match(TSLParserK_AND)
		}
		{
			p.SetState(62)
			p.LiteralValue()
		}

	case 7:
		localctx = NewInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.ColumnName()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == TSLParserK_NOT {
			{
				p.SetState(65)
				p.KeyNot()
			}

		}
		{
			p.SetState(68)
			p.Match(TSLParserK_IN)
		}

		{
			p.SetState(69)
			p.Match(TSLParserT__0)
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TSLParserT__13)|(1<<TSLParserT__14)|(1<<TSLParserNUMERIC_LITERAL)|(1<<TSLParserSTRING_LITERAL))) != 0 {
			{
				p.SetState(70)
				p.LiteralValue()
			}
			p.SetState(75)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == TSLParserT__1 {
				{
					p.SetState(71)
					p.Match(TSLParserT__1)
				}
				{
					p.SetState(72)
					p.LiteralValue()
				}

				p.SetState(77)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(80)
			p.Match(TSLParserT__2)
		}

	case 8:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(82)
			p.Match(TSLParserK_NOT)
		}
		{
			p.SetState(83)
			p.expr(4)
		}

	case 9:
		localctx = NewParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(TSLParserT__0)
		}
		{
			p.SetState(85)
			p.expr(0)
		}
		{
			p.SetState(86)
			p.Match(TSLParserT__2)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(91)
					p.Match(TSLParserK_AND)
				}
				{
					p.SetState(92)
					p.expr(4)
				}

			case 2:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TSLParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(94)
					p.Match(TSLParserK_OR)
				}
				{
					p.SetState(95)
					p.expr(3)
				}

			}

		}
		p.SetState(100)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserT__3, TSLParserT__4, TSLParserT__5, TSLParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
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
			p.SetState(102)
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
		p.SetState(105)
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
	p.EnterRule(localctx, 8, TSLParserRULE_databaseName)

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
		p.SetState(107)
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
	p.EnterRule(localctx, 10, TSLParserRULE_tableName)

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
		p.SetState(109)
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
	p.EnterRule(localctx, 12, TSLParserRULE_columnName)

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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		p.SetState(114)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(111)
				p.DatabaseName()
			}
			{
				p.SetState(112)
				p.Match(TSLParserT__12)
			}

		}
		{
			p.SetState(116)
			p.TableName()
		}
		{
			p.SetState(117)
			p.Match(TSLParserT__12)
		}

	}
	{
		p.SetState(121)
		p.Match(TSLParserIDENTIFIER)
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
	p.EnterRule(localctx, 14, TSLParserRULE_signedNumber)
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
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TSLParserT__13 || _la == TSLParserT__14 {
		{
			p.SetState(123)
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
		p.SetState(126)
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
	p.EnterRule(localctx, 16, TSLParserRULE_stringValue)

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
		p.SetState(128)
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
	p.EnterRule(localctx, 18, TSLParserRULE_literalValue)

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

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TSLParserT__13, TSLParserT__14, TSLParserNUMERIC_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.SignedNumber()
		}

	case TSLParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
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
	p.EnterRule(localctx, 20, TSLParserRULE_keyNot)

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
		p.SetState(134)
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
