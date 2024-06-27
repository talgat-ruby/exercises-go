package problem9

func factory(multiplier int) func(...int) []int {
	return func(numbers ...int) []int {
		outNumbers := make([]int, len(numbers))
		for i, number := range numbers {
			outNumbers[i] = number * multiplier
		}
		return outNumbers
	}
}
