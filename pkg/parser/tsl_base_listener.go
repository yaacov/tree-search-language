// Code generated from TSL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TSL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTSLListener is a complete listener for a parse tree produced by TSLParser.
type BaseTSLListener struct{}

var _ TSLListener = &BaseTSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTSLListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTSLListener) ExitStart(ctx *StartContext) {}

// EnterPar is called when production Par is entered.
func (s *BaseTSLListener) EnterPar(ctx *ParContext) {}

// ExitPar is called when production Par is exited.
func (s *BaseTSLListener) ExitPar(ctx *ParContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseTSLListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseTSLListener) ExitNot(ctx *NotContext) {}

// EnterLike is called when production Like is entered.
func (s *BaseTSLListener) EnterLike(ctx *LikeContext) {}

// ExitLike is called when production Like is exited.
func (s *BaseTSLListener) ExitLike(ctx *LikeContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseTSLListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseTSLListener) ExitOr(ctx *OrContext) {}

// EnterIn is called when production In is entered.
func (s *BaseTSLListener) EnterIn(ctx *InContext) {}

// ExitIn is called when production In is exited.
func (s *BaseTSLListener) ExitIn(ctx *InContext) {}

// EnterIsLiteral is called when production IsLiteral is entered.
func (s *BaseTSLListener) EnterIsLiteral(ctx *IsLiteralContext) {}

// ExitIsLiteral is called when production IsLiteral is exited.
func (s *BaseTSLListener) ExitIsLiteral(ctx *IsLiteralContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseTSLListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseTSLListener) ExitAnd(ctx *AndContext) {}

// EnterBetween is called when production Between is entered.
func (s *BaseTSLListener) EnterBetween(ctx *BetweenContext) {}

// ExitBetween is called when production Between is exited.
func (s *BaseTSLListener) ExitBetween(ctx *BetweenContext) {}

// EnterStringOps is called when production StringOps is entered.
func (s *BaseTSLListener) EnterStringOps(ctx *StringOpsContext) {}

// ExitStringOps is called when production StringOps is exited.
func (s *BaseTSLListener) ExitStringOps(ctx *StringOpsContext) {}

// EnterIsNull is called when production IsNull is entered.
func (s *BaseTSLListener) EnterIsNull(ctx *IsNullContext) {}

// ExitIsNull is called when production IsNull is exited.
func (s *BaseTSLListener) ExitIsNull(ctx *IsNullContext) {}

// EnterLiteralOps is called when production LiteralOps is entered.
func (s *BaseTSLListener) EnterLiteralOps(ctx *LiteralOpsContext) {}

// ExitLiteralOps is called when production LiteralOps is exited.
func (s *BaseTSLListener) ExitLiteralOps(ctx *LiteralOpsContext) {}

// EnterLiteralOp is called when production literalOp is entered.
func (s *BaseTSLListener) EnterLiteralOp(ctx *LiteralOpContext) {}

// ExitLiteralOp is called when production literalOp is exited.
func (s *BaseTSLListener) ExitLiteralOp(ctx *LiteralOpContext) {}

// EnterNumericOp is called when production numericOp is entered.
func (s *BaseTSLListener) EnterNumericOp(ctx *NumericOpContext) {}

// ExitNumericOp is called when production numericOp is exited.
func (s *BaseTSLListener) ExitNumericOp(ctx *NumericOpContext) {}

// EnterStringOp is called when production stringOp is entered.
func (s *BaseTSLListener) EnterStringOp(ctx *StringOpContext) {}

// ExitStringOp is called when production stringOp is exited.
func (s *BaseTSLListener) ExitStringOp(ctx *StringOpContext) {}

// EnterDatabaseName is called when production databaseName is entered.
func (s *BaseTSLListener) EnterDatabaseName(ctx *DatabaseNameContext) {}

// ExitDatabaseName is called when production databaseName is exited.
func (s *BaseTSLListener) ExitDatabaseName(ctx *DatabaseNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseTSLListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseTSLListener) ExitTableName(ctx *TableNameContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseTSLListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseTSLListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterColumn is called when production Column is entered.
func (s *BaseTSLListener) EnterColumn(ctx *ColumnContext) {}

// ExitColumn is called when production Column is exited.
func (s *BaseTSLListener) ExitColumn(ctx *ColumnContext) {}

// EnterColumnNameOp is called when production ColumnNameOp is entered.
func (s *BaseTSLListener) EnterColumnNameOp(ctx *ColumnNameOpContext) {}

// ExitColumnNameOp is called when production ColumnNameOp is exited.
func (s *BaseTSLListener) ExitColumnNameOp(ctx *ColumnNameOpContext) {}

// EnterColumnNamePar is called when production ColumnNamePar is entered.
func (s *BaseTSLListener) EnterColumnNamePar(ctx *ColumnNameParContext) {}

// ExitColumnNamePar is called when production ColumnNamePar is exited.
func (s *BaseTSLListener) ExitColumnNamePar(ctx *ColumnNameParContext) {}

// EnterSignedNumber is called when production signedNumber is entered.
func (s *BaseTSLListener) EnterSignedNumber(ctx *SignedNumberContext) {}

// ExitSignedNumber is called when production signedNumber is exited.
func (s *BaseTSLListener) ExitSignedNumber(ctx *SignedNumberContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseTSLListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseTSLListener) ExitStringValue(ctx *StringValueContext) {}

// EnterLiteralValue is called when production literalValue is entered.
func (s *BaseTSLListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *BaseTSLListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterKeyNot is called when production keyNot is entered.
func (s *BaseTSLListener) EnterKeyNot(ctx *KeyNotContext) {}

// ExitKeyNot is called when production keyNot is exited.
func (s *BaseTSLListener) ExitKeyNot(ctx *KeyNotContext) {}
