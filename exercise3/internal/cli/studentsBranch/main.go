package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

//go:embed students-branch.txt
var studentsData string

func main() {
	students := strings.Split(studentsData, "\n")
	sort.Strings(students)

	if err := switchToBranch("main"); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	for _, student := range students {
		if err := createRemoteBranch(student); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("created remote branch %s\n", student)
	}

	os.Exit(0)
}

func switchToBranch(branch string) error {
	cmd := exec.Command("git", "switch", "-C", branch)
	return cmd.Run()
}

func createRemoteBranch(branch string) error {
	if err := switchToBranch(branch); err != nil {
		log.Fatal(err)
	}

	if err := pushRemoteBranch(branch); err != nil {
		return err
	}

	if err := switchToBranch("main"); err != nil {
		log.Fatal(err)
	}

	return nil
}

func pushRemoteBranch(branch string) error {
	cmd := exec.Command("git", "push", "--set-upstream", "origin", branch)
	return cmd.Run()
}
