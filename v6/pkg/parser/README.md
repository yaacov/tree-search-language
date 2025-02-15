# TSL Parser Generation

This directory contains the source files needed to generate the Tree Search Language (TSL) parser.

## Files

- `tsl_parser.y` - Bison grammar file that defines the TSL syntax rules and AST construction
- `tsl_lexer.l` - Flex lexer file that defines the token patterns for the TSL language
- `build.sh` - Shell script that runs the code generation tools

## Code Generation

The parser code is automatically generated using the `go generate` command. The generation process is triggered by the following directive in `pkg/tsl/tsl.go`:

```go
//go:generate bash ../parser/build.sh
```

To stop autogeneration, remove or comment this line from `pkg/tsl/tsl.go`.

## Requirements

To generate the parser code, you need:

1. **Build Tools**
   - GNU Bison (3.0 or later)
   - Flex (2.6 or later)
   - GCC or compatible C compiler
   - Bash shell

## Generation Process

When you run `go generate`, the following steps occur:

1. `build.sh` is executed
2. Bison processes `tsl_parser.y` to generate:
   - `tsl_parser.tab.c` (parser implementation)
   - `tsl_parser.tab.h` (parser header)
3. Flex processes `tsl_lexer.l` to generate:
   - `lex.yy.c` (lexer implementation)

The generated files are placed in the `pkg/tsl/` directory and are used by the Go code through CGo bindings.

## Manual Generation

You can also generate the parser manually:

```bash
cd pkg/parser
./build.sh
```
