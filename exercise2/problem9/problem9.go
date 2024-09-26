package problem9

func factory(multiple int) func(nums ...int) []int {

	return func(nums ...int) []int {
		result := make([]int, len(nums))

		for i, num := range nums {
			result[i] = num * multiple
		}

		return result
	}
}
