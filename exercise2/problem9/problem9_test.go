package problem9

import (
	"slices"
	"testing"
)

func TestFactory(t *testing.T) {
	table := []struct {
		multiple int
		list     []int
		exp      []int
	}{
		{
			15,
			[]int{2, 3, 4},
			[]int{30, 45, 60},
		},
		{
			2,
			[]int{1, 2, 3},
			[]int{2, 4, 6},
		},
		{
			6,
			[]int{10, 5},
			[]int{60, 30},
		},
		{
			7,
			[]int{2, 3, 7},
			[]int{14, 21, 49},
		},
		{
			5,
			[]int{2, 1, 4},
			[]int{10, 5, 20},
		},
		{
			10,
			[]int{10, 1, 6},
			[]int{100, 10, 60},
		},
	}

	for _, r := range table {
		curry := factory(r.multiple)
		out := curry(r.list...)
		if !slices.Equal(out, r.exp) {
			t.Errorf("factory(%d)(%v) was incorrect, got: %v, expected: %v.", r.multiple, r.list, out, r.exp)
		}
	}
}
