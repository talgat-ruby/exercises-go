package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(int) {
		if _, ok := brands[brand]; !ok {
			brands[brand] = 0
		}

		return func(incr int) {
			brands[brand] += incr
		}
	}

	return brands, makeBrand
}
