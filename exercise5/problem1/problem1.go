package problem1

func incrementConcurrently(num int) int {

	result := make(chan int)

	go func() {
		result <- num + 1
	}()

	return <-result
}
