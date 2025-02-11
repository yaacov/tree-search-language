package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yaacov/tree-search-language/v6/pkg/tsl"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <expression>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nParses and displays the TSL expression tree\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	expression := flag.Arg(0)
	tree, err := tsl.ParseTSL(expression)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer tree.Free()

	fmt.Printf("Parsing: %s\n", expression)
	printer := NewASTPrinter()
	printer.Print(tree, 0)
}
