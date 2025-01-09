package problem9

func factory(num int) func(...int) []int {
	return func(args ...int) []int {
		for i, arg := range args {
			args[i] = num * arg
		}
		return args
	}
}
