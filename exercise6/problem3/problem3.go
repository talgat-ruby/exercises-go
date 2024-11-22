package problem3

import "sync/atomic"

type counter struct {
	val int64
}

func newCounter() *counter {
	return &counter{
		val: 0,
	}
}

func (c *counter) Increment() {
	atomic.AddInt64(&c.val, 1)
}

func (c *counter) Decrement() {
	atomic.AddInt64(&c.val, -1)
}

func (c *counter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}
