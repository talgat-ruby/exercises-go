package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				return "channel closed"
			}
			return msg
		case <-time.After(ttl):
			return "timeout"
		}
	}
}
