package problem9

func factory(mult int) func(numbs ...int) []int {
	return func(numbs ...int) []int {
		output := make([]int, len(numbs))
		for i, numb := range numbs {
			output[i] = numb * mult
		}
		return output
	}
}
