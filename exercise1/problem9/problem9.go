package problem9

func factory(multiple int) func(nums ...int) []int {
	var res []int
	return func(nums ...int) []int {
		for _, i := range nums {
			res = append(res, i*multiple)
		}

		return res
	}
}
