package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	var result string
	select {
	case result = <-ch:
	case <-time.After(ttl):
		result = "timeout"
	}
	return result
}
