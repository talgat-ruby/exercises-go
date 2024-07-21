package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	//var output []string
	//var wg sync.WaitGroup
	//
	//// Function to collect values from a channel until closed
	//collect := func(ch <-chan string) {
	//	defer wg.Done()
	//	for msg := range ch {
	//		output = append(output, msg)
	//	}
	//}
	//
	//// Increment the WaitGroup for each goroutine
	//wg.Add(2)
	//
	//// Start goroutines to collect from ch1 and ch2 concurrently
	//go collect(ch1)
	//go collect(ch2)
	//
	//// Wait for both goroutines to complete
	//wg.Wait()
	//
	//return output
}
