package main

import (
	"fmt"
	"os/exec"
)

func createRemoteBranch(branch string) error {
	if err := switchToBranch(branch); err != nil {
		return fmt.Errorf("could not switch to branch %s: %w", branch, err)
	}

	if err := pushRemoteBranch(branch); err != nil {
		return fmt.Errorf("could not push remote branch %s: %w", branch, err)
	}

	if err := switchToBranch("main"); err != nil {
		return fmt.Errorf("could not switch to main branch: %w", err)
	}

	return nil
}

func switchToBranch(branch string) error {
	cmd := exec.Command("git", "switch", "-C", branch)
	return cmd.Run()
}

func mergeFromLocal(branch string) error {
	cmd := exec.Command("git", "merge", branch)
	return cmd.Run()
}

func pushRemoteBranch(branch string) error {
	cmd := exec.Command("git", "push", "--set-upstream", "origin", branch)
	return cmd.Run()
}
