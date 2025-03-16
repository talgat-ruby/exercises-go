package problem7

import (
	"math/rand"
	"sync"
	"time"
)

var (
	once     sync.Once
	initLock sync.Mutex
	initRun  int
)

func initConf() {
	initLock.Lock()
	defer initLock.Unlock()
	initRun++
}

func task() {
	once.Do(initConf)

	var t *time.Timer
	t = time.AfterFunc(
		randomDuration(), func() {
			t.Reset(randomDuration())
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
