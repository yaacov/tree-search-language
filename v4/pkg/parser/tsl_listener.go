// Code generated from TSL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // TSL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TSLListener is a complete listener for a parse tree produced by TSLParser.
type TSLListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterPar is called when entering the Par production.
	EnterPar(c *ParContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterLike is called when entering the Like production.
	EnterLike(c *LikeContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterIn is called when entering the In production.
	EnterIn(c *InContext)

	// EnterIsLiteral is called when entering the IsLiteral production.
	EnterIsLiteral(c *IsLiteralContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterBetween is called when entering the Between production.
	EnterBetween(c *BetweenContext)

	// EnterStringOps is called when entering the StringOps production.
	EnterStringOps(c *StringOpsContext)

	// EnterIsNull is called when entering the IsNull production.
	EnterIsNull(c *IsNullContext)

	// EnterLiteralOps is called when entering the LiteralOps production.
	EnterLiteralOps(c *LiteralOpsContext)

	// EnterLiteralOp is called when entering the literalOp production.
	EnterLiteralOp(c *LiteralOpContext)

	// EnterStringOp is called when entering the stringOp production.
	EnterStringOp(c *StringOpContext)

	// EnterLikeOp is called when entering the likeOp production.
	EnterLikeOp(c *LikeOpContext)

	// EnterDatabaseName is called when entering the databaseName production.
	EnterDatabaseName(c *DatabaseNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterNumberLiteral is called when entering the NumberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterMathPar is called when entering the MathPar production.
	EnterMathPar(c *MathParContext)

	// EnterModOps is called when entering the ModOps production.
	EnterModOps(c *ModOpsContext)

	// EnterSubOps is called when entering the SubOps production.
	EnterSubOps(c *SubOpsContext)

	// EnterMulOps is called when entering the MulOps production.
	EnterMulOps(c *MulOpsContext)

	// EnterDivOps is called when entering the DivOps production.
	EnterDivOps(c *DivOpsContext)

	// EnterColumnIdentifier is called when entering the ColumnIdentifier production.
	EnterColumnIdentifier(c *ColumnIdentifierContext)

	// EnterAddOps is called when entering the AddOps production.
	EnterAddOps(c *AddOpsContext)

	// EnterSignedNumber is called when entering the signedNumber production.
	EnterSignedNumber(c *SignedNumberContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterKeyNot is called when entering the keyNot production.
	EnterKeyNot(c *KeyNotContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitPar is called when exiting the Par production.
	ExitPar(c *ParContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitLike is called when exiting the Like production.
	ExitLike(c *LikeContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitIn is called when exiting the In production.
	ExitIn(c *InContext)

	// ExitIsLiteral is called when exiting the IsLiteral production.
	ExitIsLiteral(c *IsLiteralContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitBetween is called when exiting the Between production.
	ExitBetween(c *BetweenContext)

	// ExitStringOps is called when exiting the StringOps production.
	ExitStringOps(c *StringOpsContext)

	// ExitIsNull is called when exiting the IsNull production.
	ExitIsNull(c *IsNullContext)

	// ExitLiteralOps is called when exiting the LiteralOps production.
	ExitLiteralOps(c *LiteralOpsContext)

	// ExitLiteralOp is called when exiting the literalOp production.
	ExitLiteralOp(c *LiteralOpContext)

	// ExitStringOp is called when exiting the stringOp production.
	ExitStringOp(c *StringOpContext)

	// ExitLikeOp is called when exiting the likeOp production.
	ExitLikeOp(c *LikeOpContext)

	// ExitDatabaseName is called when exiting the databaseName production.
	ExitDatabaseName(c *DatabaseNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitNumberLiteral is called when exiting the NumberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitMathPar is called when exiting the MathPar production.
	ExitMathPar(c *MathParContext)

	// ExitModOps is called when exiting the ModOps production.
	ExitModOps(c *ModOpsContext)

	// ExitSubOps is called when exiting the SubOps production.
	ExitSubOps(c *SubOpsContext)

	// ExitMulOps is called when exiting the MulOps production.
	ExitMulOps(c *MulOpsContext)

	// ExitDivOps is called when exiting the DivOps production.
	ExitDivOps(c *DivOpsContext)

	// ExitColumnIdentifier is called when exiting the ColumnIdentifier production.
	ExitColumnIdentifier(c *ColumnIdentifierContext)

	// ExitAddOps is called when exiting the AddOps production.
	ExitAddOps(c *AddOpsContext)

	// ExitSignedNumber is called when exiting the signedNumber production.
	ExitSignedNumber(c *SignedNumberContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitKeyNot is called when exiting the keyNot production.
	ExitKeyNot(c *KeyNotContext)
}
