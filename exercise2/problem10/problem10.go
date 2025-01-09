package problem10

func factory() (map[string]int, func(string2 string) func(int)) {
	brands := make(map[string]int)
	makeBrand := func(brandName string) func(int) {
		brands[brandName] = 0
		return func(inc int) {
			brands[brandName] = brands[brandName] + inc
		}
	}

	return brands, makeBrand
}
