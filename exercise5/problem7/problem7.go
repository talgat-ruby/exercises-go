package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var results []string
	out := make(chan string)
	done := make(chan struct{})

	go func() {
		for val := range ch1 {
			out <- val
		}
		done <- struct{}{}
	}()

	go func() {
		for val := range ch2 {
			out <- val
		}
		done <- struct{}{}
	}()

	go func() {
		<-done
		<-done
		close(out)
	}()

	for val := range out {
		results = append(results, val)
	}

	return results
}
