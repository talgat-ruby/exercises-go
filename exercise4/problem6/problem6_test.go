package problem6

import (
	"slices"
	"testing"

	"github.com/talgat-ruby/exercises-go/pkg/util"
)

func sendNums(out chan<- int, nums []int) {
	for _, n := range nums {
		out <- n
	}
	close(out)
}

func receiveResults(in <-chan int) []int {
	nums := make([]int, 0)
	for n := range in {
		nums = append(nums, n)
	}
	return nums
}

func TestPiper(t *testing.T) {
	util.SkipTestOptional(t)

	table := []struct {
		nums []int
		exp  []int
	}{
		{
			[]int{1, 2, 3},
			[]int{7, 9, 11},
		},
		{
			[]int{2, 3, 4},
			[]int{9, 11, 13},
		},
		{
			[]int{10, 1, 6},
			[]int{25, 7, 17},
		},
	}

	for _, r := range table {
		in := make(chan int)
		out := piper(in, []pipe{multiplyBy2, add5})
		go sendNums(in, r.nums)
		results := receiveResults(out)

		if !slices.Equal(results, r.exp) {
			t.Errorf("piper() was incorrect, got: %v, expected: %v.", results, r.exp)
		}
	}
}
