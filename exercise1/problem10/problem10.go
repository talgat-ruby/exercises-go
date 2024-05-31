package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)
	i := 0

	makeBrand := func(brand string) func(int) {
		brands[brand] = i

		inc := func(i int) {
			brands[brand] += i
		}

		return inc
	}

	return brands, makeBrand
}
