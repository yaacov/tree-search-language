%{
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
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

/* Bison semantic value type */
%union {
    ast_node *node;
    double num;
    char *str;
}

/* Operator precedence and associativity */
%left K_OR
%left K_AND
%right K_NOT
%left EQ NE LT LE GT GE REQ RNE
%left K_IN K_BETWEEN K_LIKE K_ILIKE
%left K_IS
%right UMINUS
%left STAR SLASH PERCENT    
%left PLUS MINUS         

/* Nonterminal types */
%type <node> input expr primary array_elements array

%start input

%%

input:
    expr { $$ = $1; ast_root = $1; }
    ;

expr:
      expr PLUS expr          { $$ = ast_create_binary(PLUS, $1, $3); }
    | expr MINUS expr         { $$ = ast_create_binary(MINUS, $1, $3); }
    | expr STAR expr          { $$ = ast_create_binary(STAR, $1, $3); }
    | expr SLASH expr         { $$ = ast_create_binary(SLASH, $1, $3); }
    | expr PERCENT expr       { $$ = ast_create_binary(PERCENT, $1, $3); }
    | expr K_BETWEEN expr K_AND expr  { 
                                      ast_node **elements = malloc(2 * sizeof(ast_node*));
                                      elements[0] = $3;
                                      elements[1] = $5;
                                      ast_node *range = ast_create_array(2, elements);
                                      $$ = ast_create_binary(K_BETWEEN, $1, range);
                                      free(elements);
                                    }
    | expr K_NOT K_BETWEEN expr K_AND expr  {
                                      ast_node **elements = malloc(2 * sizeof(ast_node*));
                                      elements[0] = $4;
                                      elements[1] = $6;
                                      ast_node *range = ast_create_array(2, elements);
                                      ast_node *between = ast_create_binary(K_BETWEEN, $1, range);
                                      $$ = ast_create_unary(K_NOT, between);
                                      free(elements);
                                    }
    | expr K_AND expr         { $$ = ast_create_binary(K_AND, $1, $3); }
    | expr K_OR expr          { $$ = ast_create_binary(K_OR, $1, $3); }
    | K_NOT expr              { $$ = ast_create_unary(K_NOT, $2); }
    | MINUS expr %prec UMINUS { $$ = ast_create_unary(MINUS, $2); }
    | expr K_IN array         { $$ = ast_create_binary(K_IN, $1, $3); }
    | expr K_NOT K_IN array   { 
        ast_node *in_node = ast_create_binary(K_IN, $1, $4);
        $$ = ast_create_unary(K_NOT, in_node);
    }
    | expr K_IS K_NULL        { $$ = ast_create_binary(K_IS, $1, ast_create_null()); }
    | expr K_IS K_NOT K_NULL  { 
        ast_node *is_node = ast_create_binary(K_IS, $1, ast_create_null());
        $$ = ast_create_unary(K_NOT, is_node);
    }
    | expr K_LIKE expr          { $$ = ast_create_binary(K_LIKE, $1, $3); }
    | expr K_ILIKE expr         { $$ = ast_create_binary(K_ILIKE, $1, $3); }
    | expr K_NOT K_LIKE expr    { 
                                 ast_node *like_node = ast_create_binary(K_LIKE, $1, $4);
                                 $$ = ast_create_unary(K_NOT, like_node);
                               }
    | expr K_NOT K_ILIKE expr   { 
                                 ast_node *ilike_node = ast_create_binary(K_ILIKE, $1, $4);
                                 $$ = ast_create_unary(K_NOT, ilike_node);
                               }
    | expr EQ expr           { $$ = ast_create_binary(EQ, $1, $3); }
    | expr NE expr           { $$ = ast_create_binary(NE, $1, $3); }
    | expr LT expr           { $$ = ast_create_binary(LT, $1, $3); }
    | expr LE expr           { $$ = ast_create_binary(LE, $1, $3); }
    | expr GT expr           { $$ = ast_create_binary(GT, $1, $3); }
    | expr GE expr           { $$ = ast_create_binary(GE, $1, $3); }
    | expr REQ expr          { $$ = ast_create_binary(REQ, $1, $3); }
    | expr RNE expr          { $$ = ast_create_binary(RNE, $1, $3); }
    | LPAREN expr RPAREN     { $$ = $2; }
    | primary                { $$ = $1; }
    | array                  { $$ = $1; }
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
      NUMERIC_LITERAL       { 
                              char *endptr;
                              double value = strtod($1, &endptr);
                              if (*endptr != '\0') {  // Has suffix
                                  double multiplier = 1024.0;  // Default to binary (Ki)
                                  if (*(endptr + 1) != 'i') {  // If not Ki/Mi/Gi, use decimal K/M/G
                                      multiplier = 1000.0;
                                  }
                                  switch(tolower(*endptr)) {
                                      case 'k': value *= multiplier; break;
                                      case 'm': value *= multiplier * multiplier; break;
                                      case 'g': value *= multiplier * multiplier * multiplier; break;
                                      case 't': value *= multiplier * multiplier * multiplier * multiplier; break;
                                      case 'p': value *= multiplier * multiplier * multiplier * multiplier * multiplier; break;
                                  }
                              }
                              $$ = ast_create_number(value); 
                              free($1); 
                            }
    | STRING_LITERAL          { $$ = ast_create_string($1); free($1); }
    | IDENTIFIER              { $$ = ast_create_identifier($1); free($1); }
    | RFC3339                 { $$ = ast_create_rfc3339($1); free($1); }
    | DATE                    { $$ = ast_create_date($1); free($1); }
    | K_TRUE                  { $$ = ast_create_boolean(1); }
    | K_FALSE                 { $$ = ast_create_boolean(0); }
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
