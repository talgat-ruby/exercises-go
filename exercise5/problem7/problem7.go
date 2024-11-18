package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string

	for {
		select {
		case val, ok := <-ch1:
			if ok {
				result = append(result, val)
				break
			}
			ch1 = nil
		case val, ok := <-ch2:
			if ok {
				result = append(result, val)
				break
			}
			ch2 = nil
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}

	return result
}
