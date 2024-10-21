package problem10

func factory() (map[string]int, func(string) func(int)) {

	brands := make(map[string]int)

	makeBrand := func(name string) func(int) {
		brands[name] = 0
		return func(increment int) {
			brands[name] += increment
		}
	}

	return brands, makeBrand
}
