package problem10

func factory() (map[string]int, func(string) func(int)) {

	brands := make(map[string]int)

	makeBrand := func(brandName string) func(int) {

		if _, exists := brands[brandName]; !exists {
			brands[brandName] = 0
		}

		return func(amount int) {
			brands[brandName] += amount
		}
	}

	return brands, makeBrand
}
