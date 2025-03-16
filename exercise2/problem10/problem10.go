package problem10

func factory() (map[string]int, func(string) func(int)) {
	m := map[string]int{}
	f := func(model string) func(int) {
		m[model] = 0
		return func(inc int) {
			m[model] += inc
		}
	}
	return m, f
}
