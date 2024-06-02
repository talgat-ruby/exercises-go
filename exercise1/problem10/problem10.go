package problem10

func factory() (map[string]int, func(brand string) func(increment int)) {
	m := map[string]int{}
	return m, func(brand string) func(increment int) {
		m[brand] = 0
		return func(inc int) {
			m[brand] += inc
		}
	}
}
