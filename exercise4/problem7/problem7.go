package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	output := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		select {
		case x := <-ch1:
			output = append(output, x)
		case x := <-ch2:
			output = append(output, x)
		}
	}
	return output
}
