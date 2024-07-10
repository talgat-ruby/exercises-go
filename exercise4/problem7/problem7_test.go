package problem7

import (
	"slices"
	"testing"
	"time"
)

func TestMultiplex(t *testing.T) {
	exp := []string{"one", "two"}
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- exp[0]
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- exp[1]
	}()

	out := multiplex(ch1, ch2)

	if !slices.Equal(out, exp) {
		t.Errorf("multiplex() was incorrect, got: %v, expected: %v.", out, exp)
	}
}
