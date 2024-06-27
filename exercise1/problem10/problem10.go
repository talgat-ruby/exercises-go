package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeBrand := func(mark string) func(int) {
		brands[mark] = 0
		return func(i int) {
			count, _ := brands[mark]
			brands[mark] = count + i
		}
	}

	return brands, makeBrand
}
