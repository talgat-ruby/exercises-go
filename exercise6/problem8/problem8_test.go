package problem8

import (
	"slices"
	"testing"
	"time"
)

func TestMultiplex(t *testing.T) {
	t.Run(
		"single value", func(t *testing.T) {
			exp := []string{"one", "two", "three"}
			ch1 := make(chan string)
			ch2 := make(chan string)
			ch3 := make(chan string)

			go func() {
				defer close(ch1)
				time.Sleep(1 * time.Second)
				ch1 <- exp[0]
			}()

			go func() {
				defer close(ch2)
				time.Sleep(2 * time.Second)
				ch2 <- exp[1]
			}()

			go func() {
				defer close(ch3)
				time.Sleep(3 * time.Second)
				ch3 <- exp[2]
			}()

			out := multiplex([]<-chan string{ch1, ch2, ch3})

			if !slices.Equal(out, exp) {
				t.Errorf("multiplex() was incorrect, got: %v, expected: %v.", out, exp)
			}
		},
	)

	t.Run(
		"multiple values", func(t *testing.T) {
			exp := []string{"one", "two", "three", "four"}
			ch1 := make(chan string)
			ch2 := make(chan string)

			go func() {
				defer close(ch1)
				time.Sleep(1 * time.Second)
				ch1 <- exp[0]
				time.Sleep(500 * time.Millisecond)
				ch1 <- exp[1]
			}()

			go func() {
				defer close(ch2)
				time.Sleep(2 * time.Second)
				ch2 <- exp[2]
				ch2 <- exp[3]
			}()

			out := multiplex([]<-chan string{ch1, ch2})

			if !slices.Equal(out, exp) {
				t.Errorf("multiplex() was incorrect, got: %v, expected: %v.", out, exp)
			}
		},
	)
}
