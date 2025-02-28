package problem3

func sum(a, b int) int {
	var c int
	ch := make(chan int) // создаем канал для передачи результата

	go func(a, b int) {
		ch <- a + b // отправляем результат в канал
	}(a, b)

	c = <-ch // получаем результат из канала и присваиваем его переменной c

	return c
}
