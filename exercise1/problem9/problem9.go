package problem9

func factory(multiplier int) func(slice ...int) []int {
	return func(slice ...int) []int {
		for i := 0; i < len(slice); i++ {
			slice[i] = multiplier * slice[i]
		}
		return slice
	}
}
