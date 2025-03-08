/* A Bison parser, made by GNU Bison 3.8.2.  */

/* Bison interface for Yacc-like parsers in C

   Copyright (C) 1984, 1989-1990, 2000-2015, 2018-2021 Free Software Foundation,
   Inc.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.  */

/* As a special exception, you may create a larger work that contains
   part or all of the Bison parser skeleton and distribute that work
   under terms of your choice, so long as that work isn't itself a
   parser generator using the skeleton or a modified version thereof
   as a parser skeleton.  Alternatively, if you modify or redistribute
   the parser skeleton itself, you may (at your option) remove this
   special exception, which will cause the skeleton and the resulting
   Bison output files to be licensed under the GNU General Public
   License without this special exception.

   This special exception was added by the Free Software Foundation in
   version 2.2 of Bison.  */

/* DO NOT RELY ON FEATURES THAT ARE NOT DOCUMENTED in the manual,
   especially those whose name start with YY_ or yy_.  They are
   private implementation details that can be changed or removed.  */

#ifndef YY_YY_TSL_TSL_PARSER_TAB_H_INCLUDED
# define YY_YY_TSL_TSL_PARSER_TAB_H_INCLUDED
/* Debug traces.  */
#ifndef YYDEBUG
# define YYDEBUG 0
#endif
#if YYDEBUG
extern int yydebug;
#endif

/* Token kinds.  */
#ifndef YYTOKENTYPE
# define YYTOKENTYPE
  enum yytokentype
  {
    YYEMPTY = -2,
    YYEOF = 0,                     /* "end of file"  */
    YYerror = 256,                 /* error  */
    YYUNDEF = 257,                 /* "invalid token"  */
    K_LIKE = 258,                  /* K_LIKE  */
    K_ILIKE = 259,                 /* K_ILIKE  */
    K_AND = 260,                   /* K_AND  */
    K_OR = 261,                    /* K_OR  */
    K_BETWEEN = 262,               /* K_BETWEEN  */
    K_IN = 263,                    /* K_IN  */
    K_IS = 264,                    /* K_IS  */
    K_NULL = 265,                  /* K_NULL  */
    K_NOT = 266,                   /* K_NOT  */
    K_TRUE = 267,                  /* K_TRUE  */
    K_FALSE = 268,                 /* K_FALSE  */
    LE = 269,                      /* LE  */
    LT = 270,                      /* LT  */
    GE = 271,                      /* GE  */
    GT = 272,                      /* GT  */
    NE = 273,                      /* NE  */
    EQ = 274,                      /* EQ  */
    REQ = 275,                     /* REQ  */
    RNE = 276,                     /* RNE  */
    RFC3339 = 277,                 /* RFC3339  */
    DATE = 278,                    /* DATE  */
    LPAREN = 279,                  /* LPAREN  */
    RPAREN = 280,                  /* RPAREN  */
    COMMA = 281,                   /* COMMA  */
    PLUS = 282,                    /* PLUS  */
    MINUS = 283,                   /* MINUS  */
    STAR = 284,                    /* STAR  */
    SLASH = 285,                   /* SLASH  */
    PERCENT = 286,                 /* PERCENT  */
    LBRACKET = 287,                /* LBRACKET  */
    RBRACKET = 288,                /* RBRACKET  */
    NUMERIC_LITERAL = 289,         /* NUMERIC_LITERAL  */
    STRING_LITERAL = 290,          /* STRING_LITERAL  */
    IDENTIFIER = 291               /* IDENTIFIER  */
  };
  typedef enum yytokentype yytoken_kind_t;
#endif

/* Value type.  */
#if ! defined YYSTYPE && ! defined YYSTYPE_IS_DECLARED
union YYSTYPE
{
#line 36 "tsl_parser.y"

    ast_node *node;
    double num;
    char *str;

#line 106 "../tsl/tsl_parser.tab.h"

};
typedef union YYSTYPE YYSTYPE;
# define YYSTYPE_IS_TRIVIAL 1
# define YYSTYPE_IS_DECLARED 1
#endif


extern YYSTYPE yylval;


int yyparse (void);


#endif /* !YY_YY_TSL_TSL_PARSER_TAB_H_INCLUDED  */
