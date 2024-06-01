package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	for k, v := range brands {
		brands[k] = v
	}

	makeBrand := func(brand string) func(int) {
		if _, ok := brands[brand]; !ok {
			brands[brand] = 0
		}
		return func(increment int) {
			brands[brand] += increment
		}
	}

	return brands, makeBrand
}
