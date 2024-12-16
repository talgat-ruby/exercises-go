package problem7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task() {
	var mu sync.Mutex
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(
		randomDuration(),
		func() {
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(time.Now().Sub(start))
			t.Reset(randomDuration())
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
