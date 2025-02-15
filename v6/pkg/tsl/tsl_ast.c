#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include "tsl_ast.h"

/* Error position tracking */
static int current_error_position = 0;

void set_error_position(int pos) {
    current_error_position = pos;
}

int get_error_position(void) {
    return current_error_position;
}

ast_node *ast_create_number(const char *value_str) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_NUMBER;

    char *endptr;
    double value = strtod(value_str, &endptr);
    
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
    
    node->data.number = value;
    return node;
}

ast_node *ast_create_string(const char *value) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_STRING;
    node->data.string = strdup(value);
    return node;
}

ast_node *ast_create_identifier(const char *value) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_IDENTIFIER;
    node->data.string = strdup(value);
    return node;
}

ast_node *ast_create_date(const char *value) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_DATE;
    node->data.string = strdup(value);
    return node;
}

ast_node *ast_create_rfc3339(const char *value) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_RFC3339;
    node->data.string = strdup(value);
    return node;
}

ast_node *ast_create_boolean(int value) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_BOOL;
    node->data.boolean = value;
    return node;
}

ast_node *ast_create_null(void) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_NULL;
    return node;
}

ast_node *ast_create_binary(int op, ast_node *left, ast_node *right) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_BINARY_OP;
    node->data.binary.op = op;
    node->data.binary.left = left;
    node->data.binary.right = right;
    return node;
}

ast_node *ast_create_unary(int op, ast_node *child) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_UNARY_OP;
    node->data.unary.op = op;
    node->data.unary.child = child;
    return node;
}

ast_node *ast_create_array(int size, ast_node **elements) {
    ast_node *node = malloc(sizeof(ast_node));
    node->type = AST_ARRAY;
    node->data.array.size = size;
    node->data.array.elements = malloc(size * sizeof(ast_node*));
    for (int i = 0; i < size; i++) {
        node->data.array.elements[i] = elements[i];
    }
    return node;
}

void ast_free(ast_node *node) {
    if (!node) return;
    
    switch (node->type) {
        case AST_STRING:
        case AST_IDENTIFIER:
        case AST_DATE:
        case AST_RFC3339:
            free(node->data.string);
            break;
        case AST_BINARY_OP:
            ast_free(node->data.binary.left);
            ast_free(node->data.binary.right);
            break;
        case AST_UNARY_OP:
            ast_free(node->data.unary.child);
            break;
        case AST_ARRAY:
            for (int i = 0; i < node->data.array.size; i++) {
                ast_free(node->data.array.elements[i]);
            }
            free(node->data.array.elements);
            break;
        default:
            break;
    }
    
    free(node);
}

ast_node *ast_clone(const ast_node *node) {
    if (!node) return NULL;
    
    ast_node *clone = malloc(sizeof(ast_node));
    clone->type = node->type;
    
    switch (node->type) {
        case AST_NUMBER:
            clone->data.number = node->data.number;
            break;
        case AST_BOOL:
            clone->data.boolean = node->data.boolean;
            break;
        case AST_STRING:
        case AST_IDENTIFIER:
        case AST_DATE:
        case AST_RFC3339:
            clone->data.string = strdup(node->data.string);
            break;
        case AST_BINARY_OP:
            clone->data.binary.op = node->data.binary.op;
            clone->data.binary.left = ast_clone(node->data.binary.left);
            clone->data.binary.right = ast_clone(node->data.binary.right);
            break;
        case AST_UNARY_OP:
            clone->data.unary.op = node->data.unary.op;
            clone->data.unary.child = ast_clone(node->data.unary.child);
            break;
        case AST_ARRAY:
            clone->data.array.size = node->data.array.size;
            clone->data.array.elements = malloc(node->data.array.size * sizeof(ast_node*));
            for (int i = 0; i < node->data.array.size; i++) {
                clone->data.array.elements[i] = ast_clone(node->data.array.elements[i]);
            }
            break;
        case AST_NULL:
            // Nothing additional needed for NULL type
            break;
    }
    return clone;
}

double get_number_value(const ast_node *node) {
    return node->data.number;
}

const char* get_string_value(const ast_node *node) {
    switch (node->type) {
        case AST_STRING:
        case AST_IDENTIFIER:
        case AST_DATE:
        case AST_RFC3339:
            return node->data.string;
        default:
            return NULL;
    }
}

int get_binary_op(const ast_node *node) {
    return node->data.binary.op;
}

ast_node* get_binary_left(const ast_node *node) {
    return node->data.binary.left;
}

ast_node* get_binary_right(const ast_node *node) {
    return node->data.binary.right;
}

int get_unary_op(const ast_node *node) {
    return node->data.unary.op;
}

ast_node* get_unary_child(const ast_node *node) {
    return node->data.unary.child;
}

int get_array_size(const ast_node *node) {
    return node->data.array.size;
}

ast_node* get_array_element(const ast_node *node, int index) {
    if (index >= 0 && index < node->data.array.size) {
        return node->data.array.elements[index];
    }
    return NULL;
}

int get_boolean_value(const ast_node *node) {
    return node->data.boolean;
}

ast_node *ast_detach_binary_left(ast_node *node) {
    if (!node || node->type != AST_BINARY_OP) return NULL;
    
    ast_node *left = node->data.binary.left;
    node->data.binary.left = NULL;
    return left;
}

ast_node *ast_detach_binary_right(ast_node *node) {
    if (!node || node->type != AST_BINARY_OP) return NULL;
    
    ast_node *right = node->data.binary.right;
    node->data.binary.right = NULL;
    return right;
}

ast_node *ast_detach_unary_child(ast_node *node) {
    if (!node || node->type != AST_UNARY_OP) return NULL;
    
    ast_node *child = node->data.unary.child;
    node->data.unary.child = NULL;
    return child;
}

int ast_attach_binary_left(ast_node *node, ast_node *child) {
    if (!node || node->type != AST_BINARY_OP) return 0;
    
    if (node->data.binary.left) {
        ast_free(node->data.binary.left);
    }
    node->data.binary.left = child;
    return 1;
}

int ast_attach_binary_right(ast_node *node, ast_node *child) {
    if (!node || node->type != AST_BINARY_OP) return 0;
    
    if (node->data.binary.right) {
        ast_free(node->data.binary.right);
    }
    node->data.binary.right = child;
    return 1;
}

int ast_attach_unary_child(ast_node *node, ast_node *child) {
    if (!node || node->type != AST_UNARY_OP) return 0;
    
    if (node->data.unary.child) {
        ast_free(node->data.unary.child);
    }
    node->data.unary.child = child;
    return 1;
}
