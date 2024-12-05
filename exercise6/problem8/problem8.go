package problem8

func multiplex(chs []<-chan string) []string {
	var slc []string
	ch1 := make(chan string)
	ch2 := make(chan struct{}, len(chs))
	for _, ch := range chs {
		go func(ch <-chan string) {
			for val := range ch {
				ch1 <- val
			}
			ch2 <- struct{}{}
		}(ch)
	}

	go func() {
		for i := 0; i < len(chs); i++ {
			<-ch2
		}
		close(ch1)
	}()
	for val := range ch1 {
		slc = append(slc, val)
	}

	return slc
}
