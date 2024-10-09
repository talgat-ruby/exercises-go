package main

import (
	"testing"
)

func TestCountVowels(t *testing.T) {
	table := []struct {
		text string
		exp  int
	}{
		{"Celebration", 5},
		{"Palm", 1},
		{"Prediction", 4},
		{"Suite", 3},
		{"Quote", 3},
		{"Portrait", 3},
		{"Steam", 2},
		{"Tape", 2},
		{"Nightmare", 3},
		{"Convention", 4},
	}

	for _, r := range table {
		out := countVowels(r.text)
		if out != r.exp {
			t.Errorf("countVowels(%s) was incorrect, expected: %d, got: %d.", r.text, r.exp, out)
		}
	}
}
