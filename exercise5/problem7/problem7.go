package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var output []string

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if ok {
				output = append(output, v)
			} else {
				ch1 = nil
			}
		case v, ok := <-ch2:
			if ok {
				output = append(output, v)
			} else {
				ch2 = nil
			}
		}
	}

	return output
}
