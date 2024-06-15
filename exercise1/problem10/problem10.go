package problem10

func factory() (map[string]int, func(brand string) func(i int)) {
	brands := make(map[string]int)

	makeBrand := func(brand string) func(i int) {
		_, ok := brands[brand]
		if !ok {
			brands[brand] = 0
		}
		return func(i int) {
			brands[brand] = brands[brand] + i
		}
	}

	return brands, makeBrand
}
