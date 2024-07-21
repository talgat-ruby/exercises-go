package problem1

func incrementConcurrently(num int) int {
	done := make(chan bool)
	go func() {
		num++
		done <- true
	}()
	<-done
	return num
}
