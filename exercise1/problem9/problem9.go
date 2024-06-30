package problem9

func factory(val int) func(...int) []int {
	return func(nums ...int) []int {
		for i, num := range nums {
			nums[i] = num * val
		}
		return nums
	}
}
