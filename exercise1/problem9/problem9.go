package problem9

func factory(multiplier int) func(numbers ...int) []int {
	return func(numbers ...int) []int {
		result := make([]int, len(numbers))
		for i, num := range numbers {
			result[i] = num * multiplier
		}
		return result
	}
}
