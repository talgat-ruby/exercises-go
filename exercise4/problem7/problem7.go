package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	result := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		select {
		case v := <-ch1:
			result = append(result, v)
		case v := <-ch2:
			result = append(result, v)
		}
	}
	return result
}
