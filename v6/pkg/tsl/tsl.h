#ifndef TSL_H
#define TSL_H

#ifdef __cplusplus
extern "C" {
#endif

#include "tsl_ast.h"

/* Parse the given input string and return the AST root.
   Returns NULL on failure. */
ast_node *parse_input_string(const char *input);

/* Get the position of the last error */
int get_error_position(void);

/* Get the error string */
const char* get_error_string(void);

/* Error handling functions */
const char* get_input_string_at_error(void);

/* Utility functions to print and free the AST */
void ast_free(ast_node *node);

/* Memory cleanup function */
void cleanup_parser_memory(void);

/* Error handling function */
void yyerror(const char *s);

/* Helper functions to access ast_node data */
double get_number_value(const ast_node *node);
const char* get_string_value(const ast_node *node);
int get_boolean_value(const ast_node *node);
int get_binary_op(const ast_node *node);
ast_node* get_binary_left(const ast_node *node);
ast_node* get_binary_right(const ast_node *node);
int get_unary_op(const ast_node *node);
ast_node* get_unary_child(const ast_node *node);
int get_array_size(const ast_node *node);
ast_node* get_array_element(const ast_node *node, int index);

#ifdef __cplusplus
}
#endif

#endif /* TSL_H */
