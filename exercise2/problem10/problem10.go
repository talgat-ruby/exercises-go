package problem10

func factory() (map[string]int, func(str string) func(int)) {
	res := make(map[string]int)
	return res, func(str string) func(int) {
		res[str] = 0
		return func(in int) {
			res[str] += in
		}
	}
}
