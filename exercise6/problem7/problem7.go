package problem7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func task() {
	start := time.Now()
	var t *time.Timer
	mx := &sync.Mutex{}

	t = time.AfterFunc(
		randomDuration(),
		func() {
			defer mx.Unlock()
			mx.Lock()
			fmt.Println(time.Now().Sub(start))

			func() {
				defer mx.Unlock()
				mx.Lock()
				t.Reset(randomDuration())
			}()
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
