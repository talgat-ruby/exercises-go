package problem10

func factory() (map[string]int, func(brand string) func(count int)) {
	outputMap := make(map[string]int)
	makeBrand := func(brand string) func(count int) {
		if _, exists := outputMap[brand]; !exists {
			outputMap[brand] = 0
		}

		return func(count int) {
			outputMap[brand] += count
		}
	}
	return outputMap, makeBrand
}
