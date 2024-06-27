package problem9

func factory(multiple int) func(...int) []int {
	return func(list ...int) []int {
		for i := range list {
			list[i] = multiple * list[i]
		}
		return list
	}
}
