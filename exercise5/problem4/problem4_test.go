package problem4

import (
	"testing"
)

func TestSum(t *testing.T) {
	table := []struct {
		nums []int
		exp  int
	}{
		{nums: []int{3, 2, 1, 0}, exp: 6},
		{nums: []int{4, 3, 2, 1}, exp: 10},
		{nums: []int{5, 4, 3, 2}, exp: 14},
		{nums: []int{6, 5, 4, 3}, exp: 18},
	}

	for _, r := range table {
		out := sum(r.nums)
		if out != r.exp {
			t.Errorf("sum(%v) was incorrect, got: %v, expected: %v.", r.nums, out, r.exp)
		}
	}
}
