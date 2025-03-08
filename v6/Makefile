#------------------------------------------------------------------------------
# Compiler and tools
#------------------------------------------------------------------------------
GO = go
GOLINT = $(shell which golangci-lint 2>/dev/null || echo $(HOME)/go/bin/golangci-lint)

#------------------------------------------------------------------------------
# Tools and flags
#------------------------------------------------------------------------------
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
.PHONY: all clean help generate lint test test-coverage install-tools format generate-parser test-stability

# Default target
all: generate tsl tsl_parser tsl_mem

help:
	@echo "Tree Search Language (TSL) Makefile Help"
	@echo "========================================"
	@echo ""
	@echo "Build targets:"
	@echo "  all               : Generate parser code and build all binaries (default)"
	@echo "  tsl               : Build the main TSL binary"
	@echo "  tsl_parser        : Build the TSL parser binary"
	@echo "  tsl_mem           : Build the TSL memory binary"
	@echo ""
	@echo "Development targets:"
	@echo "  generate          : Generate code using go generate"
	@echo "  generate-parser   : Generate parser code only"
	@echo "  format            : Format code using go fmt"
	@echo "  lint              : Run linter on code"
	@echo "  install-tools     : Install required development tools"
	@echo ""
	@echo "Testing targets:"
	@echo "  test              : Run all tests including stability tests"
	@echo "  test-stability    : Run stability tests only"
	@echo "  test-coverage     : Run tests with coverage report"
	@echo ""
	@echo "Cleanup targets:"
	@echo "  clean             : Remove build artifacts"
	@echo "  clean-all         : Remove all generated and build artifacts"

tsl: generate
	$(GO) build -o $(GO_BIN) ./$(GO_CMD)

tsl_parser: generate
	$(GO) build -o $(GO_PARSER_BIN) ./$(GO_PARSER_CMD)

tsl_mem: generate
	$(GO) build -o $(GO_MEM_BIN) ./$(GO_MEM_CMD)

generate:
	$(GO) generate ./...

format:
	$(GO) fmt ./...

lint:
	@if ! command -v $(GOLINT) >/dev/null 2>&1; then \
		echo "golangci-lint is not installed. Please run 'make install-tools' first."; \
		exit 1; \
	fi
	$(GOLINT) run ./...

test: tsl test-stability
	$(GO) test $(GO_TEST_FLAGS) ./...

test-stability: tsl
	@echo "Running stability tests..."
	@bash test/stability_test.sh

test-coverage: tsl
	$(GO) test $(GO_TEST_COVERAGE_FLAGS) ./...
	$(GO) tool cover -html=coverage.out

generate-parser:
	cd pkg/parser && ./build.sh

clean:
	rm -rf $(GO_BIN) $(GO_PARSER_BIN) $(GO_MEM_BIN) coverage.out

clean-all:
	rm -rf pkg/tsl/tsl_parser.tab.* pkg/tsl/lex.yy.c $(GO_BIN) $(GO_PARSER_BIN) $(GO_MEM_BIN) coverage.out

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

