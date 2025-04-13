package problem10

func factory() (map[string]int, func(brand string) func(int)) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(int) {
		if _, exists := brands[brand]; !exists {
			brands[brand] = 0
		}

		return func(count int) {
			brands[brand] += count
		}
	}

	return brands, makeBrand
}
