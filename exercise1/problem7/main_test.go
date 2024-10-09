package main

import (
	"testing"
)

func TestHighestDigit(t *testing.T) {
	table := []struct {
		num int
		exp int
	}{
		{51, 5},
		{0, 0},
		{7495037, 9},
		{222222, 2},
	}

	for _, r := range table {
		out := highestDigit(r.num)
		if out != r.exp {
			t.Errorf("emojify(%d) was incorrect, expected: %d, got: %d.", r.num, r.exp, out)
		}
	}
}
