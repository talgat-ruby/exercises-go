package problem1

func incrementConcurrently(num int) int {
	ch := make(chan bool)
	go func() {
		num++
		ch <- true
	}()
	<-ch
	return num
}
