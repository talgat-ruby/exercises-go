package problem9

import "fmt"

func factory(multiplier int) func(nums ...int) []int {

	return func(nums ...int) []int {
		result := make([]int, len(nums))

		for i, num := range nums {
			result[i] = num * multiplier
		}

		return result
	}
}

func main() {
	first := factory(15)
	fmt.Println(first(2, 3, 4))

	second := factory(2)
	fmt.Println(second(1, 2, 3, 4))
}
