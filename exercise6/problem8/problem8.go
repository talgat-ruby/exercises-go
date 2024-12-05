package problem8

import "sync"

func multiplex(chs []<-chan string) []string {
	var result []string
	var wg sync.WaitGroup
	cha := make(chan string)

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for value := range ch {
				cha <- value
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(cha)
	}()

	for val := range cha {
		result = append(result, val)
	}

	return result
}
