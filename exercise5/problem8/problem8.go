package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {}
