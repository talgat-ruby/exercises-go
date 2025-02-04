package problem10

func factory() (map[string]int, func(brand string) func(int)) {
	result := make(map[string]int)
	return result, func(brand string) func(int) {
		result[brand] = 0
		return func(i int) {
			result[brand] = result[brand] + i
		}
	}
}
