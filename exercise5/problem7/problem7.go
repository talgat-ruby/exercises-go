package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {

	var result []string

	for {
		select {
		case value1, ok1 := <-ch1:
			if ok1 {
				result = append(result, value1)
			} else {
				ch1 = nil
			}
		case value2, ok2 := <-ch2:
			if ok2 {
				result = append(result, value2)
			} else {
				ch2 = nil
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
	return result
}
