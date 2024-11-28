package problem10

func factory() (map[string]int, func(brand string) func(value int)) {
	brands := make(map[string]int)
	increase := func(brand string) func(value int) {
		brands[brand] = 0
		return func(value int) {
			brands[brand] = value + brands[brand]
		}
	}

	return brands, increase
}
