package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	select {
	case s := <-ch:
		return s
	case <-time.After(ttl):
		return "timeout"
	}
}
