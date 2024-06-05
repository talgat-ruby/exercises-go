package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)
	return brands, makeBrand(brands)
}

func makeBrand(brands map[string]int) func(string) func(int) {
	return func(brand string) func(int) {
		brands[brand] = 0
		return func(count int) {
			brands[brand] += count
		}
	}
}
