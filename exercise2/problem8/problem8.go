package problem8

func simplify(list []string) map[string]int {
	indMap := make(map[string]int) // Создаем карту с использованием make

	// Передаем карту и срез напрямую без указателей
	load(indMap, list)

	return indMap
}

func load(m map[string]int, students []string) {
	for i, name := range students {
		m[name] = i // Добавляем пару (name: i) в карту
	}
}
