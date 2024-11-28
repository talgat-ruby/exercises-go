package main

import (
	"testing"
)

func TestNumberSquares(t *testing.T) {
	table := []struct {
		num int
		exp int
	}{
		{3, 14},
		{10, 385},
		{12, 650},
		{5, 55},
		{9, 285},
		{11, 506},
		{15, 1240},
	}

	for _, r := range table {
		out := numberSquares(r.num)
		if out != r.exp {
			t.Errorf("numberSquares(%d) was incorrect, expected: %d, got: %d.", r.num, r.exp, out)
		}
	}
}
