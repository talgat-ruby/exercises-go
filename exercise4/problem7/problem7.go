package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	out := make(chan string, 2)
	done := make(chan bool)
	go func() {
		for len(out) < 2 {
			select {
			case x := <-ch1:
				out <- x
			case y := <-ch2:
				out <- y
			default:
				// wait
			}
		}
		done <- true
		close(out)
	}()
	<-done
	close(done)
	strings := make([]string, 0, 2)
	for v := range out {
		strings = append(strings, v)
	}
	return strings
}
