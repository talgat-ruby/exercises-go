package problem11

func keysAndValues[K comparable, V any](inp map[K]V) ([]K, []V) {
	var keys []K
	var values []V

	if inp == nil {
		return keys, values
	}

	// Создаем срезы для ключей и значений
	for key, value := range inp {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}
