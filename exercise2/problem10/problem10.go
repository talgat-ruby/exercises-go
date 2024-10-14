package problem10

import "sync"

func factory() (map[string]int, func(string) func(int)) {
	brands := make(map[string]int)
	var mu sync.Mutex

	makeBrand := func(brand string) func(int) {
		mu.Lock()
		if _, exists := brands[brand]; !exists {
			brands[brand] = 0
		}
		mu.Unlock()
		return func(increment int) {
			mu.Lock()
			brands[brand] += increment
			mu.Unlock()
		}
	}

	return brands, makeBrand
}
