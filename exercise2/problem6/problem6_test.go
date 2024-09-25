package problem6

import (
	"testing"

	"github.com/talgat-ruby/exercises-go/pkg/util"
)

func TestSumOfTwo(t *testing.T) {
	util.SkipTestOptional(t)

	table := []struct {
		a   []int
		b   []int
		sum int
		exp bool
	}{
		{
			[]int{1, 2, 3},
			[]int{10, 20, 30, 40, 50},
			42,
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{10, 20, 30, 40, 50},
			11,
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{10, 20, 30, 40, 50},
			60,
			false,
		},
		{
			[]int{1, 2, 3},
			[]int{10, 20, 30, 40, 50},
			53,
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{10, 20, 30, 40, 50},
			4,
			false,
		},
		{
			[]int{1, 2},
			[]int{4, 5, 6},
			5,
			true,
		},
		{
			[]int{1, 2},
			[]int{4, 5, 6},
			8,
			true,
		},
		{
			[]int{1, 2},
			[]int{4, 5, 6},
			3,
			false,
		},
		{
			[]int{1, 2},
			[]int{4, 5, 6},
			9,
			false,
		},
	}

	for _, r := range table {
		out := sumOfTwo(r.a, r.b, r.sum)
		if out != r.exp {
			t.Errorf("sumOfTwo(%v, %v, %d) was incorrect, got: %t, expected: %t.", r.a, r.b, r.sum, out, r.exp)
		}
	}
}
