package problem1

import (
	"testing"
)

func TestIsChangeEnough(t *testing.T) {
	table := []struct {
		changes [4]int
		total   float32
		exp     bool
	}{
		{[4]int{2, 100, 0, 0}, 14.11, false},
		{[4]int{0, 0, 20, 5}, 0.75, true},
		{[4]int{30, 40, 20, 5}, 12.55, true},
		{[4]int{10, 0, 0, 50}, 13.85, false},
		{[4]int{1, 0, 5, 219}, 19.99, false},
		{[4]int{1, 0, 2555, 219}, 127.75, true},
		{[4]int{1, 335, 0, 219}, 35.21, true},
	}

	for _, r := range table {
		out := isChangeEnough(r.changes, r.total)
		if out != r.exp {
			t.Errorf("isChangeEnough(%v, %f) was incorrect, got: %t, expected: %t.",
				r.changes, r.total, out, r.exp)
		}
	}
}
