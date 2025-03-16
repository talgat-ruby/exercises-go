package problem7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var once sync.Once

func task() {
	start := time.Now()

	once.Do(func() {
		t := time.AfterFunc(randomDuration(), func() {
			fmt.Println("Duration:", time.Now().Sub(start))
		})
		t.Reset(randomDuration())
	})
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
