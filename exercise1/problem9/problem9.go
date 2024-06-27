package problem9

func factory(mult int) func(...int) []int {
	return func(values ...int) []int {
		for i, x := range values {
			values[i] = x * mult
		}

		return values
	}
}
