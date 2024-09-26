package problem11

import (
	"sort"
)

func keysAndValues[K string | int, V any](m map[K]V) ([]K, []V) {
	if m == nil {
		return []K{}, []V{}
	}

	// Создаем срез для хранения ключей
	keys := make([]K, 0, len(m))

	// Заполняем срез ключами из карты
	for k := range m {
		keys = append(keys, k)
	}

	// Сортируем ключи
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Создаем срез для хранения значений
	values := make([]V, 0, len(m))

	// Заполняем срез значениями в порядке отсортированных ключей
	for _, k := range keys {
		values = append(values, m[k])
	}

	return keys, values
}
