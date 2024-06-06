package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"
)

type subcmdMergeMain struct {
	cmd      *flag.FlagSet
	students []string
}

func newSubcmdMergeMain() *subcmdMergeMain {
	cmd := flag.NewFlagSet(cmdNameMergeMain, flag.ExitOnError)

	students := strings.Split(strings.TrimSpace(studentsData), "\n")
	sort.Strings(students)

	return &subcmdMergeMain{
		cmd:      cmd,
		students: students,
	}
}

func (s *subcmdMergeMain) getCmdName() string {
	return s.cmd.Name()
}

func (s *subcmdMergeMain) printDefaults() {
	s.cmd.PrintDefaults()
}

func (s *subcmdMergeMain) parse(args []string) error {
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
		if err := switchToBranch(student); err != nil {
			return fmt.Errorf("could not switch to %s branch: %w", student, err)
		}
		if err := mergeFromLocal("main"); err != nil {
			return fmt.Errorf("could not merge main to %s branch: %w", student, err)
		}
		if err := pushRemoteBranch(student); err != nil {
			return fmt.Errorf("could not push to %s branch: %w", student, err)
		}
		fmt.Printf("merged main to %s\n", student)
	}

	if err := switchToBranch("main"); err != nil {
		return fmt.Errorf("could not switch to main branch: %w", err)
	}

	return nil
}
