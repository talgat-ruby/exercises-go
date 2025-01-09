package problem5

func products(obj map[string]int, amount int) []string {
	keys := make([]string, 0)
	reversedKeys := make([]string, 0)

	for key, value := range obj {
		if value > amount {
			keys = append(keys, key)
		}
	}

	// sort
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if obj[keys[i]] < obj[keys[j]] {
				temp := keys[j]
				keys[j] = keys[i]
				keys[i] = temp
			}
		}
	}

	return reversedKeys
}
