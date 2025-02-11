#!/bin/bash
set -e

# Change to script's directory
cd "$(dirname "${BASH_SOURCE[0]}")"

# Check for required tools
command -v bison >/dev/null 2>&1 || { echo "Error: bison is required but not installed" >&2; exit 1; }
command -v flex >/dev/null 2>&1 || { echo "Error: flex is required but not installed" >&2; exit 1; }

# Generate parser and lexer
bison -d -o ../tsl/tsl_parser.tab.c tsl_parser.y
flex -o ../tsl/lex.yy.c tsl_lexer.l
