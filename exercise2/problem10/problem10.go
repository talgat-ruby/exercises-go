package problem10

func factory() (map[string]int, func(string) func(int)) {
	brandInfo := make(map[string]int)
	return brandInfo, func(brandName string) func(int) {
		brandInfo[brandName] = 0
		return func(incrementValue int) {
			brandInfo[brandName] += incrementValue
		}
	}
}
