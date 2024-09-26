package problem8

func simplify(slice []string) map[string]int {
	result := make(map[string]int)

	for i, value := range slice {
		result[value] = i
	}

	return result
}
