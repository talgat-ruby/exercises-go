package problem10

func factory() (map[string]int, func(brand string) func(n int)) {
	m := make(map[string]int)
	f := func(brand string) func(n int) {
		m[brand] = 0
		f1 := func(n int) {
			m[brand] += n
		}
		return f1
	}
	return m, f
}
