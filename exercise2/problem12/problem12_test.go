package problem11

import (
	"slices"
	"testing"
)

type row[K comparable, V any] struct {
	inp  map[K]V
	exp1 []K
	exp2 []V
}

func TestKeysAndValues(t *testing.T) {
	row1 := row[string, int]{
		map[string]int{"a": 1, "b": 2, "c": 3},
		[]string{"a", "b", "c"},
		[]int{1, 2, 3},
	}
	out1_1, out1_2 := keysAndValues(row1.inp)
	if !slices.Equal(out1_1, row1.exp1) || !slices.Equal(out1_2, row1.exp2) {
		t.Errorf("keysAndValues(%v) was incorrect, got: %v, %v, expected: %v, %v.", row1.inp, out1_1, out1_2, row1.exp1, row1.exp2)
	}

	row2 := row[string, string]{
		map[string]string{"a": "Apple", "b": "Microsoft", "c": "Google"},
		[]string{"a", "b", "c"},
		[]string{"Apple", "Microsoft", "Google"},
	}
	out2_1, out2_2 := keysAndValues(row2.inp)
	if !slices.Equal(out2_1, row2.exp1) || !slices.Equal(out2_2, row2.exp2) {
		t.Errorf("keysAndValues(%v) was incorrect, got: %v, %v, expected: %v, %v.", row2.inp, out2_1, out2_2, row2.exp1, row2.exp2)
	}

	row3 := row[int, bool]{
		map[int]bool{1: true, 2: false, 3: false},
		[]int{1, 2, 3},
		[]bool{true, false, false},
	}
	out3_1, out3_2 := keysAndValues(row3.inp)
	if !slices.Equal(out3_1, row3.exp1) || !slices.Equal(out3_2, row3.exp2) {
		t.Errorf("keysAndValues(%v) was incorrect, got: %v, %v, expected: %v, %v.", row3.inp, out3_1, out3_2, row3.exp1, row3.exp2)
	}

	row4 := row[int, bool]{
		nil,
		[]int{},
		[]bool{},
	}
	out4_1, out4_2 := keysAndValues(row4.inp)
	if !slices.Equal(out4_1, row4.exp1) || !slices.Equal(out4_2, row4.exp2) {
		t.Errorf("keysAndValues(%v) was incorrect, got: %v, %v, expected: %v, %v.", row4.inp, out4_1, out4_2, row4.exp1, row4.exp2)
	}
}
