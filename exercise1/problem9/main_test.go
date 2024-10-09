package main

import (
	"testing"
)

func TestBitwiseOperator(t *testing.T) {
	table := []struct {
		a, b                  int
		expAND, expOR, expXOR int
	}{
		{a: 7, b: 12, expAND: 4, expOR: 15, expXOR: 11},
		{a: 32, b: 17, expAND: 0, expOR: 49, expXOR: 49},
		{a: 13, b: 19, expAND: 1, expOR: 31, expXOR: 30},
	}

	for _, r := range table {
		outAND := bitwiseAND(r.a, r.b)
		outOR := bitwiseOR(r.a, r.b)
		outXOR := bitwiseXOR(r.a, r.b)

		if outAND != r.expAND {
			t.Errorf("bitwiseAND(%d, %d) was incorrect, expected: %d, got: %d.", r.a, r.b, r.expAND, outAND)
		}
		if outOR != r.expOR {
			t.Errorf("bitwiseOR(%d, %d) was incorrect, expected: %d, got: %d.", r.a, r.b, r.expOR, outOR)
		}
		if outXOR != r.expXOR {
			t.Errorf("bitwiseXOR(%d, %d) was incorrect, expected: %d, got: %d.", r.a, r.b, r.expXOR, outXOR)
		}
	}
}
