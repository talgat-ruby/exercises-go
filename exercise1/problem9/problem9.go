package problem9

func factory(multiplier int) func(...int) []int {
	return func(numbers ...int) []int {
		var result []int
		for _, number := range numbers {
			result = append(result, number*multiplier)
		}
		return result
	}
}
