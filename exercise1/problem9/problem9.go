package problem9

func factory(multiple int) func(arr ...int) []int {
	f := func(arr ...int) []int {
		for i := 0; i < len(arr); i++ {
			arr[i] *= multiple
		}
		return arr
	}
	return f
}
