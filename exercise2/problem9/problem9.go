package problem9

func factory(a int) func(vals ...int) []int {
	fun := func(vals ...int) []int {
		for i, val := range vals {
			vals[i] = val * a
		}
		return vals
	}

	return fun
}
