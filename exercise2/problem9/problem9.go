package problem9

func factory(multiple int) func(...int) []int {
	return func(a ...int) []int {
		var result []int
		for _, i := range a {
			result = append(result, i*multiple)
		}

		return result
	}
}
