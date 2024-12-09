package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(int) {
		if _, exists := brands[brand]; !exists {
			brands[brand] = 0
		}

		return func(amount int) {
			brands[brand] += amount
		}
	}

	return brands, makeBrand
}
