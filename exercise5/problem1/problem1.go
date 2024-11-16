package problem1

func incrementConcurrently(num int) int {
	done := make(chan int)

	go func() {
		num++
		done <- num
	}()

	return <-done
}
