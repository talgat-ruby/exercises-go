package problem9

func factory(multiple int) func(...int) []int {
	return func(numbers ...int) []int {
		result := make([]int, len(numbers))

		for i, num := range numbers {
			result[i] = num * multiple
		}
		return result
	}
}
