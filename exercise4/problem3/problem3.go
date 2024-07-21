package problem3

func sum(a, b int) int {
	resultChan := make(chan int) // Создаем канал для передачи результата

	go func(a, b int) {
		result := a + b
		resultChan <- result // Отправляем результат в канал
	}(a, b)

	c := <-resultChan // Ожидаем результат из канала
	return c
}
