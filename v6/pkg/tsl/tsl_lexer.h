#ifndef TSL_LEXER_H
#define TSL_LEXER_H

#include <stdio.h>

int yylex(void);
void yyrestart(FILE *input);
void yyerror(const char *s);

#endif /* TSL_LEXER_H */
