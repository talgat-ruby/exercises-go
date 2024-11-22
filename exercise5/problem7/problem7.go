package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var res []string
	done := make(chan struct{})

	go func() {
		for v := range ch1 {
			res = append(res, v)
		}
		done <- struct{}{}
	}()

	go func() {
		for v := range ch2 {
			res = append(res, v)
		}
		done <- struct{}{}
	}()

	<-done
	<-done

	return res

}
