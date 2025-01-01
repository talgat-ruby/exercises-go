package problem8

func multiplex(chs []<-chan string) []string {
	out := make([]string, 0)
	done := make(chan struct{})
	var remaining = len(chs)

	for _, ch := range chs {
		go func(c <-chan string) {
			for msg := range c {
				out = append(out, msg)
			}
			remaining--
			if remaining == 0 {
				close(done)
			}
		}(ch)
	}

	<-done

	return out
}
