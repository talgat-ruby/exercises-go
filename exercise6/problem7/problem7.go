package problem7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task() {
	start := time.Now()
	var m sync.Mutex
	var t *time.Timer

	resetTimer := func() {
		m.Lock()
		defer m.Unlock()
		if t != nil {
			t.Reset(randomDuration())
		}
	}

	m.Lock() 
	t = time.AfterFunc(
		randomDuration(),
		func() {
			fmt.Println(time.Since(start))
			resetTimer() 
		},
	)
	m.Unlock() 

	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
