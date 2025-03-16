package problem1

func incrementConcurrently(num int) int {
	ch := make(chan bool)

	go func(ch chan bool) {
		num++
		ch <- true
	}(ch)

	<-ch

	return num
}
