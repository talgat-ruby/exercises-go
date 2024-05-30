package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	brand := func(name string) func(int) {
		brands[name] = 0
		return func(cnt int) {
			brands[name] += cnt
		}
	}
	return brands, brand
}
