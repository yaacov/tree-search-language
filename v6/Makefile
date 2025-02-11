#------------------------------------------------------------------------------
# Compiler and tools
#------------------------------------------------------------------------------
GO = go

#------------------------------------------------------------------------------
# Directories
#------------------------------------------------------------------------------
GO_CMD = cmd/tsl

#------------------------------------------------------------------------------
# Output files
#------------------------------------------------------------------------------
GO_BIN = tsl

#------------------------------------------------------------------------------
# Phony targets
#------------------------------------------------------------------------------
.PHONY: all clean help generate

# Default target
all: generate tsl

help:
	@echo "Available targets:"
	@echo "  all      - Build everything (default)"
	@echo "  tsl      - Build Go CLI tool"
	@echo "  generate - Run go generate"
	@echo "  clean    - Remove built files and directories"
	@echo "  help     - Show this help message"

generate:
	$(GO) generate ./...

tsl: generate
	$(GO) build -o $(GO_BIN) ./$(GO_CMD)

clean:
	rm -rf pkg/tsl/tsl_parser.tab.* pkg/tsl/lex.yy.c $(GO_BIN)
