package problem9

func factory(a int) func(l ...int) []int {
	fn := func(l ...int) []int {
		var result []int

		for _, num := range l {
			result = append(result, num*a)
		}

		return result
	}

	return fn
}
