package problem9

func factory(multiple int) func(...int) []int {
	return func(list ...int) []int {
		result := make([]int, len(list))
		for i, val := range list {
			result[i] = val * multiple
		}
		return result
	}
}
