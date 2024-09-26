package problem10

func factory() (map[string]int, func(string) func(int)) {

	brands := make(map[string]int)

	createBrand := func(brand string) func(int) {
		return func(amount int) {
			brands[brand] += amount
		}
	}
	return brands, createBrand
}
