package problem7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//TODO: identify the data race
// fix the issue.

var (
	once     sync.Once // Ensures init runs only once
	initLock sync.Mutex
	initRun  int // To count how many times init actually runs
)

func initConf() {
	initLock.Lock()
	defer initLock.Unlock()
	fmt.Println("Initializing configuration...")
	initRun++
}

func task() {
	// Ensure `initConf` is called only once
	once.Do(initConf)

	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(
		randomDuration(), func() {
			fmt.Println("Elapsed:", time.Now().Sub(start))
			t.Reset(randomDuration())
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
