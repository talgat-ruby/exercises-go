package problem9

func factory(factor int) func(...int) []int {
	return func(nums ...int) []int {
		results := make([]int, len(nums))
		for i, num := range nums {
			results[i] = num * factor
		}
		return results
	}
}
