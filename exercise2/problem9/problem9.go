package problem9

func factory(num int) func(...int) []int {
	return func(number ...int) []int {
		res := make([]int, len(number))
		for i, n := range number {
			res[i] = n * num
		}
		return res
	}
}
