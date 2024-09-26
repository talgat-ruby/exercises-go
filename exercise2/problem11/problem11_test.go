package problem11

import (
	"slices"
	"testing"
)

type row[T comparable] struct {
	inp []T
	exp []T
}

func TestRemoveDups(t *testing.T) {
	row1 := row[int]{
		[]int{1, 0, 1, 0},
		[]int{1, 0},
	}
	out1 := removeDups(row1.inp)
	if !slices.Equal(out1, row1.exp) {
		t.Errorf("removeDups(%v) was incorrect, got: %v, expected: %v.", row1.inp, out1, row1.exp)
	}

	row2 := row[bool]{
		[]bool{true, false, false, true},
		[]bool{true, false},
	}
	out2 := removeDups(row2.inp)
	if !slices.Equal(out2, row2.exp) {
		t.Errorf("removeDups(%v) was incorrect, got: %v, expected: %v.", row2.inp, out2, row2.exp)
	}

	row3 := row[string]{
		[]string{"John", "Taylor", "John"},
		[]string{"John", "Taylor"},
	}
	out3 := removeDups(row3.inp)
	if !slices.Equal(out3, row3.exp) {
		t.Errorf("removeDups(%v) was incorrect, got: %v, expected: %v.", row3.inp, out3, row3.exp)
	}
}
