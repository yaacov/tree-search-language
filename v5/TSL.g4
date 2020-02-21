// TSL.g4
grammar TSL;

// Rules
start : expr EOF;

expr
  : mathExp literalOp literalOrMath                                          # LiteralOps
  | mathExp stringOp literalOrMath                                           # StringOps
  | mathExp keyNot? likeOp literalValue                                      # Like
  | mathExp K_IS keyNot? K_NULL                                              # IsNull
  | mathExp K_IS keyNot? literalValue                                        # IsLiteral
  | mathExp keyNot? K_BETWEEN literalValue K_AND literalValue                # Between
  | mathExp keyNot? K_IN ( '(' ( literalValue ( ',' literalValue )* )? ')' ) # In
  | K_NOT expr                                                               # Not
  | expr K_AND expr                                                          # And
  | expr K_OR expr                                                           # Or
  | '(' expr ')'                                                             # Par
  ;

literalOp
  : ( '<' | '<=' | '>' | '>=' )
  | ( '=' | '!=' | '<>' )
  ;

stringOp
  : ( '~=' | '~!' )
  ;

likeOp
  : ( K_LIKE | K_ILIKE )
  ;

databaseName
  : IDENTIFIER
  ;

tableName
  : IDENTIFIER
  ;

columnName
  : IDENTIFIER
  ;

literalValue
  : signedNumber         # NumberLiteral
  | stringValue          # StringLiteral
  | ( K_TRUE | K_FALSE ) # BooleanLiteral
  ;

literalOrMath
  : ( literalValue | mathExp )
  ;

mathExp
  : columnName                             # ColumnIdentifier
  | mathExp '*' literalOrMath # MulOps
  | mathExp '/' literalOrMath # DivOps
  | mathExp '%' literalOrMath # ModOps
  | mathExp '+' literalOrMath # AddOps
  | mathExp '-' literalOrMath # SubOps
  | '(' mathExp ')'                        # MathPar
  ;

signedNumber
  : ( '+' | '-' )? NUMERIC_LITERAL
  ;

stringValue
  : STRING_LITERAL
  ;

keyNot
 : K_NOT
 ;

// Words
K_LIKE : L I K E;
K_ILIKE : I L I K E;
K_AND : A N D;
K_OR : O R;
K_BETWEEN : B E T W E E N;
K_IN : I N;
K_IS : I S;
K_NULL : N U L L;
K_NOT : N O T;
K_TRUE : T R U E;
K_FALSE : F A L S E;

IDENTIFIER
  : '"' (~'"' | '""')* '"'
  | '`' (~'`' | '``')* '`'
  | '[' ~']'* ']'
  | [a-zA-Z_] [a-zA-Z_./0-9]* [a-zA-Z_0-9]* // TODO check: needs more chars in set
  ;

NUMERIC_LITERAL
  : DIGIT+ ('Ki' | 'Mi' | 'Gi' | 'Ti' | 'Pi')
  | DIGIT+ ('K' | 'M' | 'G' | 'T' | 'P')
  | DIGIT+ ( '.' DIGIT* )? ( E [-+]? DIGIT+ )?
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
