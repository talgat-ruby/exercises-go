package main

import (
	"testing"
)

func TestGreetings(t *testing.T) {
	table := []struct {
		exp string
	}{
		{"Hello World!"},
	}

	for _, r := range table {
		out := greetings()
		if out != r.exp {
			t.Errorf("greetings() was incorrect, expected: %s, got: %s.", r.exp, out)
		}
	}
}
