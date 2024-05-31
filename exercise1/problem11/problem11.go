package problem11

func removeDups[T comparable](input []T) []T {
	// Ключами являются элементы среза, а значениями - булевы значения, указывающие, встречался ли элемент уже или нет
	check := make(map[T]bool)
	result := []T{}
	// Если элемент встречается впервые, мы добавляем его в результат.
	for _, val := range input {
		if !check[val] {
			check[val] = true
			// возвращаем срез result, который содержит только уникальные элементы в том же порядке, что и в исходном срезе.
			result = append(result, val)
		}
	}
	return result
}

// обобщенные типы данных (generics) в Go, обозначенные символом [T comparable] после имени функции
