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

func addConcurrently(numbers []int) int64 {
	if len(numbers) == 0 {
		return 0
	}
	lenNum := len(numbers)
	var sum1, sum2, sum3, sum4 int64
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i := 0; i < lenNum/4; i++ {
			sum1 += int64(numbers[i])
		}
	}()

	go func() {
		defer wg.Done()
		for i := lenNum / 4; i < lenNum/2; i++ {
			sum2 += int64(numbers[i])
		}
	}()

	go func() {
		defer wg.Done()
		for i := lenNum / 2; i < lenNum/4*3; i++ {
			sum3 += int64(numbers[i])
		}
	}()

	go func() {
		defer wg.Done()
		for i := lenNum / 4 * 3; i < lenNum; i++ {
			sum4 += int64(numbers[i])
		}
	}()

	wg.Wait()

	return sum1 + sum2 + sum3 + sum4
}
