package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

type subcmdCreateBranches struct {
	cmd      *flag.FlagSet
	students []string
}

func newSubcmdCreateBranches() *subcmdCreateBranches {
	cmd := flag.NewFlagSet(cmdNameCreateBranches, flag.ExitOnError)

	students := strings.Split(strings.TrimSpace(studentsData), "\n")
	sort.Strings(students)

	return &subcmdCreateBranches{
		cmd:      cmd,
		students: students,
	}
}

func (s *subcmdCreateBranches) getCmdName() string {
	return s.cmd.Name()
}

func (s *subcmdCreateBranches) printDefaults() {
	s.cmd.PrintDefaults()
}

func (s *subcmdCreateBranches) parse(args []string) error {
	if err := s.cmd.Parse(args); err != nil {
		return err
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if !s.cmd.Parsed() {
		return fmt.Errorf("please provide correct arguments to %s command", s.cmd.Name())
	}

	if err := switchToBranch("main"); err != nil {
		return fmt.Errorf("could not switch to main branch: %w", err)
	}

	for _, student := range s.students {
		if err := createRemoteBranch(student); err != nil {
			return fmt.Errorf("could not create remote branch: %w", err)
		}
		fmt.Printf("created remote branch %s\n", student)
	}

	if err := switchToBranch("main"); err != nil {
		return fmt.Errorf("could not switch to main branch: %w", err)
	}

	return nil
}
