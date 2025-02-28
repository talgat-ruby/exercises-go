package problem9

func factory(multiplier int) func(...int) []int {
	return func(nums ...int) []int {

		result := make([]int, len(nums))

		for i, num := range nums {
			result[i] = num * multiplier
		}

		return result
	}
}
