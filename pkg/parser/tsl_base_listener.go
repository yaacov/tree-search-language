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

// EnterStringOp is called when production stringOp is entered.
func (s *BaseTSLListener) EnterStringOp(ctx *StringOpContext) {}

// ExitStringOp is called when production stringOp is exited.
func (s *BaseTSLListener) ExitStringOp(ctx *StringOpContext) {}

// EnterLikeOp is called when production likeOp is entered.
func (s *BaseTSLListener) EnterLikeOp(ctx *LikeOpContext) {}

// ExitLikeOp is called when production likeOp is exited.
func (s *BaseTSLListener) ExitLikeOp(ctx *LikeOpContext) {}

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

// EnterNumberLiteral is called when production NumberLiteral is entered.
func (s *BaseTSLListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production NumberLiteral is exited.
func (s *BaseTSLListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseTSLListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseTSLListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterMathPar is called when production MathPar is entered.
func (s *BaseTSLListener) EnterMathPar(ctx *MathParContext) {}

// ExitMathPar is called when production MathPar is exited.
func (s *BaseTSLListener) ExitMathPar(ctx *MathParContext) {}

// EnterModOps is called when production ModOps is entered.
func (s *BaseTSLListener) EnterModOps(ctx *ModOpsContext) {}

// ExitModOps is called when production ModOps is exited.
func (s *BaseTSLListener) ExitModOps(ctx *ModOpsContext) {}

// EnterSubOps is called when production SubOps is entered.
func (s *BaseTSLListener) EnterSubOps(ctx *SubOpsContext) {}

// ExitSubOps is called when production SubOps is exited.
func (s *BaseTSLListener) ExitSubOps(ctx *SubOpsContext) {}

// EnterMulOps is called when production MulOps is entered.
func (s *BaseTSLListener) EnterMulOps(ctx *MulOpsContext) {}

// ExitMulOps is called when production MulOps is exited.
func (s *BaseTSLListener) ExitMulOps(ctx *MulOpsContext) {}

// EnterDivOps is called when production DivOps is entered.
func (s *BaseTSLListener) EnterDivOps(ctx *DivOpsContext) {}

// ExitDivOps is called when production DivOps is exited.
func (s *BaseTSLListener) ExitDivOps(ctx *DivOpsContext) {}

// EnterColumnIdentifier is called when production ColumnIdentifier is entered.
func (s *BaseTSLListener) EnterColumnIdentifier(ctx *ColumnIdentifierContext) {}

// ExitColumnIdentifier is called when production ColumnIdentifier is exited.
func (s *BaseTSLListener) ExitColumnIdentifier(ctx *ColumnIdentifierContext) {}

// EnterAddOps is called when production AddOps is entered.
func (s *BaseTSLListener) EnterAddOps(ctx *AddOpsContext) {}

// ExitAddOps is called when production AddOps is exited.
func (s *BaseTSLListener) ExitAddOps(ctx *AddOpsContext) {}

// EnterSignedNumber is called when production signedNumber is entered.
func (s *BaseTSLListener) EnterSignedNumber(ctx *SignedNumberContext) {}

// ExitSignedNumber is called when production signedNumber is exited.
func (s *BaseTSLListener) ExitSignedNumber(ctx *SignedNumberContext) {}

// EnterStringValue is called when production stringValue is entered.
func (s *BaseTSLListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production stringValue is exited.
func (s *BaseTSLListener) ExitStringValue(ctx *StringValueContext) {}

// EnterKeyNot is called when production keyNot is entered.
func (s *BaseTSLListener) EnterKeyNot(ctx *KeyNotContext) {}

// ExitKeyNot is called when production keyNot is exited.
func (s *BaseTSLListener) ExitKeyNot(ctx *KeyNotContext) {}
