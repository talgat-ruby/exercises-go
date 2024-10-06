package problem9

import (
	"fmt"
)

func factory(multiplier int) func(nums ...int) []int {
	return func(nums ...int) []int {

		results := make([]int, len(nums))
		for i, num := range nums {
			results[i] = num * multiplier
		}
		return results
	}
}

func main() {

	first := factory(15)
	fmt.Println(first(2, 3, 4))

	second := factory(2)
	fmt.Println(second(1, 2, 3, 4))
}
