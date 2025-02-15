#------------------------------------------------------------------------------
# Compiler and tools
#------------------------------------------------------------------------------
GO = go

#------------------------------------------------------------------------------
# Tools and flags
#------------------------------------------------------------------------------
GOLINT = golangci-lint
GO_TEST_FLAGS = -v -race
GO_TEST_COVERAGE_FLAGS = -coverprofile=coverage.out

#------------------------------------------------------------------------------
# Directories
#------------------------------------------------------------------------------
GO_CMD = cmd/tsl
GO_PARSER_CMD = cmd/tsl_parser
GO_MEM_CMD = cmd/tsl_mem

#------------------------------------------------------------------------------
# Output files
#------------------------------------------------------------------------------
GO_BIN = tsl
GO_PARSER_BIN = tsl_parser
GO_MEM_BIN = tsl_mem

#------------------------------------------------------------------------------
# Phony targets
#------------------------------------------------------------------------------
.PHONY: all clean help generate lint test test-coverage test-file install-tools format

# Default target
all: generate tsl tsl_parser tsl_mem

help:
	@echo "Available targets:"
	@echo "  all           - Build everything (default)"
	@echo "  tsl           - Build Go CLI tool"
	@echo "  tsl_parser    - Build TSL parser tool"
	@echo "  tsl_mem       - Build TSL memory demo tool"
	@echo "  generate      - Run go generate"
	@echo "  lint          - Run linter checks"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  test-file     - Run tests for a specific file"
	@echo "  clean         - Remove built files and directories"
	@echo "  clean         - Remove built files and directories and generated files"
	@echo "  help          - Show this help message"
	@echo "  format        - Format the code using go fmt"

generate:
	$(GO) generate ./...

lint:
	$(GOLINT) run ./...

test: tsl
	$(GO) test $(GO_TEST_FLAGS) ./...

test-coverage: tsl
	$(GO) test $(GO_TEST_COVERAGE_FLAGS) ./...
	$(GO) tool cover -html=coverage.out

# Add new target for testing specific file
test-file:
	@if [ -z "$(FILE)" ]; then \
		echo "Please specify a test file with FILE=<path>"; \
		echo "Example: make test-file FILE=./pkg/walkers/sql/walk_test.go"; \
		exit 1; \
	fi
	$(GO) test $(GO_TEST_FLAGS) $(FILE)

format:
	$(GO) fmt ./...

tsl: generate
	$(GO) build -o $(GO_BIN) ./$(GO_CMD)

tsl_parser: generate
	$(GO) build -o $(GO_PARSER_BIN) ./$(GO_PARSER_CMD)

tsl_mem: generate
	$(GO) build -o $(GO_MEM_BIN) ./$(GO_MEM_CMD)

clean:
	rm -rf $(GO_BIN) $(GO_PARSER_BIN) $(GO_MEM_BIN) coverage.out

clean-all:
	rm -rf pkg/tsl/tsl_parser.tab.* pkg/tsl/lex.yy.c $(GO_BIN) $(GO_PARSER_BIN) $(GO_MEM_BIN) coverage.out
