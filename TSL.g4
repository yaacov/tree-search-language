// TSL.g4
grammar TSL;

// Rules
start : expr EOF;

expr
  : columnOp literalOp literalValue                                           # LiteralOps
  | columnOp stringOp stringValue                                             # StringOps
  | columnOp keyNot? K_LIKE stringValue                                       # Like
  | columnOp K_IS keyNot? K_NULL                                              # IsNull
  | columnOp K_IS keyNot? literalValue                                        # IsLiteral
  | columnOp keyNot? K_BETWEEN literalValue K_AND literalValue                # Between
  | columnOp keyNot? K_IN ( '(' ( literalValue ( ',' literalValue )* )? ')' ) # In
  | K_NOT expr                                                                # Not
  | expr K_AND expr                                                           # And
  | expr K_OR expr                                                            # Or
  | '(' expr ')'                                                              # Par
  ;

literalOp
  : ( '<' | '<=' | '>' | '>=' )
  | ( '=' | '!=' | '<>' )
  ;

numericOp
  : ( '+' | '-' | '*' | '/' | '%' )
  ;

stringOp
  : ( '~=' | '~!' )
  ;

databaseName
  : IDENTIFIER
  ;

tableName
  : IDENTIFIER
  ;

columnName
  : ( ( databaseName '.' )? tableName '.' )? IDENTIFIER
  ;

columnOp
  : columnName                  # Column
  | columnOp numericOp columnOp # ColumnNameOp
  | '(' columnOp ')'            # ColumnNamePar
  ;

signedNumber
  : ( '+' | '-' )? NUMERIC_LITERAL
  ;

stringValue
  : STRING_LITERAL
  ;

literalValue
  : signedNumber
  | stringValue
  ;

keyNot
 : K_NOT
 ;

// Words
K_LIKE : L I K E;
K_AND : A N D;
K_OR : O R;
K_BETWEEN : B E T W E E N;
K_IN : I N;
K_IS : I S;
K_NULL : N U L L;
K_NOT : N O T;

IDENTIFIER
  : '"' (~'"' | '""')* '"'
  | '`' (~'`' | '``')* '`'
  | '[' ~']'* ']'
  | [a-zA-Z_] [a-zA-Z_0-9]* // TODO check: needs more chars in set
  ;

NUMERIC_LITERAL
  : DIGIT+ ( '.' DIGIT* )? ( E [-+]? DIGIT+ )?
  | '.' DIGIT+ ( E [-+]? DIGIT+ )?
  ;

STRING_LITERAL
  : '\'' ( ~'\'' | '\'\'' )* '\''
  ;

SPACES
  : [ \u000B\t\r\n] -> channel(HIDDEN)
  ;

// Fragments
fragment DIGIT : [0-9];

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];
