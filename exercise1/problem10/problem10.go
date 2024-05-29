package problem10

func factory() (map[string]int, func(string) func(int)) {
	brands := map[string]int{}
	return brands, func(brand string) func(int) {
		brands[brand] = 0
		return func(incrementVal int) {
			brands[brand] += incrementVal
		}
	}
}
