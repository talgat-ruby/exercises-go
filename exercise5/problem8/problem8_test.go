package problem8

import (
	"testing"
	"time"
)

func TestWithTimeout(t *testing.T) {
	table := []struct {
		message  string
		duration time.Duration
		ttl      time.Duration
		exp      string
	}{
		{
			message:  "success",
			duration: time.Second,
			ttl:      2 * time.Second,
			exp:      "success",
		},
		{
			message:  "success",
			duration: 3 * time.Second,
			ttl:      2 * time.Second,
			exp:      "timeout",
		},
		{
			message:  "fail",
			duration: 3 * time.Second,
			ttl:      5 * time.Second,
			exp:      "fail",
		},
	}

	for _, r := range table {
		ch := make(chan string)
		go func() {
			time.Sleep(r.duration)
			ch <- r.message
		}()
		out := withTimeout(ch, r.ttl)

		if out != r.exp {
			t.Errorf("withTimeout(ch, %v) was incorrect, got: %s, expected: %s.", r.ttl, out, r.exp)
		}
	}
}
