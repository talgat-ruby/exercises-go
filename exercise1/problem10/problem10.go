package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeNewBrand := func(newBrand string) func(newCount int) {
		_, contains := brands[newBrand]
		if !contains {
			brands[newBrand] = 0
		}

		return func(newCount int) {
			brands[newBrand] += newCount
		}
	}

	return brands, makeNewBrand
}
