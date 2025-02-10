package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	select {
	case msg, ok := <-ch:
		if !ok {
			return "fail"
		}
		return msg
	case <-time.After(ttl):
		return "timeout"
	}
}
