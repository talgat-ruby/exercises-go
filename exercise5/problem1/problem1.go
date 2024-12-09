package problem1

import "time"

func incrementConcurrently(num int) int {
	go func() {
		num++
	}()
	time.Sleep(1 * time.Second)
	return num
}
