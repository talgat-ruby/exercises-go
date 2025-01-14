package problem1

import (
	"testing"
)

func TestIncrementConcurrently(t *testing.T) {
	inp := 2
	exp := 3
	out := incrementConcurrently(inp)
	if out != exp {
		t.Errorf("incrementConcurrently(%d) was incorrect, got: %d, expected: %d.", inp, out, exp)
	}
}
