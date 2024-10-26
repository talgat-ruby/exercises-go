package problem1

func incrementConcurrently(num int) int {
	go func() {
		num++
	}()

	return num
}
