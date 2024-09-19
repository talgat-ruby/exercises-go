package problem8

import (
	"maps"
	"reflect"
	"testing"
)

func TestSimplify(t *testing.T) {
	table := []struct {
		inp []string
		exp map[string]int
	}{
		{
			[]string{"a", "b", "c"},
			map[string]int{"a": 0, "b": 1, "c": 2},
		},
		{
			[]string{"z", "y", "x", "u", "v"},
			map[string]int{"z": 0, "y": 1, "x": 2, "u": 3, "v": 4},
		},
	}

	for _, r := range table {
		out := simplify(r.inp)
		if !maps.Equal(out, r.exp) {
			t.Errorf("simplify(%v) was incorrect, got: %v, expected: %v.", r.inp, out, r.exp)
		}
	}

	tp := reflect.TypeOf(load)
	if tp.Kind() != reflect.Func || tp.NumOut() != 0 || tp.NumIn() != 2 || tp.In(0).Kind() == reflect.Pointer || tp.In(1).Kind() == reflect.Pointer {
		t.Errorf("load wrong format")
	}
}
