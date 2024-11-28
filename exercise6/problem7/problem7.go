package problem7

import (
	"fmt"
	"math/rand"
	"time"
)

func task() {
	start := time.Now()

	done := make(chan struct{})

	var t *time.Timer

	timerFunc := func() {
		fmt.Println(time.Now().Sub(start))
		select {
		case <-done:
			return
		default:

			t.Reset(randomDuration())
		}
	}

	t = time.AfterFunc(randomDuration(), timerFunc)

	time.Sleep(5 * time.Second)

	t.Stop()
	close(done)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
