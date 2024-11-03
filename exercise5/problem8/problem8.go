package problem8

import (
	"context"
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	ctx, cancel := context.WithTimeout(context.Background(), ttl)
	defer cancel()
	select {
	case <-ctx.Done():
		return "timeout"
	case result, ok := <-ch:
		if ok {
			return result
		} else {
			return "fail"
		}
	}
}
