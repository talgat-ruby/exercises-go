package problem9

func factory(num int) func(...int) []int {
	return func(a ...int) []int {
		for i := 0; i < len(a); i++ {
			a[i] = a[i] * num
		}
		return a
	}
}
