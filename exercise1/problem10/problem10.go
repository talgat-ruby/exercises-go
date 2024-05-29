package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(int) {
		brands[brand] = 0
		return func(increment int) {
			brands[brand] += increment
		}
	}

	return brands, makeBrand
}
