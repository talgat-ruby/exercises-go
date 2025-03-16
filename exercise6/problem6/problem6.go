package problem6

import (
	"sync"
)

func runTasks(init func()) {
	var once sync.Once

	for range 10 {
		once.Do(func() {
			init()
		})
	}

}
