package problem3

import "sync/atomic"

type counter struct {
	val int64
}

func (c *counter) inc() {
	atomic.AddInt64(&c.val, 1)
}

func (c *counter) dec() {
	atomic.AddInt64(&c.val, -1)
}

func (c *counter) value() int64 {

	return atomic.LoadInt64(&c.val)
}

func newCounter() *counter {
	return &counter{
		val: 0,
	}
}
