package problem6

import (
	"sync/atomic"
	"testing"
)

func initGen(counter *atomic.Int64) func() {
	return func() {
		counter.Add(1)
	}
}

func TestInitFunc(t *testing.T) {
	t.Run(
		"init should run one only once", func(t *testing.T) {
			var exp int64 = 1

			var counter atomic.Int64

			init := initGen(&counter)
			runTasks(init)

			out := counter.Load()

			if exp != out {
				t.Errorf("init function run times wrong, expected: %v, got: %v.", exp, out)
			}
		},
	)
}
