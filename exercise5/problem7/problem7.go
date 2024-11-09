package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string
	for v := range ch1 {
		result = append(result, v)
	}

	for v := range ch2 {
		result = append(result, v)
	}
	return result
}
