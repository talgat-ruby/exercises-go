package problem9

func factory(num int) func(...int) []int {
	// Возвращаемая функция принимает переменное количество аргументов типа int и возвращает срез int.
	return func(args ...int) []int {
		// Создаем срез для хранения результатов.
		result := make([]int, len(args))
		// Заполняем срез, умножая каждый аргумент на число num.
		for i, arg := range args {
			result[i] = arg * num
		}
		return result
	}
}
