package problem9

func factory(multiplier int) func(nums ...int) []int {
	return func(nums ...int) []int {
		results := make([]int, len(nums))
		for i, num := range nums {
			results[i] = num * multiplier
		}
		return results
	}
}
