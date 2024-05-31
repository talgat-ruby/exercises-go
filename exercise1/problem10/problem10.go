package problem10

func factory() (map[string]int, func(string) func(int)) {
	// Инициализация карты брендов.
	brands := make(map[string]int)

	// Определение функции makeBrand для создания инкрементеров брендов.
	makeBrand := func(brand string) func(int) {
		// Инициализация счетчика для бренда.
		counter := 0
		brands[brand] = counter

		// Возвращаем функцию инкрементации.
		return func(inc int) {
			counter += inc
			brands[brand] = counter
		}
	}

	// Возвращаем карту брендов и функцию makeBrand.
	return brands, makeBrand
}
