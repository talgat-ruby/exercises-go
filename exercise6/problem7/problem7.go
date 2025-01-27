package problem7

import (
	"fmt"
	"math/rand"
	"time"
)

func task() {
	start := time.Now()
	time.AfterFunc(
		randomDuration(), func() {
			fmt.Println(time.Now().Sub(start))
		},
	)
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
