package problem9

func factory(multiple int) func(...int) []int {
	return func(nums ...int) []int {
		var result []int
		for _, num := range nums {
			result = append(result, num*multiple)
		}
		return result
	}
}
