package problem11

func keysAndValues[T string | int, Y comparable](slice map[T]Y) ([]T, []Y) {
	keys := make([]T, 0, len(slice))
	values := make([]Y, 0, len(slice))

	for key := range slice {
		keys = append(keys, key)
	}

	// sort
	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				temp := keys[j]
				keys[j] = keys[i]
				keys[i] = temp
			}
		}
	}

	for _, key := range keys {
		values = append(values, slice[key])
	}

	return keys, values
}
