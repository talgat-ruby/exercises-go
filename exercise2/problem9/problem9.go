package problem9

func factory(n int) func(...int) []int {
	return func(a ...int) []int {
		s := make([]int, len(a))
		for i, v := range a {
			s[i] = v * n
		}
		return s
	}
}
