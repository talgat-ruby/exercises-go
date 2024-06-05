package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(int) {
		_, found := brands[brand]
		if !found {
			brands[brand] = 0
		}

		return func(incrementor int) {
			brands[brand] += incrementor
		}
	}

	return brands, makeBrand
}
