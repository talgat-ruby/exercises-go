package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)
	makeBrand := func(name string) func(int) {
		if _, exists := brands[name]; !exists {
			brands[name] = 0
		}
		return func(i int) {
			brands[name] += i
		}
	}
	return brands, makeBrand
}
