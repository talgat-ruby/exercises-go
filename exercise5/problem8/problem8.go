package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	select {
	case val := <-ch:
		return val
	case <-time.After(ttl):
		return "timeout"
	}
}
