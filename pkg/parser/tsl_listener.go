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

	// EnterDatabaseName is called when entering the databaseName production.
	EnterDatabaseName(c *DatabaseNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterSignedNumber is called when entering the signedNumber production.
	EnterSignedNumber(c *SignedNumberContext)

	// EnterStringValue is called when entering the stringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterLiteralValue is called when entering the literalValue production.
	EnterLiteralValue(c *LiteralValueContext)

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

	// ExitDatabaseName is called when exiting the databaseName production.
	ExitDatabaseName(c *DatabaseNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitSignedNumber is called when exiting the signedNumber production.
	ExitSignedNumber(c *SignedNumberContext)

	// ExitStringValue is called when exiting the stringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitLiteralValue is called when exiting the literalValue production.
	ExitLiteralValue(c *LiteralValueContext)

	// ExitKeyNot is called when exiting the keyNot production.
	ExitKeyNot(c *KeyNotContext)
}
