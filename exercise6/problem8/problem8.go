package problem8

import "sync"

func multiplex(chs []<-chan string) []string {
	var result []string
	var wg sync.WaitGroup
	var mux sync.Mutex 

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch { 
				mux.Lock()
				result = append(result, msg) 
				mux.Unlock()
			}
		}(ch)
	}

	wg.Wait() 
	return result
}
