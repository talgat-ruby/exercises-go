package main

import (
	"flag"
	"fmt"
	"github.com/webaiz/exercise7/blogging-platform/internal/db/seeds"
)

type subcmdSeed struct {
	cmd *flag.FlagSet
}

func newSubcmdSeed() subcmd {
	cmd := flag.NewFlagSet(cmdNameSeed, flag.ExitOnError)

	return &subcmdSeed{
		cmd,
	}
}

func (s *subcmdSeed) getCmdName() string {
	return s.cmd.Name()
}

func (s *subcmdSeed) printDefaults() {
	s.cmd.PrintDefaults()
}

func (s *subcmdSeed) parse(args []string) error {
	if err := s.cmd.Parse(args); err != nil {
		return err
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	if !s.cmd.Parsed() {
		return fmt.Errorf("please provide correct arguments to %s command", s.cmd.Name())
	}

	seeder, err := seeds.New()
	if err != nil {
		return err
	}

	err = seeder.Populate()
	if err != nil {
		return err
	}

	return nil
}
