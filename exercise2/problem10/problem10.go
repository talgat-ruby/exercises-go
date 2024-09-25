package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)

	makeB := func(brand string) func(int) {
		return func(count int) {
			brands[brand] += count
		}
	}

	return brands, makeB
}
