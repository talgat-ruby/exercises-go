package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	select {
	case message := <-ch:
		return message
	case <-time.After(ttl):
		return "timeout"
	}
}
