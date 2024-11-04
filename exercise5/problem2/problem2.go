package problem2

import "sync"

// add - sequential code to add numbers, don't update it, just to illustrate concept
func add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

// Функция addConcurrently — параллельное сложение чисел
func addConcurrently(numbers []int) int64 {
	var sum int64
	var wg sync.WaitGroup
	sumChan := make(chan int64) // Небуферизированный канал

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sumChan <- int64(n) // Отправляем число в канал
		}(num)
	}

	// Горутина для чтения из канала и суммирования
	go func() {
		wg.Wait()      // Ждем завершения всех горутин
		close(sumChan) // Закрываем канал после завершения всех горутин
	}()

	// Читаем из канала и суммируем
	for partialSum := range sumChan {
		sum += partialSum
	}

	return sum
}
