package problem9

func factory(factor int) func(nums ...int) []int {
	return func(nums ...int) []int {
		output := make([]int, len(nums))
		for i, num := range nums {
			output[i] = factor * num
		}
		return output
	}
}
