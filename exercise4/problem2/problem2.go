package problem2

import (
	"runtime"
	"sync"
)

// add - sequential code to add numbers, don't update it, just to illustrate concept
func add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

func addConcurrently(numbers []int) int64 {
	var sum int64
	numCores := runtime.NumCPU() // Получаем количество доступных ядер
	runtime.GOMAXPROCS(numCores) // Устанавливаем использование всех ядер

	partSize := (len(numbers) + numCores - 1) / numCores // Определяем размер части
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < numCores; i++ {
		start := i * partSize
		end := start + partSize
		if end > len(numbers) {
			end = len(numbers)
		}
		wg.Add(1)

		go func(part []int) {
			defer wg.Done()
			var partSum int64
			for _, n := range part {
				partSum += int64(n)
			}
			mu.Lock()
			sum += partSum
			mu.Unlock()
		}(numbers[start:end])
	}

	wg.Wait()
	return sum
}
