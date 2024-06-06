package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
)

//go:embed students-branch.txt
var studentsData string

func main() {
	cmdCreateBranches := newSubcmdCreateBranches()
	cmdMergeMain := newSubcmdMergeMain()

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Printf(
			"one of subcommands: %v is required\n",
			[]string{
				cmdCreateBranches.getCmdName(),
				cmdMergeMain.getCmdName(),
			},
		)
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case cmdCreateBranches.getCmdName():
		if err := cmdCreateBranches.parse(os.Args[2:]); err != nil {
			fmt.Printf("error in cmd %s: %+v", cmdCreateBranches.getCmdName(), err)
			cmdCreateBranches.printDefaults()
			os.Exit(1)
		}
	case cmdMergeMain.getCmdName():
		if err := cmdMergeMain.parse(os.Args[2:]); err != nil {
			fmt.Printf("error in cmd %s: %+v", cmdMergeMain.getCmdName(), err)
			cmdMergeMain.printDefaults()
			os.Exit(1)
		}
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
