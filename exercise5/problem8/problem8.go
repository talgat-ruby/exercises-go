package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {

	select {
	case message, ok := <-ch:
		if ok {
			return message
		}
	case <-time.After(ttl):
		break
	}
	return "timeout"
}
