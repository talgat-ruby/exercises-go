package problem9

func factory(number int) func(nums ...int) []int {
	return func(nums ...int) []int {
		result := make([]int, len(nums))
		for i, num := range nums {
			result[i] = num * number
		}
		return result
	}
}
