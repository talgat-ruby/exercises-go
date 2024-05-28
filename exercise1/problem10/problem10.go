package problem10

func factory() (map[string]int, func(brandName string) func(increment int)) {
	brand := map[string]int{}
	return brand, func(brandName string) func(increment int) {
		brand[brandName] = 0
		return func(increment int) {
			brand[brandName] += increment
		}
	}
}
