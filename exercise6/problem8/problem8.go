package problem8

func multiplex(chs []<-chan string) []string {
	var result []string
	d := make(chan bool)

	for _, ch := range chs {
		go func(ch <-chan string) {
			for v := range ch {
				result = append(result, v)
			}
			d <- true
		}(ch)
	}

	for range chs {
		<-d
	}
	return result
}
