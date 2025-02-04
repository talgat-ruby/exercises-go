package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	for {
		select {
		case v := <-ch:
			return v
		case <-time.After(ttl):
			return "timeout"
		}
	}
}
