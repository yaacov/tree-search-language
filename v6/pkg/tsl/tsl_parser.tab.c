/* A Bison parser, made by GNU Bison 3.8.2.  */

/* Bison implementation for Yacc-like parsers in C

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

/* C LALR(1) parser skeleton written by Richard Stallman, by
   simplifying the original so-called "semantic" parser.  */

/* DO NOT RELY ON FEATURES THAT ARE NOT DOCUMENTED in the manual,
   especially those whose name start with YY_ or yy_.  They are
   private implementation details that can be changed or removed.  */

/* All symbols defined below should begin with yy or YY, to avoid
   infringing on user name space.  This should be done even for local
   variables, as they might otherwise be expanded by user macros.
   There are some unavoidable exceptions within include files to
   define necessary library symbols; they are noted "INFRINGES ON
   USER NAME SPACE" below.  */

/* Identify Bison output, and Bison version.  */
#define YYBISON 30802

/* Bison version string.  */
#define YYBISON_VERSION "3.8.2"

/* Skeleton name.  */
#define YYSKELETON_NAME "yacc.c"

/* Pure parsers.  */
#define YYPURE 0

/* Push parsers.  */
#define YYPUSH 0

/* Pull parsers.  */
#define YYPULL 1




/* First part of user prologue.  */
#line 1 "tsl_parser.y"

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


#line 105 "../tsl/tsl_parser.tab.c"

# ifndef YY_CAST
#  ifdef __cplusplus
#   define YY_CAST(Type, Val) static_cast<Type> (Val)
#   define YY_REINTERPRET_CAST(Type, Val) reinterpret_cast<Type> (Val)
#  else
#   define YY_CAST(Type, Val) ((Type) (Val))
#   define YY_REINTERPRET_CAST(Type, Val) ((Type) (Val))
#  endif
# endif
# ifndef YY_NULLPTR
#  if defined __cplusplus
#   if 201103L <= __cplusplus
#    define YY_NULLPTR nullptr
#   else
#    define YY_NULLPTR 0
#   endif
#  else
#   define YY_NULLPTR ((void*)0)
#  endif
# endif

#include "tsl_parser.tab.h"
/* Symbol kind.  */
enum yysymbol_kind_t
{
  YYSYMBOL_YYEMPTY = -2,
  YYSYMBOL_YYEOF = 0,                      /* "end of file"  */
  YYSYMBOL_YYerror = 1,                    /* error  */
  YYSYMBOL_YYUNDEF = 2,                    /* "invalid token"  */
  YYSYMBOL_K_LIKE = 3,                     /* K_LIKE  */
  YYSYMBOL_K_ILIKE = 4,                    /* K_ILIKE  */
  YYSYMBOL_K_AND = 5,                      /* K_AND  */
  YYSYMBOL_K_OR = 6,                       /* K_OR  */
  YYSYMBOL_K_BETWEEN = 7,                  /* K_BETWEEN  */
  YYSYMBOL_K_IN = 8,                       /* K_IN  */
  YYSYMBOL_K_IS = 9,                       /* K_IS  */
  YYSYMBOL_K_NULL = 10,                    /* K_NULL  */
  YYSYMBOL_K_NOT = 11,                     /* K_NOT  */
  YYSYMBOL_K_TRUE = 12,                    /* K_TRUE  */
  YYSYMBOL_K_FALSE = 13,                   /* K_FALSE  */
  YYSYMBOL_K_LEN = 14,                     /* K_LEN  */
  YYSYMBOL_K_ANY = 15,                     /* K_ANY  */
  YYSYMBOL_K_ALL = 16,                     /* K_ALL  */
  YYSYMBOL_K_SUM = 17,                     /* K_SUM  */
  YYSYMBOL_RFC3339 = 18,                   /* RFC3339  */
  YYSYMBOL_DATE = 19,                      /* DATE  */
  YYSYMBOL_LPAREN = 20,                    /* LPAREN  */
  YYSYMBOL_RPAREN = 21,                    /* RPAREN  */
  YYSYMBOL_COMMA = 22,                     /* COMMA  */
  YYSYMBOL_PLUS = 23,                      /* PLUS  */
  YYSYMBOL_MINUS = 24,                     /* MINUS  */
  YYSYMBOL_STAR = 25,                      /* STAR  */
  YYSYMBOL_SLASH = 26,                     /* SLASH  */
  YYSYMBOL_PERCENT = 27,                   /* PERCENT  */
  YYSYMBOL_LBRACKET = 28,                  /* LBRACKET  */
  YYSYMBOL_RBRACKET = 29,                  /* RBRACKET  */
  YYSYMBOL_NUMERIC_LITERAL = 30,           /* NUMERIC_LITERAL  */
  YYSYMBOL_STRING_LITERAL = 31,            /* STRING_LITERAL  */
  YYSYMBOL_IDENTIFIER = 32,                /* IDENTIFIER  */
  YYSYMBOL_EQ = 33,                        /* EQ  */
  YYSYMBOL_NE = 34,                        /* NE  */
  YYSYMBOL_LT = 35,                        /* LT  */
  YYSYMBOL_LE = 36,                        /* LE  */
  YYSYMBOL_GT = 37,                        /* GT  */
  YYSYMBOL_GE = 38,                        /* GE  */
  YYSYMBOL_REQ = 39,                       /* REQ  */
  YYSYMBOL_RNE = 40,                       /* RNE  */
  YYSYMBOL_UMINUS = 41,                    /* UMINUS  */
  YYSYMBOL_YYACCEPT = 42,                  /* $accept  */
  YYSYMBOL_input = 43,                     /* input  */
  YYSYMBOL_expr = 44,                      /* expr  */
  YYSYMBOL_or_expr = 45,                   /* or_expr  */
  YYSYMBOL_and_expr = 46,                  /* and_expr  */
  YYSYMBOL_comparison_expr = 47,           /* comparison_expr  */
  YYSYMBOL_additive_expr = 48,             /* additive_expr  */
  YYSYMBOL_multiplicative_expr = 49,       /* multiplicative_expr  */
  YYSYMBOL_not_expr = 50,                  /* not_expr  */
  YYSYMBOL_unary_expr = 51,                /* unary_expr  */
  YYSYMBOL_array = 52,                     /* array  */
  YYSYMBOL_opt_array_elements = 53,        /* opt_array_elements  */
  YYSYMBOL_array_elements = 54,            /* array_elements  */
  YYSYMBOL_primary = 55                    /* primary  */
};
typedef enum yysymbol_kind_t yysymbol_kind_t;




#ifdef short
# undef short
#endif

/* On compilers that do not define __PTRDIFF_MAX__ etc., make sure
   <limits.h> and (if available) <stdint.h> are included
   so that the code can choose integer types of a good width.  */

#ifndef __PTRDIFF_MAX__
# include <limits.h> /* INFRINGES ON USER NAME SPACE */
# if defined __STDC_VERSION__ && 199901 <= __STDC_VERSION__
#  include <stdint.h> /* INFRINGES ON USER NAME SPACE */
#  define YY_STDINT_H
# endif
#endif

/* Narrow types that promote to a signed type and that can represent a
   signed or unsigned integer of at least N bits.  In tables they can
   save space and decrease cache pressure.  Promoting to a signed type
   helps avoid bugs in integer arithmetic.  */

#ifdef __INT_LEAST8_MAX__
typedef __INT_LEAST8_TYPE__ yytype_int8;
#elif defined YY_STDINT_H
typedef int_least8_t yytype_int8;
#else
typedef signed char yytype_int8;
#endif

#ifdef __INT_LEAST16_MAX__
typedef __INT_LEAST16_TYPE__ yytype_int16;
#elif defined YY_STDINT_H
typedef int_least16_t yytype_int16;
#else
typedef short yytype_int16;
#endif

/* Work around bug in HP-UX 11.23, which defines these macros
   incorrectly for preprocessor constants.  This workaround can likely
   be removed in 2023, as HPE has promised support for HP-UX 11.23
   (aka HP-UX 11i v2) only through the end of 2022; see Table 2 of
   <https://h20195.www2.hpe.com/V2/getpdf.aspx/4AA4-7673ENW.pdf>.  */
#ifdef __hpux
# undef UINT_LEAST8_MAX
# undef UINT_LEAST16_MAX
# define UINT_LEAST8_MAX 255
# define UINT_LEAST16_MAX 65535
#endif

#if defined __UINT_LEAST8_MAX__ && __UINT_LEAST8_MAX__ <= __INT_MAX__
typedef __UINT_LEAST8_TYPE__ yytype_uint8;
#elif (!defined __UINT_LEAST8_MAX__ && defined YY_STDINT_H \
       && UINT_LEAST8_MAX <= INT_MAX)
typedef uint_least8_t yytype_uint8;
#elif !defined __UINT_LEAST8_MAX__ && UCHAR_MAX <= INT_MAX
typedef unsigned char yytype_uint8;
#else
typedef short yytype_uint8;
#endif

#if defined __UINT_LEAST16_MAX__ && __UINT_LEAST16_MAX__ <= __INT_MAX__
typedef __UINT_LEAST16_TYPE__ yytype_uint16;
#elif (!defined __UINT_LEAST16_MAX__ && defined YY_STDINT_H \
       && UINT_LEAST16_MAX <= INT_MAX)
typedef uint_least16_t yytype_uint16;
#elif !defined __UINT_LEAST16_MAX__ && USHRT_MAX <= INT_MAX
typedef unsigned short yytype_uint16;
#else
typedef int yytype_uint16;
#endif

#ifndef YYPTRDIFF_T
# if defined __PTRDIFF_TYPE__ && defined __PTRDIFF_MAX__
#  define YYPTRDIFF_T __PTRDIFF_TYPE__
#  define YYPTRDIFF_MAXIMUM __PTRDIFF_MAX__
# elif defined PTRDIFF_MAX
#  ifndef ptrdiff_t
#   include <stddef.h> /* INFRINGES ON USER NAME SPACE */
#  endif
#  define YYPTRDIFF_T ptrdiff_t
#  define YYPTRDIFF_MAXIMUM PTRDIFF_MAX
# else
#  define YYPTRDIFF_T long
#  define YYPTRDIFF_MAXIMUM LONG_MAX
# endif
#endif

#ifndef YYSIZE_T
# ifdef __SIZE_TYPE__
#  define YYSIZE_T __SIZE_TYPE__
# elif defined size_t
#  define YYSIZE_T size_t
# elif defined __STDC_VERSION__ && 199901 <= __STDC_VERSION__
#  include <stddef.h> /* INFRINGES ON USER NAME SPACE */
#  define YYSIZE_T size_t
# else
#  define YYSIZE_T unsigned
# endif
#endif

#define YYSIZE_MAXIMUM                                  \
  YY_CAST (YYPTRDIFF_T,                                 \
           (YYPTRDIFF_MAXIMUM < YY_CAST (YYSIZE_T, -1)  \
            ? YYPTRDIFF_MAXIMUM                         \
            : YY_CAST (YYSIZE_T, -1)))

#define YYSIZEOF(X) YY_CAST (YYPTRDIFF_T, sizeof (X))


/* Stored state numbers (used for stacks). */
typedef yytype_int8 yy_state_t;

/* State numbers in computations.  */
typedef int yy_state_fast_t;

#ifndef YY_
# if defined YYENABLE_NLS && YYENABLE_NLS
#  if ENABLE_NLS
#   include <libintl.h> /* INFRINGES ON USER NAME SPACE */
#   define YY_(Msgid) dgettext ("bison-runtime", Msgid)
#  endif
# endif
# ifndef YY_
#  define YY_(Msgid) Msgid
# endif
#endif


#ifndef YY_ATTRIBUTE_PURE
# if defined __GNUC__ && 2 < __GNUC__ + (96 <= __GNUC_MINOR__)
#  define YY_ATTRIBUTE_PURE __attribute__ ((__pure__))
# else
#  define YY_ATTRIBUTE_PURE
# endif
#endif

#ifndef YY_ATTRIBUTE_UNUSED
# if defined __GNUC__ && 2 < __GNUC__ + (7 <= __GNUC_MINOR__)
#  define YY_ATTRIBUTE_UNUSED __attribute__ ((__unused__))
# else
#  define YY_ATTRIBUTE_UNUSED
# endif
#endif

/* Suppress unused-variable warnings by "using" E.  */
#if ! defined lint || defined __GNUC__
# define YY_USE(E) ((void) (E))
#else
# define YY_USE(E) /* empty */
#endif

/* Suppress an incorrect diagnostic about yylval being uninitialized.  */
#if defined __GNUC__ && ! defined __ICC && 406 <= __GNUC__ * 100 + __GNUC_MINOR__
# if __GNUC__ * 100 + __GNUC_MINOR__ < 407
#  define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN                           \
    _Pragma ("GCC diagnostic push")                                     \
    _Pragma ("GCC diagnostic ignored \"-Wuninitialized\"")
# else
#  define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN                           \
    _Pragma ("GCC diagnostic push")                                     \
    _Pragma ("GCC diagnostic ignored \"-Wuninitialized\"")              \
    _Pragma ("GCC diagnostic ignored \"-Wmaybe-uninitialized\"")
# endif
# define YY_IGNORE_MAYBE_UNINITIALIZED_END      \
    _Pragma ("GCC diagnostic pop")
#else
# define YY_INITIAL_VALUE(Value) Value
#endif
#ifndef YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
# define YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
# define YY_IGNORE_MAYBE_UNINITIALIZED_END
#endif
#ifndef YY_INITIAL_VALUE
# define YY_INITIAL_VALUE(Value) /* Nothing. */
#endif

#if defined __cplusplus && defined __GNUC__ && ! defined __ICC && 6 <= __GNUC__
# define YY_IGNORE_USELESS_CAST_BEGIN                          \
    _Pragma ("GCC diagnostic push")                            \
    _Pragma ("GCC diagnostic ignored \"-Wuseless-cast\"")
# define YY_IGNORE_USELESS_CAST_END            \
    _Pragma ("GCC diagnostic pop")
#endif
#ifndef YY_IGNORE_USELESS_CAST_BEGIN
# define YY_IGNORE_USELESS_CAST_BEGIN
# define YY_IGNORE_USELESS_CAST_END
#endif


#define YY_ASSERT(E) ((void) (0 && (E)))

#if !defined yyoverflow

/* The parser invokes alloca or malloc; define the necessary symbols.  */

# ifdef YYSTACK_USE_ALLOCA
#  if YYSTACK_USE_ALLOCA
#   ifdef __GNUC__
#    define YYSTACK_ALLOC __builtin_alloca
#   elif defined __BUILTIN_VA_ARG_INCR
#    include <alloca.h> /* INFRINGES ON USER NAME SPACE */
#   elif defined _AIX
#    define YYSTACK_ALLOC __alloca
#   elif defined _MSC_VER
#    include <malloc.h> /* INFRINGES ON USER NAME SPACE */
#    define alloca _alloca
#   else
#    define YYSTACK_ALLOC alloca
#    if ! defined _ALLOCA_H && ! defined EXIT_SUCCESS
#     include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
      /* Use EXIT_SUCCESS as a witness for stdlib.h.  */
#     ifndef EXIT_SUCCESS
#      define EXIT_SUCCESS 0
#     endif
#    endif
#   endif
#  endif
# endif

# ifdef YYSTACK_ALLOC
   /* Pacify GCC's 'empty if-body' warning.  */
#  define YYSTACK_FREE(Ptr) do { /* empty */; } while (0)
#  ifndef YYSTACK_ALLOC_MAXIMUM
    /* The OS might guarantee only one guard page at the bottom of the stack,
       and a page size can be as small as 4096 bytes.  So we cannot safely
       invoke alloca (N) if N exceeds 4096.  Use a slightly smaller number
       to allow for a few compiler-allocated temporary stack slots.  */
#   define YYSTACK_ALLOC_MAXIMUM 4032 /* reasonable circa 2006 */
#  endif
# else
#  define YYSTACK_ALLOC YYMALLOC
#  define YYSTACK_FREE YYFREE
#  ifndef YYSTACK_ALLOC_MAXIMUM
#   define YYSTACK_ALLOC_MAXIMUM YYSIZE_MAXIMUM
#  endif
#  if (defined __cplusplus && ! defined EXIT_SUCCESS \
       && ! ((defined YYMALLOC || defined malloc) \
             && (defined YYFREE || defined free)))
#   include <stdlib.h> /* INFRINGES ON USER NAME SPACE */
#   ifndef EXIT_SUCCESS
#    define EXIT_SUCCESS 0
#   endif
#  endif
#  ifndef YYMALLOC
#   define YYMALLOC malloc
#   if ! defined malloc && ! defined EXIT_SUCCESS
void *malloc (YYSIZE_T); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
#  ifndef YYFREE
#   define YYFREE free
#   if ! defined free && ! defined EXIT_SUCCESS
void free (void *); /* INFRINGES ON USER NAME SPACE */
#   endif
#  endif
# endif
#endif /* !defined yyoverflow */

#if (! defined yyoverflow \
     && (! defined __cplusplus \
         || (defined YYSTYPE_IS_TRIVIAL && YYSTYPE_IS_TRIVIAL)))

/* A type that is properly aligned for any stack member.  */
union yyalloc
{
  yy_state_t yyss_alloc;
  YYSTYPE yyvs_alloc;
};

/* The size of the maximum gap between one aligned stack and the next.  */
# define YYSTACK_GAP_MAXIMUM (YYSIZEOF (union yyalloc) - 1)

/* The size of an array large to enough to hold all stacks, each with
   N elements.  */
# define YYSTACK_BYTES(N) \
     ((N) * (YYSIZEOF (yy_state_t) + YYSIZEOF (YYSTYPE)) \
      + YYSTACK_GAP_MAXIMUM)

# define YYCOPY_NEEDED 1

/* Relocate STACK from its old location to the new one.  The
   local variables YYSIZE and YYSTACKSIZE give the old and new number of
   elements in the stack, and YYPTR gives the new location of the
   stack.  Advance YYPTR to a properly aligned location for the next
   stack.  */
# define YYSTACK_RELOCATE(Stack_alloc, Stack)                           \
    do                                                                  \
      {                                                                 \
        YYPTRDIFF_T yynewbytes;                                         \
        YYCOPY (&yyptr->Stack_alloc, Stack, yysize);                    \
        Stack = &yyptr->Stack_alloc;                                    \
        yynewbytes = yystacksize * YYSIZEOF (*Stack) + YYSTACK_GAP_MAXIMUM; \
        yyptr += yynewbytes / YYSIZEOF (*yyptr);                        \
      }                                                                 \
    while (0)

#endif

#if defined YYCOPY_NEEDED && YYCOPY_NEEDED
/* Copy COUNT objects from SRC to DST.  The source and destination do
   not overlap.  */
# ifndef YYCOPY
#  if defined __GNUC__ && 1 < __GNUC__
#   define YYCOPY(Dst, Src, Count) \
      __builtin_memcpy (Dst, Src, YY_CAST (YYSIZE_T, (Count)) * sizeof (*(Src)))
#  else
#   define YYCOPY(Dst, Src, Count)              \
      do                                        \
        {                                       \
          YYPTRDIFF_T yyi;                      \
          for (yyi = 0; yyi < (Count); yyi++)   \
            (Dst)[yyi] = (Src)[yyi];            \
        }                                       \
      while (0)
#  endif
# endif
#endif /* !YYCOPY_NEEDED */

/* YYFINAL -- State number of the termination state.  */
#define YYFINAL  39
/* YYLAST -- Last index in YYTABLE.  */
#define YYLAST   125

/* YYNTOKENS -- Number of terminals.  */
#define YYNTOKENS  42
/* YYNNTS -- Number of nonterminals.  */
#define YYNNTS  14
/* YYNRULES -- Number of rules.  */
#define YYNRULES  57
/* YYNSTATES -- Number of states.  */
#define YYNSTATES  99

/* YYMAXUTOK -- Last valid token kind.  */
#define YYMAXUTOK   296


/* YYTRANSLATE(TOKEN-NUM) -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex, with out-of-bounds checking.  */
#define YYTRANSLATE(YYX)                                \
  (0 <= (YYX) && (YYX) <= YYMAXUTOK                     \
   ? YY_CAST (yysymbol_kind_t, yytranslate[YYX])        \
   : YYSYMBOL_YYUNDEF)

/* YYTRANSLATE[TOKEN-NUM] -- Symbol number corresponding to TOKEN-NUM
   as returned by yylex.  */
static const yytype_int8 yytranslate[] =
{
       0,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     1,     2,     3,     4,
       5,     6,     7,     8,     9,    10,    11,    12,    13,    14,
      15,    16,    17,    18,    19,    20,    21,    22,    23,    24,
      25,    26,    27,    28,    29,    30,    31,    32,    33,    34,
      35,    36,    37,    38,    39,    40,    41
};

#if YYDEBUG
/* YYRLINE[YYN] -- Source line where rule number YYN was defined.  */
static const yytype_uint8 yyrline[] =
{
       0,    72,    72,    76,    80,    81,    85,    86,    90,    91,
      92,    93,    94,    95,    96,    97,    98,    99,   100,   101,
     105,   109,   110,   114,   122,   131,   132,   139,   140,   141,
     145,   146,   147,   148,   152,   153,   154,   155,   156,   157,
     161,   162,   163,   164,   165,   169,   173,   174,   175,   179,
     185,   199,   200,   201,   202,   203,   204,   205
};
#endif

/** Accessing symbol of state STATE.  */
#define YY_ACCESSING_SYMBOL(State) YY_CAST (yysymbol_kind_t, yystos[State])

#if YYDEBUG || 0
/* The user-facing name of the symbol whose (internal) number is
   YYSYMBOL.  No bounds checking.  */
static const char *yysymbol_name (yysymbol_kind_t yysymbol) YY_ATTRIBUTE_UNUSED;

/* YYTNAME[SYMBOL-NUM] -- String name of the symbol SYMBOL-NUM.
   First, the terminals, then, starting at YYNTOKENS, nonterminals.  */
static const char *const yytname[] =
{
  "\"end of file\"", "error", "\"invalid token\"", "K_LIKE", "K_ILIKE",
  "K_AND", "K_OR", "K_BETWEEN", "K_IN", "K_IS", "K_NULL", "K_NOT",
  "K_TRUE", "K_FALSE", "K_LEN", "K_ANY", "K_ALL", "K_SUM", "RFC3339",
  "DATE", "LPAREN", "RPAREN", "COMMA", "PLUS", "MINUS", "STAR", "SLASH",
  "PERCENT", "LBRACKET", "RBRACKET", "NUMERIC_LITERAL", "STRING_LITERAL",
  "IDENTIFIER", "EQ", "NE", "LT", "LE", "GT", "GE", "REQ", "RNE", "UMINUS",
  "$accept", "input", "expr", "or_expr", "and_expr", "comparison_expr",
  "additive_expr", "multiplicative_expr", "not_expr", "unary_expr",
  "array", "opt_array_elements", "array_elements", "primary", YY_NULLPTR
};

static const char *
yysymbol_name (yysymbol_kind_t yysymbol)
{
  return yytname[yysymbol];
}
#endif

#define YYPACT_NINF (-36)

#define yypact_value_is_default(Yyn) \
  ((Yyn) == YYPACT_NINF)

#define YYTABLE_NINF (-1)

#define yytable_value_is_error(Yyn) \
  0

/* YYPACT[STATE-NUM] -- Index in YYTABLE of the portion describing
   STATE-NUM.  */
static const yytype_int8 yypact[] =
{
      72,    72,   -36,   -36,    72,    72,    72,    72,   -36,   -36,
      72,    93,    93,    72,   -36,   -36,   -36,    12,   -36,    16,
      31,    42,   -22,     8,   -36,   -36,   -36,   -36,   -36,   -36,
     -36,   -36,   -36,    33,   -36,   -36,   -36,    23,    34,   -36,
      72,    72,    72,    72,    72,    72,    15,    24,    72,    72,
      72,    72,    72,    72,    72,    72,    72,    72,    72,    72,
      72,   -36,   -36,    72,    31,    42,   -22,   -22,     6,   -22,
     -36,    50,    72,    72,    72,    72,   -22,   -22,   -22,   -22,
     -22,   -22,   -22,   -22,     8,     8,   -36,   -36,   -36,   -36,
      72,   -36,   -22,   -22,    18,   -22,   -22,    72,   -22
};

/* YYDEFACT[STATE-NUM] -- Default reduction number in state STATE-NUM.
   Performed when YYTABLE does not specify something else to do.  Zero
   means the default is an error.  */
static const yytype_int8 yydefact[] =
{
       0,     0,    56,    57,     0,     0,     0,     0,    54,    55,
       0,     0,     0,    46,    51,    52,    53,     0,     2,     3,
       4,     6,     8,    27,    30,    34,    44,    40,    35,    36,
      37,    38,    39,     0,    42,    41,    49,     0,    47,     1,
       0,     0,     0,     0,     0,     0,     0,     0,     0,     0,
       0,     0,     0,     0,     0,     0,     0,     0,     0,     0,
       0,    43,    45,    48,     5,     7,    17,    18,     0,    25,
      21,     0,     0,     0,     0,     0,     9,    10,    11,    12,
      13,    14,    15,    16,    28,    29,    31,    32,    33,    50,
       0,    22,    19,    20,     0,    26,    23,     0,    24
};

/* YYPGOTO[NTERM-NUM].  */
static const yytype_int8 yypgoto[] =
{
     -36,   -36,    11,   -36,    21,    22,   -35,   -13,    -1,    36,
     -36,   -36,   -36,   -36
};

/* YYDEFGOTO[NTERM-NUM].  */
static const yytype_int8 yydefgoto[] =
{
       0,    17,    18,    19,    20,    21,    22,    23,    24,    25,
      26,    37,    38,    27
};

/* YYTABLE[YYPACT[STATE-NUM]] -- What to do in state STATE-NUM.  If
   positive, shift that token.  If negative, reduce the rule whose
   number is the opposite.  If YYTABLE_NINF, syntax error.  */
static const yytype_int8 yytable[] =
{
      28,    56,    57,    29,    30,    31,    32,    66,    67,    68,
      69,    90,    39,    76,    77,    78,    79,    80,    81,    82,
      83,    33,    40,    97,    36,    70,    71,    72,    73,    56,
      57,    74,    75,    58,    59,    60,    41,    92,    93,    94,
      95,    56,    57,    84,    85,    42,    43,    34,    35,    44,
      45,    46,    62,    47,    61,    96,    63,    86,    87,    88,
      91,    64,    98,    65,     0,     0,     0,     0,     0,     0,
       0,     0,     0,     0,    89,    48,    49,    50,    51,    52,
      53,    54,    55,     1,     2,     3,     4,     5,     6,     7,
       8,     9,    10,     0,     0,    11,    12,     0,     0,     0,
      13,     0,    14,    15,    16,     2,     3,     0,     0,     0,
       0,     8,     9,    10,     0,     0,    11,    12,     0,     0,
       0,    13,     0,    14,    15,    16
};

static const yytype_int8 yycheck[] =
{
       1,    23,    24,     4,     5,     6,     7,    42,    43,    44,
      45,     5,     0,    48,    49,    50,    51,    52,    53,    54,
      55,    10,     6,     5,    13,    10,    11,     3,     4,    23,
      24,     7,     8,    25,    26,    27,     5,    72,    73,    74,
      75,    23,    24,    56,    57,     3,     4,    11,    12,     7,
       8,     9,    29,    11,    21,    90,    22,    58,    59,    60,
      10,    40,    97,    41,    -1,    -1,    -1,    -1,    -1,    -1,
      -1,    -1,    -1,    -1,    63,    33,    34,    35,    36,    37,
      38,    39,    40,    11,    12,    13,    14,    15,    16,    17,
      18,    19,    20,    -1,    -1,    23,    24,    -1,    -1,    -1,
      28,    -1,    30,    31,    32,    12,    13,    -1,    -1,    -1,
      -1,    18,    19,    20,    -1,    -1,    23,    24,    -1,    -1,
      -1,    28,    -1,    30,    31,    32
};

/* YYSTOS[STATE-NUM] -- The symbol kind of the accessing symbol of
   state STATE-NUM.  */
static const yytype_int8 yystos[] =
{
       0,    11,    12,    13,    14,    15,    16,    17,    18,    19,
      20,    23,    24,    28,    30,    31,    32,    43,    44,    45,
      46,    47,    48,    49,    50,    51,    52,    55,    50,    50,
      50,    50,    50,    44,    51,    51,    44,    53,    54,     0,
       6,     5,     3,     4,     7,     8,     9,    11,    33,    34,
      35,    36,    37,    38,    39,    40,    23,    24,    25,    26,
      27,    21,    29,    22,    46,    47,    48,    48,    48,    48,
      10,    11,     3,     4,     7,     8,    48,    48,    48,    48,
      48,    48,    48,    48,    49,    49,    50,    50,    50,    44,
       5,    10,    48,    48,    48,    48,    48,     5,    48
};

/* YYR1[RULE-NUM] -- Symbol kind of the left-hand side of rule RULE-NUM.  */
static const yytype_int8 yyr1[] =
{
       0,    42,    43,    44,    45,    45,    46,    46,    47,    47,
      47,    47,    47,    47,    47,    47,    47,    47,    47,    47,
      47,    47,    47,    47,    47,    47,    47,    48,    48,    48,
      49,    49,    49,    49,    50,    50,    50,    50,    50,    50,
      51,    51,    51,    51,    51,    52,    53,    53,    53,    54,
      54,    55,    55,    55,    55,    55,    55,    55
};

/* YYR2[RULE-NUM] -- Number of symbols on the right-hand side of rule RULE-NUM.  */
static const yytype_int8 yyr2[] =
{
       0,     2,     1,     1,     1,     3,     1,     3,     1,     3,
       3,     3,     3,     3,     3,     3,     3,     3,     3,     4,
       4,     3,     4,     5,     6,     3,     4,     1,     3,     3,
       1,     3,     3,     3,     1,     2,     2,     2,     2,     2,
       1,     2,     2,     3,     1,     3,     0,     1,     2,     1,
       3,     1,     1,     1,     1,     1,     1,     1
};


enum { YYENOMEM = -2 };

#define yyerrok         (yyerrstatus = 0)
#define yyclearin       (yychar = YYEMPTY)

#define YYACCEPT        goto yyacceptlab
#define YYABORT         goto yyabortlab
#define YYERROR         goto yyerrorlab
#define YYNOMEM         goto yyexhaustedlab


#define YYRECOVERING()  (!!yyerrstatus)

#define YYBACKUP(Token, Value)                                    \
  do                                                              \
    if (yychar == YYEMPTY)                                        \
      {                                                           \
        yychar = (Token);                                         \
        yylval = (Value);                                         \
        YYPOPSTACK (yylen);                                       \
        yystate = *yyssp;                                         \
        goto yybackup;                                            \
      }                                                           \
    else                                                          \
      {                                                           \
        yyerror (YY_("syntax error: cannot back up")); \
        YYERROR;                                                  \
      }                                                           \
  while (0)

/* Backward compatibility with an undocumented macro.
   Use YYerror or YYUNDEF. */
#define YYERRCODE YYUNDEF


/* Enable debugging if requested.  */
#if YYDEBUG

# ifndef YYFPRINTF
#  include <stdio.h> /* INFRINGES ON USER NAME SPACE */
#  define YYFPRINTF fprintf
# endif

# define YYDPRINTF(Args)                        \
do {                                            \
  if (yydebug)                                  \
    YYFPRINTF Args;                             \
} while (0)




# define YY_SYMBOL_PRINT(Title, Kind, Value, Location)                    \
do {                                                                      \
  if (yydebug)                                                            \
    {                                                                     \
      YYFPRINTF (stderr, "%s ", Title);                                   \
      yy_symbol_print (stderr,                                            \
                  Kind, Value); \
      YYFPRINTF (stderr, "\n");                                           \
    }                                                                     \
} while (0)


/*-----------------------------------.
| Print this symbol's value on YYO.  |
`-----------------------------------*/

static void
yy_symbol_value_print (FILE *yyo,
                       yysymbol_kind_t yykind, YYSTYPE const * const yyvaluep)
{
  FILE *yyoutput = yyo;
  YY_USE (yyoutput);
  if (!yyvaluep)
    return;
  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  YY_USE (yykind);
  YY_IGNORE_MAYBE_UNINITIALIZED_END
}


/*---------------------------.
| Print this symbol on YYO.  |
`---------------------------*/

static void
yy_symbol_print (FILE *yyo,
                 yysymbol_kind_t yykind, YYSTYPE const * const yyvaluep)
{
  YYFPRINTF (yyo, "%s %s (",
             yykind < YYNTOKENS ? "token" : "nterm", yysymbol_name (yykind));

  yy_symbol_value_print (yyo, yykind, yyvaluep);
  YYFPRINTF (yyo, ")");
}

/*------------------------------------------------------------------.
| yy_stack_print -- Print the state stack from its BOTTOM up to its |
| TOP (included).                                                   |
`------------------------------------------------------------------*/

static void
yy_stack_print (yy_state_t *yybottom, yy_state_t *yytop)
{
  YYFPRINTF (stderr, "Stack now");
  for (; yybottom <= yytop; yybottom++)
    {
      int yybot = *yybottom;
      YYFPRINTF (stderr, " %d", yybot);
    }
  YYFPRINTF (stderr, "\n");
}

# define YY_STACK_PRINT(Bottom, Top)                            \
do {                                                            \
  if (yydebug)                                                  \
    yy_stack_print ((Bottom), (Top));                           \
} while (0)


/*------------------------------------------------.
| Report that the YYRULE is going to be reduced.  |
`------------------------------------------------*/

static void
yy_reduce_print (yy_state_t *yyssp, YYSTYPE *yyvsp,
                 int yyrule)
{
  int yylno = yyrline[yyrule];
  int yynrhs = yyr2[yyrule];
  int yyi;
  YYFPRINTF (stderr, "Reducing stack by rule %d (line %d):\n",
             yyrule - 1, yylno);
  /* The symbols being reduced.  */
  for (yyi = 0; yyi < yynrhs; yyi++)
    {
      YYFPRINTF (stderr, "   $%d = ", yyi + 1);
      yy_symbol_print (stderr,
                       YY_ACCESSING_SYMBOL (+yyssp[yyi + 1 - yynrhs]),
                       &yyvsp[(yyi + 1) - (yynrhs)]);
      YYFPRINTF (stderr, "\n");
    }
}

# define YY_REDUCE_PRINT(Rule)          \
do {                                    \
  if (yydebug)                          \
    yy_reduce_print (yyssp, yyvsp, Rule); \
} while (0)

/* Nonzero means print parse trace.  It is left uninitialized so that
   multiple parsers can coexist.  */
int yydebug;
#else /* !YYDEBUG */
# define YYDPRINTF(Args) ((void) 0)
# define YY_SYMBOL_PRINT(Title, Kind, Value, Location)
# define YY_STACK_PRINT(Bottom, Top)
# define YY_REDUCE_PRINT(Rule)
#endif /* !YYDEBUG */


/* YYINITDEPTH -- initial size of the parser's stacks.  */
#ifndef YYINITDEPTH
# define YYINITDEPTH 200
#endif

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

#ifndef YYMAXDEPTH
# define YYMAXDEPTH 10000
#endif






/*-----------------------------------------------.
| Release the memory associated to this symbol.  |
`-----------------------------------------------*/

static void
yydestruct (const char *yymsg,
            yysymbol_kind_t yykind, YYSTYPE *yyvaluep)
{
  YY_USE (yyvaluep);
  if (!yymsg)
    yymsg = "Deleting";
  YY_SYMBOL_PRINT (yymsg, yykind, yyvaluep, yylocationp);

  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  YY_USE (yykind);
  YY_IGNORE_MAYBE_UNINITIALIZED_END
}


/* Lookahead token kind.  */
int yychar;

/* The semantic value of the lookahead symbol.  */
YYSTYPE yylval;
/* Number of syntax errors so far.  */
int yynerrs;




/*----------.
| yyparse.  |
`----------*/

int
yyparse (void)
{
    yy_state_fast_t yystate = 0;
    /* Number of tokens to shift before error messages enabled.  */
    int yyerrstatus = 0;

    /* Refer to the stacks through separate pointers, to allow yyoverflow
       to reallocate them elsewhere.  */

    /* Their size.  */
    YYPTRDIFF_T yystacksize = YYINITDEPTH;

    /* The state stack: array, bottom, top.  */
    yy_state_t yyssa[YYINITDEPTH];
    yy_state_t *yyss = yyssa;
    yy_state_t *yyssp = yyss;

    /* The semantic value stack: array, bottom, top.  */
    YYSTYPE yyvsa[YYINITDEPTH];
    YYSTYPE *yyvs = yyvsa;
    YYSTYPE *yyvsp = yyvs;

  int yyn;
  /* The return value of yyparse.  */
  int yyresult;
  /* Lookahead symbol kind.  */
  yysymbol_kind_t yytoken = YYSYMBOL_YYEMPTY;
  /* The variables used to return semantic value and location from the
     action routines.  */
  YYSTYPE yyval;



#define YYPOPSTACK(N)   (yyvsp -= (N), yyssp -= (N))

  /* The number of symbols on the RHS of the reduced rule.
     Keep to zero when no symbol should be popped.  */
  int yylen = 0;

  YYDPRINTF ((stderr, "Starting parse\n"));

  yychar = YYEMPTY; /* Cause a token to be read.  */

  goto yysetstate;


/*------------------------------------------------------------.
| yynewstate -- push a new state, which is found in yystate.  |
`------------------------------------------------------------*/
yynewstate:
  /* In all cases, when you get here, the value and location stacks
     have just been pushed.  So pushing a state here evens the stacks.  */
  yyssp++;


/*--------------------------------------------------------------------.
| yysetstate -- set current state (the top of the stack) to yystate.  |
`--------------------------------------------------------------------*/
yysetstate:
  YYDPRINTF ((stderr, "Entering state %d\n", yystate));
  YY_ASSERT (0 <= yystate && yystate < YYNSTATES);
  YY_IGNORE_USELESS_CAST_BEGIN
  *yyssp = YY_CAST (yy_state_t, yystate);
  YY_IGNORE_USELESS_CAST_END
  YY_STACK_PRINT (yyss, yyssp);

  if (yyss + yystacksize - 1 <= yyssp)
#if !defined yyoverflow && !defined YYSTACK_RELOCATE
    YYNOMEM;
#else
    {
      /* Get the current used size of the three stacks, in elements.  */
      YYPTRDIFF_T yysize = yyssp - yyss + 1;

# if defined yyoverflow
      {
        /* Give user a chance to reallocate the stack.  Use copies of
           these so that the &'s don't force the real ones into
           memory.  */
        yy_state_t *yyss1 = yyss;
        YYSTYPE *yyvs1 = yyvs;

        /* Each stack pointer address is followed by the size of the
           data in use in that stack, in bytes.  This used to be a
           conditional around just the two extra args, but that might
           be undefined if yyoverflow is a macro.  */
        yyoverflow (YY_("memory exhausted"),
                    &yyss1, yysize * YYSIZEOF (*yyssp),
                    &yyvs1, yysize * YYSIZEOF (*yyvsp),
                    &yystacksize);
        yyss = yyss1;
        yyvs = yyvs1;
      }
# else /* defined YYSTACK_RELOCATE */
      /* Extend the stack our own way.  */
      if (YYMAXDEPTH <= yystacksize)
        YYNOMEM;
      yystacksize *= 2;
      if (YYMAXDEPTH < yystacksize)
        yystacksize = YYMAXDEPTH;

      {
        yy_state_t *yyss1 = yyss;
        union yyalloc *yyptr =
          YY_CAST (union yyalloc *,
                   YYSTACK_ALLOC (YY_CAST (YYSIZE_T, YYSTACK_BYTES (yystacksize))));
        if (! yyptr)
          YYNOMEM;
        YYSTACK_RELOCATE (yyss_alloc, yyss);
        YYSTACK_RELOCATE (yyvs_alloc, yyvs);
#  undef YYSTACK_RELOCATE
        if (yyss1 != yyssa)
          YYSTACK_FREE (yyss1);
      }
# endif

      yyssp = yyss + yysize - 1;
      yyvsp = yyvs + yysize - 1;

      YY_IGNORE_USELESS_CAST_BEGIN
      YYDPRINTF ((stderr, "Stack size increased to %ld\n",
                  YY_CAST (long, yystacksize)));
      YY_IGNORE_USELESS_CAST_END

      if (yyss + yystacksize - 1 <= yyssp)
        YYABORT;
    }
#endif /* !defined yyoverflow && !defined YYSTACK_RELOCATE */


  if (yystate == YYFINAL)
    YYACCEPT;

  goto yybackup;


/*-----------.
| yybackup.  |
`-----------*/
yybackup:
  /* Do appropriate processing given the current state.  Read a
     lookahead token if we need one and don't already have one.  */

  /* First try to decide what to do without reference to lookahead token.  */
  yyn = yypact[yystate];
  if (yypact_value_is_default (yyn))
    goto yydefault;

  /* Not known => get a lookahead token if don't already have one.  */

  /* YYCHAR is either empty, or end-of-input, or a valid lookahead.  */
  if (yychar == YYEMPTY)
    {
      YYDPRINTF ((stderr, "Reading a token\n"));
      yychar = yylex ();
    }

  if (yychar <= YYEOF)
    {
      yychar = YYEOF;
      yytoken = YYSYMBOL_YYEOF;
      YYDPRINTF ((stderr, "Now at end of input.\n"));
    }
  else if (yychar == YYerror)
    {
      /* The scanner already issued an error message, process directly
         to error recovery.  But do not keep the error token as
         lookahead, it is too special and may lead us to an endless
         loop in error recovery. */
      yychar = YYUNDEF;
      yytoken = YYSYMBOL_YYerror;
      goto yyerrlab1;
    }
  else
    {
      yytoken = YYTRANSLATE (yychar);
      YY_SYMBOL_PRINT ("Next token is", yytoken, &yylval, &yylloc);
    }

  /* If the proper action on seeing token YYTOKEN is to reduce or to
     detect an error, take that action.  */
  yyn += yytoken;
  if (yyn < 0 || YYLAST < yyn || yycheck[yyn] != yytoken)
    goto yydefault;
  yyn = yytable[yyn];
  if (yyn <= 0)
    {
      if (yytable_value_is_error (yyn))
        goto yyerrlab;
      yyn = -yyn;
      goto yyreduce;
    }

  /* Count tokens shifted since error; after three, turn off error
     status.  */
  if (yyerrstatus)
    yyerrstatus--;

  /* Shift the lookahead token.  */
  YY_SYMBOL_PRINT ("Shifting", yytoken, &yylval, &yylloc);
  yystate = yyn;
  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  *++yyvsp = yylval;
  YY_IGNORE_MAYBE_UNINITIALIZED_END

  /* Discard the shifted token.  */
  yychar = YYEMPTY;
  goto yynewstate;


/*-----------------------------------------------------------.
| yydefault -- do the default action for the current state.  |
`-----------------------------------------------------------*/
yydefault:
  yyn = yydefact[yystate];
  if (yyn == 0)
    goto yyerrlab;
  goto yyreduce;


/*-----------------------------.
| yyreduce -- do a reduction.  |
`-----------------------------*/
yyreduce:
  /* yyn is the number of a rule to reduce with.  */
  yylen = yyr2[yyn];

  /* If YYLEN is nonzero, implement the default value of the action:
     '$$ = $1'.

     Otherwise, the following line sets YYVAL to garbage.
     This behavior is undocumented and Bison
     users should not rely upon it.  Assigning to YYVAL
     unconditionally makes the parser a bit smaller, and it avoids a
     GCC warning that YYVAL may be used uninitialized.  */
  yyval = yyvsp[1-yylen];


  YY_REDUCE_PRINT (yyn);
  switch (yyn)
    {
  case 2: /* input: expr  */
#line 72 "tsl_parser.y"
         { (yyval.node) = (yyvsp[0].node); ast_root = (yyvsp[0].node); }
#line 1212 "../tsl/tsl_parser.tab.c"
    break;

  case 5: /* or_expr: or_expr K_OR and_expr  */
#line 81 "tsl_parser.y"
                                   { (yyval.node) = ast_create_binary(K_OR, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1218 "../tsl/tsl_parser.tab.c"
    break;

  case 7: /* and_expr: and_expr K_AND comparison_expr  */
#line 86 "tsl_parser.y"
                                          { (yyval.node) = ast_create_binary(K_AND, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1224 "../tsl/tsl_parser.tab.c"
    break;

  case 9: /* comparison_expr: comparison_expr EQ additive_expr  */
#line 91 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(EQ, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1230 "../tsl/tsl_parser.tab.c"
    break;

  case 10: /* comparison_expr: comparison_expr NE additive_expr  */
#line 92 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(NE, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1236 "../tsl/tsl_parser.tab.c"
    break;

  case 11: /* comparison_expr: comparison_expr LT additive_expr  */
#line 93 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(LT, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1242 "../tsl/tsl_parser.tab.c"
    break;

  case 12: /* comparison_expr: comparison_expr LE additive_expr  */
#line 94 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(LE, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1248 "../tsl/tsl_parser.tab.c"
    break;

  case 13: /* comparison_expr: comparison_expr GT additive_expr  */
#line 95 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(GT, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1254 "../tsl/tsl_parser.tab.c"
    break;

  case 14: /* comparison_expr: comparison_expr GE additive_expr  */
#line 96 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(GE, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1260 "../tsl/tsl_parser.tab.c"
    break;

  case 15: /* comparison_expr: comparison_expr REQ additive_expr  */
#line 97 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(REQ, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1266 "../tsl/tsl_parser.tab.c"
    break;

  case 16: /* comparison_expr: comparison_expr RNE additive_expr  */
#line 98 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(RNE, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1272 "../tsl/tsl_parser.tab.c"
    break;

  case 17: /* comparison_expr: comparison_expr K_LIKE additive_expr  */
#line 99 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(K_LIKE, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1278 "../tsl/tsl_parser.tab.c"
    break;

  case 18: /* comparison_expr: comparison_expr K_ILIKE additive_expr  */
#line 100 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(K_ILIKE, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1284 "../tsl/tsl_parser.tab.c"
    break;

  case 19: /* comparison_expr: comparison_expr K_NOT K_LIKE additive_expr  */
#line 101 "tsl_parser.y"
                                                  { 
        ast_node *like_expr = ast_create_binary(K_LIKE, (yyvsp[-3].node), (yyvsp[0].node));
        (yyval.node) = ast_create_unary(K_NOT, like_expr);
    }
#line 1293 "../tsl/tsl_parser.tab.c"
    break;

  case 20: /* comparison_expr: comparison_expr K_NOT K_ILIKE additive_expr  */
#line 105 "tsl_parser.y"
                                                  { 
        ast_node *ilike_expr = ast_create_binary(K_ILIKE, (yyvsp[-3].node), (yyvsp[0].node));
        (yyval.node) = ast_create_unary(K_NOT, ilike_expr);
    }
#line 1302 "../tsl/tsl_parser.tab.c"
    break;

  case 21: /* comparison_expr: comparison_expr K_IS K_NULL  */
#line 109 "tsl_parser.y"
                                            { (yyval.node) = ast_create_binary(K_IS, (yyvsp[-2].node), ast_create_null()); }
#line 1308 "../tsl/tsl_parser.tab.c"
    break;

  case 22: /* comparison_expr: comparison_expr K_IS K_NOT K_NULL  */
#line 110 "tsl_parser.y"
                                            { 
        ast_node *is_null = ast_create_binary(K_IS, (yyvsp[-3].node), ast_create_null());
        (yyval.node) = ast_create_unary(K_NOT, is_null);
    }
#line 1317 "../tsl/tsl_parser.tab.c"
    break;

  case 23: /* comparison_expr: comparison_expr K_BETWEEN additive_expr K_AND additive_expr  */
#line 114 "tsl_parser.y"
                                                                  {
        ast_node **elements = malloc(2 * sizeof(ast_node*));
        elements[0] = (yyvsp[-2].node);
        elements[1] = (yyvsp[0].node);
        ast_node *range = ast_create_array(2, elements);
        (yyval.node) = ast_create_binary(K_BETWEEN, (yyvsp[-4].node), range);
        free(elements);
    }
#line 1330 "../tsl/tsl_parser.tab.c"
    break;

  case 24: /* comparison_expr: comparison_expr K_NOT K_BETWEEN additive_expr K_AND additive_expr  */
#line 122 "tsl_parser.y"
                                                                        {
        ast_node **elements = malloc(2 * sizeof(ast_node*));
        elements[0] = (yyvsp[-2].node);
        elements[1] = (yyvsp[0].node);
        ast_node *range = ast_create_array(2, elements);
        ast_node *between = ast_create_binary(K_BETWEEN, (yyvsp[-5].node), range);
        (yyval.node) = ast_create_unary(K_NOT, between);
        free(elements);
    }
#line 1344 "../tsl/tsl_parser.tab.c"
    break;

  case 25: /* comparison_expr: comparison_expr K_IN additive_expr  */
#line 131 "tsl_parser.y"
                                             { (yyval.node) = ast_create_binary(K_IN, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1350 "../tsl/tsl_parser.tab.c"
    break;

  case 26: /* comparison_expr: comparison_expr K_NOT K_IN additive_expr  */
#line 132 "tsl_parser.y"
                                               {
        ast_node *in_expr = ast_create_binary(K_IN, (yyvsp[-3].node), (yyvsp[0].node));
        (yyval.node) = ast_create_unary(K_NOT, in_expr);
    }
#line 1359 "../tsl/tsl_parser.tab.c"
    break;

  case 28: /* additive_expr: additive_expr PLUS multiplicative_expr  */
#line 140 "tsl_parser.y"
                                               { (yyval.node) = ast_create_binary(PLUS, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1365 "../tsl/tsl_parser.tab.c"
    break;

  case 29: /* additive_expr: additive_expr MINUS multiplicative_expr  */
#line 141 "tsl_parser.y"
                                               { (yyval.node) = ast_create_binary(MINUS, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1371 "../tsl/tsl_parser.tab.c"
    break;

  case 31: /* multiplicative_expr: multiplicative_expr STAR not_expr  */
#line 146 "tsl_parser.y"
                                           { (yyval.node) = ast_create_binary(STAR, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1377 "../tsl/tsl_parser.tab.c"
    break;

  case 32: /* multiplicative_expr: multiplicative_expr SLASH not_expr  */
#line 147 "tsl_parser.y"
                                           { (yyval.node) = ast_create_binary(SLASH, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1383 "../tsl/tsl_parser.tab.c"
    break;

  case 33: /* multiplicative_expr: multiplicative_expr PERCENT not_expr  */
#line 148 "tsl_parser.y"
                                           { (yyval.node) = ast_create_binary(PERCENT, (yyvsp[-2].node), (yyvsp[0].node)); }
#line 1389 "../tsl/tsl_parser.tab.c"
    break;

  case 35: /* not_expr: K_NOT not_expr  */
#line 153 "tsl_parser.y"
                                   { (yyval.node) = ast_create_unary(K_NOT, (yyvsp[0].node)); }
#line 1395 "../tsl/tsl_parser.tab.c"
    break;

  case 36: /* not_expr: K_LEN not_expr  */
#line 154 "tsl_parser.y"
                                   { (yyval.node) = ast_create_unary(K_LEN, (yyvsp[0].node)); }
#line 1401 "../tsl/tsl_parser.tab.c"
    break;

  case 37: /* not_expr: K_ANY not_expr  */
#line 155 "tsl_parser.y"
                                   { (yyval.node) = ast_create_unary(K_ANY, (yyvsp[0].node)); }
#line 1407 "../tsl/tsl_parser.tab.c"
    break;

  case 38: /* not_expr: K_ALL not_expr  */
#line 156 "tsl_parser.y"
                                   { (yyval.node) = ast_create_unary(K_ALL, (yyvsp[0].node)); }
#line 1413 "../tsl/tsl_parser.tab.c"
    break;

  case 39: /* not_expr: K_SUM not_expr  */
#line 157 "tsl_parser.y"
                                   { (yyval.node) = ast_create_unary(K_SUM, (yyvsp[0].node)); }
#line 1419 "../tsl/tsl_parser.tab.c"
    break;

  case 41: /* unary_expr: MINUS unary_expr  */
#line 162 "tsl_parser.y"
                                   { (yyval.node) = ast_create_unary(UMINUS, (yyvsp[0].node)); }
#line 1425 "../tsl/tsl_parser.tab.c"
    break;

  case 42: /* unary_expr: PLUS unary_expr  */
#line 163 "tsl_parser.y"
                                   { (yyval.node) = (yyvsp[0].node); }
#line 1431 "../tsl/tsl_parser.tab.c"
    break;

  case 43: /* unary_expr: LPAREN expr RPAREN  */
#line 164 "tsl_parser.y"
                                   { (yyval.node) = (yyvsp[-1].node); }
#line 1437 "../tsl/tsl_parser.tab.c"
    break;

  case 44: /* unary_expr: array  */
#line 165 "tsl_parser.y"
                                   { (yyval.node) = (yyvsp[0].node); }
#line 1443 "../tsl/tsl_parser.tab.c"
    break;

  case 45: /* array: LBRACKET opt_array_elements RBRACKET  */
#line 169 "tsl_parser.y"
                                           { (yyval.node) = (yyvsp[-1].node); }
#line 1449 "../tsl/tsl_parser.tab.c"
    break;

  case 46: /* opt_array_elements: %empty  */
#line 173 "tsl_parser.y"
                                  { (yyval.node) = ast_create_array(0, NULL); }
#line 1455 "../tsl/tsl_parser.tab.c"
    break;

  case 47: /* opt_array_elements: array_elements  */
#line 174 "tsl_parser.y"
                                  { (yyval.node) = (yyvsp[0].node); }
#line 1461 "../tsl/tsl_parser.tab.c"
    break;

  case 48: /* opt_array_elements: array_elements COMMA  */
#line 175 "tsl_parser.y"
                                  { (yyval.node) = (yyvsp[-1].node); }
#line 1467 "../tsl/tsl_parser.tab.c"
    break;

  case 49: /* array_elements: expr  */
#line 179 "tsl_parser.y"
                                    { 
                                      ast_node **elements = malloc(sizeof(ast_node*));
                                      elements[0] = (yyvsp[0].node);
                                      (yyval.node) = ast_create_array(1, elements);
                                      free(elements);
                                    }
#line 1478 "../tsl/tsl_parser.tab.c"
    break;

  case 50: /* array_elements: array_elements COMMA expr  */
#line 185 "tsl_parser.y"
                                    {
                                      ast_node **elements = malloc(((yyvsp[-2].node)->data.array.size + 1) * sizeof(ast_node*));
                                      // Copy existing elements
                                      for(int i = 0; i < (yyvsp[-2].node)->data.array.size; i++) {
                                          elements[i] = ast_clone((yyvsp[-2].node)->data.array.elements[i]);
                                      }
                                      elements[(yyvsp[-2].node)->data.array.size] = ast_clone((yyvsp[0].node));
                                      (yyval.node) = ast_create_array((yyvsp[-2].node)->data.array.size + 1, elements);
                                      free(elements);
                                      ast_free((yyvsp[-2].node));
                                    }
#line 1494 "../tsl/tsl_parser.tab.c"
    break;

  case 51: /* primary: NUMERIC_LITERAL  */
#line 199 "tsl_parser.y"
                            { (yyval.node) = ast_create_number((yyvsp[0].str)); free((yyvsp[0].str)); }
#line 1500 "../tsl/tsl_parser.tab.c"
    break;

  case 52: /* primary: STRING_LITERAL  */
#line 200 "tsl_parser.y"
                            { (yyval.node) = ast_create_string((yyvsp[0].str)); free((yyvsp[0].str)); }
#line 1506 "../tsl/tsl_parser.tab.c"
    break;

  case 53: /* primary: IDENTIFIER  */
#line 201 "tsl_parser.y"
                            { (yyval.node) = ast_create_identifier((yyvsp[0].str)); free((yyvsp[0].str)); }
#line 1512 "../tsl/tsl_parser.tab.c"
    break;

  case 54: /* primary: RFC3339  */
#line 202 "tsl_parser.y"
                            { (yyval.node) = ast_create_rfc3339((yyvsp[0].str)); free((yyvsp[0].str)); }
#line 1518 "../tsl/tsl_parser.tab.c"
    break;

  case 55: /* primary: DATE  */
#line 203 "tsl_parser.y"
                            { (yyval.node) = ast_create_date((yyvsp[0].str)); free((yyvsp[0].str)); }
#line 1524 "../tsl/tsl_parser.tab.c"
    break;

  case 56: /* primary: K_TRUE  */
#line 204 "tsl_parser.y"
                            { (yyval.node) = ast_create_boolean(1); }
#line 1530 "../tsl/tsl_parser.tab.c"
    break;

  case 57: /* primary: K_FALSE  */
#line 205 "tsl_parser.y"
                            { (yyval.node) = ast_create_boolean(0); }
#line 1536 "../tsl/tsl_parser.tab.c"
    break;


#line 1540 "../tsl/tsl_parser.tab.c"

      default: break;
    }
  /* User semantic actions sometimes alter yychar, and that requires
     that yytoken be updated with the new translation.  We take the
     approach of translating immediately before every use of yytoken.
     One alternative is translating here after every semantic action,
     but that translation would be missed if the semantic action invokes
     YYABORT, YYACCEPT, or YYERROR immediately after altering yychar or
     if it invokes YYBACKUP.  In the case of YYABORT or YYACCEPT, an
     incorrect destructor might then be invoked immediately.  In the
     case of YYERROR or YYBACKUP, subsequent parser actions might lead
     to an incorrect destructor call or verbose syntax error message
     before the lookahead is translated.  */
  YY_SYMBOL_PRINT ("-> $$ =", YY_CAST (yysymbol_kind_t, yyr1[yyn]), &yyval, &yyloc);

  YYPOPSTACK (yylen);
  yylen = 0;

  *++yyvsp = yyval;

  /* Now 'shift' the result of the reduction.  Determine what state
     that goes to, based on the state we popped back to and the rule
     number reduced by.  */
  {
    const int yylhs = yyr1[yyn] - YYNTOKENS;
    const int yyi = yypgoto[yylhs] + *yyssp;
    yystate = (0 <= yyi && yyi <= YYLAST && yycheck[yyi] == *yyssp
               ? yytable[yyi]
               : yydefgoto[yylhs]);
  }

  goto yynewstate;


/*--------------------------------------.
| yyerrlab -- here on detecting error.  |
`--------------------------------------*/
yyerrlab:
  /* Make sure we have latest lookahead translation.  See comments at
     user semantic actions for why this is necessary.  */
  yytoken = yychar == YYEMPTY ? YYSYMBOL_YYEMPTY : YYTRANSLATE (yychar);
  /* If not already recovering from an error, report this error.  */
  if (!yyerrstatus)
    {
      ++yynerrs;
      yyerror (YY_("syntax error"));
    }

  if (yyerrstatus == 3)
    {
      /* If just tried and failed to reuse lookahead token after an
         error, discard it.  */

      if (yychar <= YYEOF)
        {
          /* Return failure if at end of input.  */
          if (yychar == YYEOF)
            YYABORT;
        }
      else
        {
          yydestruct ("Error: discarding",
                      yytoken, &yylval);
          yychar = YYEMPTY;
        }
    }

  /* Else will try to reuse lookahead token after shifting the error
     token.  */
  goto yyerrlab1;


/*---------------------------------------------------.
| yyerrorlab -- error raised explicitly by YYERROR.  |
`---------------------------------------------------*/
yyerrorlab:
  /* Pacify compilers when the user code never invokes YYERROR and the
     label yyerrorlab therefore never appears in user code.  */
  if (0)
    YYERROR;
  ++yynerrs;

  /* Do not reclaim the symbols of the rule whose action triggered
     this YYERROR.  */
  YYPOPSTACK (yylen);
  yylen = 0;
  YY_STACK_PRINT (yyss, yyssp);
  yystate = *yyssp;
  goto yyerrlab1;


/*-------------------------------------------------------------.
| yyerrlab1 -- common code for both syntax error and YYERROR.  |
`-------------------------------------------------------------*/
yyerrlab1:
  yyerrstatus = 3;      /* Each real token shifted decrements this.  */

  /* Pop stack until we find a state that shifts the error token.  */
  for (;;)
    {
      yyn = yypact[yystate];
      if (!yypact_value_is_default (yyn))
        {
          yyn += YYSYMBOL_YYerror;
          if (0 <= yyn && yyn <= YYLAST && yycheck[yyn] == YYSYMBOL_YYerror)
            {
              yyn = yytable[yyn];
              if (0 < yyn)
                break;
            }
        }

      /* Pop the current state because it cannot handle the error token.  */
      if (yyssp == yyss)
        YYABORT;


      yydestruct ("Error: popping",
                  YY_ACCESSING_SYMBOL (yystate), yyvsp);
      YYPOPSTACK (1);
      yystate = *yyssp;
      YY_STACK_PRINT (yyss, yyssp);
    }

  YY_IGNORE_MAYBE_UNINITIALIZED_BEGIN
  *++yyvsp = yylval;
  YY_IGNORE_MAYBE_UNINITIALIZED_END


  /* Shift the error token.  */
  YY_SYMBOL_PRINT ("Shifting", YY_ACCESSING_SYMBOL (yyn), yyvsp, yylsp);

  yystate = yyn;
  goto yynewstate;


/*-------------------------------------.
| yyacceptlab -- YYACCEPT comes here.  |
`-------------------------------------*/
yyacceptlab:
  yyresult = 0;
  goto yyreturnlab;


/*-----------------------------------.
| yyabortlab -- YYABORT comes here.  |
`-----------------------------------*/
yyabortlab:
  yyresult = 1;
  goto yyreturnlab;


/*-----------------------------------------------------------.
| yyexhaustedlab -- YYNOMEM (memory exhaustion) comes here.  |
`-----------------------------------------------------------*/
yyexhaustedlab:
  yyerror (YY_("memory exhausted"));
  yyresult = 2;
  goto yyreturnlab;


/*----------------------------------------------------------.
| yyreturnlab -- parsing is finished, clean up and return.  |
`----------------------------------------------------------*/
yyreturnlab:
  if (yychar != YYEMPTY)
    {
      /* Make sure we have latest lookahead translation.  See comments at
         user semantic actions for why this is necessary.  */
      yytoken = YYTRANSLATE (yychar);
      yydestruct ("Cleanup: discarding lookahead",
                  yytoken, &yylval);
    }
  /* Do not reclaim the symbols of the rule whose action triggered
     this YYABORT or YYACCEPT.  */
  YYPOPSTACK (yylen);
  YY_STACK_PRINT (yyss, yyssp);
  while (yyssp != yyss)
    {
      yydestruct ("Cleanup: popping",
                  YY_ACCESSING_SYMBOL (+*yyssp), yyvsp);
      YYPOPSTACK (1);
    }
#ifndef yyoverflow
  if (yyss != yyssa)
    YYSTACK_FREE (yyss);
#endif

  return yyresult;
}

#line 208 "tsl_parser.y"


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
