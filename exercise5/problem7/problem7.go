package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	out := make([]string, 0)

	for ch1 != nil || ch2 != nil {
		select {
		case val, ok := <-ch1:
			if ok {
				out = append(out, val)
			} else {
				ch1 = nil
			}
		case val, ok := <-ch2:
			if ok {
				out = append(out, val)
			} else {
				ch2 = nil
			}
		}
	}

	return out
}
