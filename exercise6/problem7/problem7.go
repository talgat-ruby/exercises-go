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
	m := sync.Mutex{}

	resetTimer := func() {
		m.Lock() // Lock before accessing the timer
		defer m.Unlock()

		// Reset the timer inside the critical section
		t.Reset(randomDuration())
	}

	t = time.AfterFunc(
		randomDuration(),
		func() {
			m.Lock() // Lock before accessing the timer
			defer m.Unlock()
			fmt.Println(time.Now().Sub(start))
			resetTimer()
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
