package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	output := make([]string, 0)

	for {
		v, more := <-ch1
		if more {
			output = append(output, v)
		} else {
			break
		}
	}

	for {
		v, more := <-ch2
		if more {
			output = append(output, v)
		} else {
			break
		}
	}

	return output
}
