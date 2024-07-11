package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	timer := time.After(ttl)
	for {
		select {
		case <-timer:
			return "timeout"
		case x := <-ch:
			return x
		default:
		}
	}
}
