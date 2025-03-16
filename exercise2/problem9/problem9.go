package problem9

func factory(a int) func (nums ...int) []int {
	return func(nums ...int) []int {
		res := make([]int, len(nums))
		for i, num := range nums {
			res[i] = num * a
		}
		return res
	}
}
