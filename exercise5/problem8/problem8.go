package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	select {
	case msq := <-ch:
		return msq
	case <-time.After(ttl):
		return "timeout"
	}
}
