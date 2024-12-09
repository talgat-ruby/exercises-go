package problem3

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run(
		"concurrently increment & decrement counter", func(t *testing.T) {
			cntr := newCounter()
			var wg sync.WaitGroup

			inc := 200
			wg.Add(inc)
			go func() {
				for range inc {
					go func() {
						defer wg.Done()
						cntr.inc()
					}()
				}
			}()

			dec := 100
			wg.Add(dec)
			go func() {
				for range dec {
					go func() {
						defer wg.Done()
						cntr.dec()
					}()
				}
			}()

			wg.Wait()

			out := int64(inc - dec)
			if cntr.value() != out {
				t.Errorf("counter value was incorrect, expected: %d, got: %d.", out, cntr.value())
			}
		},
	)
}
