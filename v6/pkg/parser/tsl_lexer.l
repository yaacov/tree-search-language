%{
/* C Lexer for TSL Grammar */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "tsl_ast.h"
#include "tsl_lexer.h"
#include "tsl_parser.tab.h"

/* External variables for error reporting */
extern int yycolumn;
extern int error_pos;

/* Function to return the next token */
#define YY_USER_ACTION \
    error_pos = yycolumn; \
    yycolumn += yyleng;

/* Check if string is a date/time format and return appropriate token */
int check_date_formats(char *s) {
    if (strlen(s) == 10 && s[4] == '-' && s[7] == '-') {
        yylval.str = s;
        return DATE;
    } else if (strstr(s, "T") && (strstr(s, "Z") || strstr(s, "+"))) {
        yylval.str = s;
        return RFC3339;
    }
    return 0;
}

/* Process escaped characters in a string */
char* process_escapes(char* text, int len) {
    char* result = malloc(len + 1);
    int i, j = 0;
    
    for (i = 0; i < len; i++) {
        if (text[i] == '\\' && i + 1 < len) {
            i++; // Skip backslash
            switch (text[i]) {
                case 'n': result[j++] = '\n'; break;
                case 't': result[j++] = '\t'; break;
                case 'r': result[j++] = '\r'; break;
                case 'b': result[j++] = '\b'; break;
                case 'f': result[j++] = '\f'; break;
                case '\\': result[j++] = '\\'; break;
                case '\'': result[j++] = '\''; break;
                case '\"': result[j++] = '\"'; break;
                case '`': result[j++] = '`'; break;
                default: result[j++] = text[i]; break;
            }
        } else {
            result[j++] = text[i];
        }
    }
    result[j] = '\0';
    return result;
}

/* Strip quotes and process string */
char* strip_quotes(char* text, int offset) {
    int len = strlen(text) - (offset * 2);
    char *raw = malloc(len + 1);
    strncpy(raw, text + offset, len);
    raw[len] = '\0';
    
    char *processed = process_escapes(raw, len);
    free(raw);
    return processed;
}
%}

/* Basic building blocks */
DIGIT           [0-9]
LETTER          [a-zA-Z]
UNDERSCORE      [_]

/* Date and time components */
YEAR            {DIGIT}{4}
MONTH           {DIGIT}{2}
DAY             {DIGIT}{2}
HOUR            {DIGIT}{2}
MINUTE          {DIGIT}{2}
SEC             {DIGIT}{2}
SEC_FRAC        \.{DIGIT}+
DATE            {YEAR}-{MONTH}-{DAY}
TIME            {HOUR}:{MINUTE}:{SEC}{SEC_FRAC}?
TIME_OFFSET     [+-]{HOUR}:{MINUTE}
RFC3339         {DATE}T{TIME}(Z|{TIME_OFFSET})

/* Identifier patterns */
ID_START        {LETTER}|{UNDERSCORE}
ARRAY_SUFFIX    (\[[0-9]+\]|\[\*\]|\[[a-zA-Z0-9_./-]+\])
ID_CONT         {LETTER}|{DIGIT}|{UNDERSCORE}|[./]
ID_PART         ({ID_CONT}|{ARRAY_SUFFIX})*
IDENTIFIER      {ID_START}{ID_PART}

/* Numeric patterns */
INTEGER         [+-]?{DIGIT}+
DECIMAL         {INTEGER}\.{DIGIT}*
SCIENTIFIC      ({INTEGER}|{DECIMAL})[Ee][+-]?{DIGIT}+
SIZE_SUFFIX     [kKmMgGtTpP][iI]?
NUMBER_WITH_SUFFIX ({INTEGER}|{DECIMAL}){SIZE_SUFFIX}

%option noinput
%option nounput
%option noyywrap

%%

    /* Keywords (case-insensitive) */
(?i:like)                    { return K_LIKE; }
(?i:ilike)                   { return K_ILIKE; }
(?i:and)                     { return K_AND; }
(?i:or)                      { return K_OR; }
(?i:between)                 { return K_BETWEEN; }
(?i:in)                      { return K_IN; }
(?i:is)                      { return K_IS; }
(?i:null)                    { return K_NULL; }
(?i:not)                     { return K_NOT; }
(?i:true)                    { return K_TRUE; }
(?i:false)                   { return K_FALSE; }
(?i:len)                     { return K_LEN; }
(?i:any)                     { return K_ANY; }
(?i:all)                     { return K_ALL; }
(?i:sum)                     { return K_SUM; }

    /* Comparison operators */
"<="                        { return LE; }
"<"                         { return LT; }
">="                        { return GE; }
">"                         { return GT; }
"!="|"<>"                   { return NE; }
"="                         { return EQ; }
"~="                        { return REQ; }
"~!"                        { return RNE; }

    /* Date and time literals */
{RFC3339}                   { yylval.str = strdup(yytext); return RFC3339; }
{DATE}                      { yylval.str = strdup(yytext); return DATE; }

    /* Numeric literals */
{INTEGER}                   { yylval.str = strdup(yytext); return NUMERIC_LITERAL; }
{DECIMAL}                   { yylval.str = strdup(yytext); return NUMERIC_LITERAL; }
{SCIENTIFIC}                { yylval.str = strdup(yytext); return NUMERIC_LITERAL; }
{NUMBER_WITH_SUFFIX}        { yylval.str = strdup(yytext); return NUMERIC_LITERAL; }

    /* String literals and identifiers */
'([^'\\]|\\.)*'           {
                            char *s = strip_quotes(yytext, 1);
                            int token = check_date_formats(s);
                            if (token) return token;
                            yylval.str = s;
                            return STRING_LITERAL;
                           }

\"([^"\\]|\\.)*\"         {
                            char *s = strip_quotes(yytext, 1);
                            int token = check_date_formats(s);
                            if (token) return token;
                            yylval.str = s;
                            return STRING_LITERAL;
                           }

`([^`\\]|\\.)*`           {
                            char *s = strip_quotes(yytext, 1);
                            int token = check_date_formats(s);
                            if (token) return token;
                            yylval.str = s;
                            return STRING_LITERAL;
                           }

{IDENTIFIER}               { yylval.str = strdup(yytext); return IDENTIFIER; }

    /* Operators and delimiters */
"("                        { return LPAREN; }
")"                        { return RPAREN; }
","                        { return COMMA; }
"+"                        { return PLUS; }
"-"                        { return MINUS; }
"*"                        { return STAR; }
"/"                        { return SLASH; }
"%"                        { return PERCENT; }
"["                        { return LBRACKET; }
"]"                        { return RBRACKET; }

    /* Whitespace handling */
\r?\n                     { yycolumn = 1; /* Reset column on newline */ }
[ \t\r\n\f\v]+            { /* Track column for whitespace */ }

    /* Error handling */
.                         { 
                           fprintf(stderr, "Error: Unknown character '%s' at column %d\n", 
                                   yytext, yycolumn); 
                           return -1; 
                          }

%%
