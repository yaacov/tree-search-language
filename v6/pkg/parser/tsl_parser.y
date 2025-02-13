%{
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "tsl.h"
#include "tsl_ast.h"
#include "tsl_lexer.h"

/* Forward declarations */
int yylex(void);
void yyerror(const char *s);
extern int yylineno;
extern int yycolumn;
extern void set_error_position(int pos);

/* Global variable to hold the AST root after parsing */
ast_node *ast_root = NULL;

/* Declarations for Flex helper functions */
extern void *yy_scan_string(const char *);
extern void yy_delete_buffer(void *);

/* API function to parse an input string and return the AST */
ast_node *parse_input_string(const char *input);

/* Global variables for error reporting */
int yycolumn = 1;
int error_pos = 0;
char *input_string = NULL;
char *error_string = NULL;

%}

/* Token definitions */
%token K_LIKE K_ILIKE K_AND K_OR K_BETWEEN K_IN K_IS K_NULL
%token K_NOT K_TRUE K_FALSE
%token LE LT GE GT NE EQ REQ RNE
%token <str> RFC3339 DATE 
%token LPAREN RPAREN COMMA
%token PLUS MINUS STAR SLASH PERCENT
%token LBRACKET RBRACKET
%token <str> NUMERIC_LITERAL STRING_LITERAL IDENTIFIER

/* Operator precedence and associativity */
%left K_OR                         /* lowest precedence */
%left K_AND
%left EQ NE LT LE GT GE REQ RNE
%left K_LIKE K_ILIKE K_IS K_BETWEEN K_IN
%left PLUS MINUS                   /* + - */
%left STAR SLASH PERCENT           /* * / % */
%right K_NOT                       /* unary NOT */

/* Bison semantic value type */
%union {
    ast_node *node;
    double num;
    char *str;
}

/* Nonterminal types */
%type <node> input expr or_expr and_expr comparison_expr
%type <node> additive_expr multiplicative_expr not_expr unary_expr
%type <node> primary array_elements array

%start input

%%

input:
    expr { $$ = $1; ast_root = $1; }
    ;

expr:
    or_expr
    ;

or_expr:
      and_expr
    | or_expr K_OR and_expr        { $$ = ast_create_binary(K_OR, $1, $3); }
    ;

and_expr:
      comparison_expr
    | and_expr K_AND comparison_expr      { $$ = ast_create_binary(K_AND, $1, $3); }
    ;

comparison_expr:
      additive_expr
    | comparison_expr EQ additive_expr      { $$ = ast_create_binary(EQ, $1, $3); }
    | comparison_expr NE additive_expr      { $$ = ast_create_binary(NE, $1, $3); }
    | comparison_expr LT additive_expr      { $$ = ast_create_binary(LT, $1, $3); }
    | comparison_expr LE additive_expr      { $$ = ast_create_binary(LE, $1, $3); }
    | comparison_expr GT additive_expr      { $$ = ast_create_binary(GT, $1, $3); }
    | comparison_expr GE additive_expr      { $$ = ast_create_binary(GE, $1, $3); }
    | comparison_expr REQ additive_expr     { $$ = ast_create_binary(REQ, $1, $3); }
    | comparison_expr RNE additive_expr     { $$ = ast_create_binary(RNE, $1, $3); }
    | comparison_expr K_LIKE additive_expr  { $$ = ast_create_binary(K_LIKE, $1, $3); }
    | comparison_expr K_ILIKE additive_expr { $$ = ast_create_binary(K_ILIKE, $1, $3); }
    | comparison_expr K_NOT K_LIKE additive_expr  { 
        ast_node *like_expr = ast_create_binary(K_LIKE, $1, $4);
        $$ = ast_create_unary(K_NOT, like_expr);
    }
    | comparison_expr K_NOT K_ILIKE additive_expr { 
        ast_node *ilike_expr = ast_create_binary(K_ILIKE, $1, $4);
        $$ = ast_create_unary(K_NOT, ilike_expr);
    }
    | comparison_expr K_IS K_NULL           { $$ = ast_create_binary(K_IS, $1, ast_create_null()); }
    | comparison_expr K_IS K_NOT K_NULL     { 
        ast_node *is_null = ast_create_binary(K_IS, $1, ast_create_null());
        $$ = ast_create_unary(K_NOT, is_null);
    }
    | comparison_expr K_BETWEEN additive_expr K_AND additive_expr {
        ast_node **elements = malloc(2 * sizeof(ast_node*));
        elements[0] = $3;
        elements[1] = $5;
        ast_node *range = ast_create_array(2, elements);
        $$ = ast_create_binary(K_BETWEEN, $1, range);
        free(elements);
    }
    | comparison_expr K_NOT K_BETWEEN additive_expr K_AND additive_expr {
        ast_node **elements = malloc(2 * sizeof(ast_node*));
        elements[0] = $4;
        elements[1] = $6;
        ast_node *range = ast_create_array(2, elements);
        ast_node *between = ast_create_binary(K_BETWEEN, $1, range);
        $$ = ast_create_unary(K_NOT, between);
        free(elements);
    }
    | comparison_expr K_IN array           { $$ = ast_create_binary(K_IN, $1, $3); }
    | comparison_expr K_NOT K_IN array     {
        ast_node *in_expr = ast_create_binary(K_IN, $1, $4);
        $$ = ast_create_unary(K_NOT, in_expr);
    }
    | K_NOT comparison_expr               { $$ = ast_create_unary(K_NOT, $2); }
    ;

additive_expr:
      multiplicative_expr
    | additive_expr PLUS multiplicative_expr   { $$ = ast_create_binary(PLUS, $1, $3); }
    | additive_expr MINUS multiplicative_expr  { $$ = ast_create_binary(MINUS, $1, $3); }
    ;

multiplicative_expr:
      not_expr
    | multiplicative_expr STAR not_expr    { $$ = ast_create_binary(STAR, $1, $3); }
    | multiplicative_expr SLASH not_expr   { $$ = ast_create_binary(SLASH, $1, $3); }
    | multiplicative_expr PERCENT not_expr { $$ = ast_create_binary(PERCENT, $1, $3); }
    ;

not_expr:
      unary_expr
    | K_NOT not_expr               { $$ = ast_create_unary(K_NOT, $2); }
    ;

unary_expr:
      primary
    | MINUS primary                { $$ = ast_create_unary(MINUS, $2); }
    | LPAREN expr RPAREN           { $$ = $2; }
    | array                        { $$ = $1; }
    ;

array:
      LPAREN array_elements RPAREN     { $$ = $2; }
    | LBRACKET array_elements RBRACKET { $$ = $2; }
    ;

array_elements:
      expr                          { 
                                      ast_node **elements = malloc(sizeof(ast_node*));
                                      elements[0] = $1;
                                      $$ = ast_create_array(1, elements);
                                      free(elements);
                                    }
    | array_elements COMMA expr     {
                                      ast_node **elements = malloc(($1->data.array.size + 1) * sizeof(ast_node*));
                                      // Copy existing elements
                                      for(int i = 0; i < $1->data.array.size; i++) {
                                          elements[i] = ast_clone($1->data.array.elements[i]);
                                      }
                                      elements[$1->data.array.size] = ast_clone($3);
                                      $$ = ast_create_array($1->data.array.size + 1, elements);
                                      free(elements);
                                      ast_free($1);
                                    }
    ;

primary:
      NUMERIC_LITERAL       { $$ = ast_create_number($1); free($1); }
    | STRING_LITERAL        { $$ = ast_create_string($1); free($1); }
    | IDENTIFIER            { $$ = ast_create_identifier($1); free($1); }
    | RFC3339               { $$ = ast_create_rfc3339($1); free($1); }
    | DATE                  { $$ = ast_create_date($1); free($1); }
    | K_TRUE                { $$ = ast_create_boolean(1); }
    | K_FALSE               { $$ = ast_create_boolean(0); }
    ;

%%

/* API function implementation */
ast_node *parse_input_string(const char *input) {
    void *buffer = yy_scan_string(input);
    ast_root = NULL;
    
    // Store input for error reporting
    if (input_string) {
        free(input_string);
    }
    input_string = strdup(input);
    
    if (error_string) {
        free(error_string);
        error_string = NULL;
    }
    
    yycolumn = 1;  // Reset column counter
    error_pos = 0;
    
    int result = yyparse();
    yy_delete_buffer(buffer);
    
    ast_node *return_node = NULL;
    if (result == 0) {
        return_node = ast_root;
    } else if (ast_root) {
        ast_free(ast_root);
    }
    
    // Don't free input_string here - it might still be needed for error reporting
    return return_node;
}

// Add cleanup function
void cleanup_parser_memory(void) {
    if (input_string) {
        free(input_string);
        input_string = NULL;
    }
    if (error_string) {
        free(error_string);
        error_string = NULL;
    }
}

/* Enhanced error handling function */
void yyerror(const char *s) {
    set_error_position(yycolumn - 2);  // Subtract 2 to account for token start
    if (error_string) {
        free(error_string);
    }
    error_string = strdup(s);
}

/* Error string getter */
const char* get_error_string(void) {
    return error_string ? error_string : "Unknown error";
}

/* Input string getter */
const char* get_input_string_at_error(void) {
    if (!input_string || error_pos <= 0) {
        return "";
    }
    
    static char context[512];
    int len = strlen(input_string);
    int i = error_pos - 1;
    
    // Skip whitespace backwards to find token start
    while (i > 0 && isspace(input_string[i])) i--;
    
    // Find token start
    int start = i;
    while (start > 0 && !isspace(input_string[start-1])) start--;
    
    // Find token end
    int end = error_pos;
    while (end < len && !isspace(input_string[end])) end++;
    
    // Extract the token
    snprintf(context, sizeof(context), "%.*s", 
             end - start, input_string + start);
    
    return context;
}
