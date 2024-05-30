package problem9

func factory(a int) func(v ...int) []int {
	return func(v ...int) []int {
		res := make([]int, len(v))
		for i := 0; i < len(v); i++ {
			res[i] = a * v[i]
		}
		return res
	}
}
