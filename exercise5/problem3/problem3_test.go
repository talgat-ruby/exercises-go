package problem3

import (
	"testing"
)

func TestSum(t *testing.T) {
	table := []struct {
		a, b int
		exp  int
	}{
		{2, 100, 102},
		{0, 0, 0},
		{30, 40, 70},
		{10, 50, 60},
		{1, 219, 220},
		{25, 219, 244},
		{1, 0, 1},
	}

	for _, r := range table {
		out := sum(r.a, r.b)
		if out != r.exp {
			t.Errorf("sum(%d, %d) was incorrect, got: %v, expected: %v.", r.a, r.b, out, r.exp)
		}
	}
}
