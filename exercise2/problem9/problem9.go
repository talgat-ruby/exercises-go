package problem9

func factory(multiple int) func(nums ...int) []int {
	return func(nums ...int) []int {
		result := make([]int, 0)
		for _, v := range nums {
			result = append(result, v*multiple)
		}
		return result
	}
}
