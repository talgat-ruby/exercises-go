package problem9

func factory(size int) func(args ...int) []int {
	return func(args ...int) []int {
		s := make([]int, 0)
		for _, arg := range args {
			s = append(s, arg*size)
		}
		return s
	}
}
