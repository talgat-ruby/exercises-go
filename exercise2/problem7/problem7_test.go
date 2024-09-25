package problem7

import (
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	outA, outB := b, a
	swap(&a, &b)
	if a != outA && b != outB {
		t.Errorf("swap was incorrect, got: %d, %d expected: %d, %d.", a, b, outA, outB)
	}
}
