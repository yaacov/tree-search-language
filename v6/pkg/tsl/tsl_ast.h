#ifndef TSL_AST_H
#define TSL_AST_H

#ifdef __cplusplus
extern "C" {
#endif

/* AST node types */
typedef enum {
    AST_NUMBER,
    AST_STRING,
    AST_IDENTIFIER,
    AST_BINARY_OP,
    AST_UNARY_OP,
    AST_DATE,
    AST_RFC3339,
    AST_ARRAY,       
    AST_BOOL,        
    AST_NULL       
} ast_node_type;

/* AST node structure */
typedef struct ast_node {
    ast_node_type type;
    union {
        double number;       /* for numeric literals */
        char *string;        /* for string literals and identifiers */
        int boolean;         /* for boolean value */
        struct {
            int op;          /* operator token (e.g., PLUS, MINUS, etc.) */
            struct ast_node *left;
            struct ast_node *right;
        } binary;
        struct {
            int op;          /* operator token (e.g., MINUS for unary minus, K_NOT for NOT) */
            struct ast_node *child;
        } unary;
        struct {
            int size;
            struct ast_node **elements;
        } array;
    } data;
} ast_node;

/* AST node creation functions */
ast_node *ast_create_null(void);
ast_node *ast_create_boolean(int value);
ast_node *ast_create_number(const char *value_str);
ast_node *ast_create_string(const char *value);
ast_node *ast_create_identifier(const char *value);
ast_node *ast_create_binary(int op, ast_node *left, ast_node *right);
ast_node *ast_create_unary(int op, ast_node *child);
ast_node *ast_create_date(const char *value);
ast_node *ast_create_rfc3339(const char *value);
ast_node *ast_create_array(int size, ast_node **elements);
ast_node *ast_clone(const ast_node *node);

/* AST node manipulation functions */
ast_node *ast_detach_binary_left(ast_node *node);
ast_node *ast_detach_binary_right(ast_node *node);
ast_node *ast_detach_unary_child(ast_node *node);
int ast_attach_binary_left(ast_node *node, ast_node *child);
int ast_attach_binary_right(ast_node *node, ast_node *child);
int ast_attach_unary_child(ast_node *node, ast_node *child);

/* AST printing functions */
void ast_print(const ast_node *node);
void ast_print_indent(const ast_node *node, int indent);

/* Error position functions */
void set_error_position(int pos);
int get_error_position(void);

#ifdef __cplusplus
}
#endif

#endif /* TSL_AST_H */
