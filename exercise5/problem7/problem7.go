package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string
	var done1 bool
	var done2 bool
	for {
		select {
		case val1, ok := <-ch1:
			if !ok {
				done1 = true
			} else {
				result = append(result, val1)
			}
		case val2, ok := <-ch2:
			if !ok {
				done2 = true
			} else {
				result = append(result, val2)
			}
		}
		if done1 && done2 {
			break
		}

	}
	return result
}
