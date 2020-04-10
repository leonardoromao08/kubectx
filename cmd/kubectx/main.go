package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// parse command-line flags
	var op Op
	op = parseArgs(os.Args[1:])

	// TODO consider addin Run() operation to each operation type
	switch v := op.(type) {
	case HelpOp:
		printHelp(os.Stdout)
	case ListOp:
		printListContexts(os.Stdout)
	case SwitchOp:
		// TODO implement
		panic("not implemented")
	case UnknownOp:
		fmt.Printf("%s unsupported operation: %s\n",
			color.RedString("error:"),
			strings.Join(v.Args, " "))
		printHelp(os.Stdout)
		os.Exit(1)
	default:
		fmt.Printf("internal error: operation type %T not handled", op)
	}
}
