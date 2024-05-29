package problem9

func factory(multiplier int) func(...int) []int {
	return func(args ...int) []int {
		result := make([]int, len(args))
		for i := 0; i < len(result); i++ {
			result[i] = multiplier * int(args[i])
		}
		return result
	}
}
