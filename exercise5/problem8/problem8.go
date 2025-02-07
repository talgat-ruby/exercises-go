package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	timer := time.NewTimer(ttl)
	defer timer.Stop()

	select {
	case msg := <-ch:
		return msg
	case <-timer.C:
		return "timeout"
	}
}
