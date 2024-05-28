package problem4

import (
	"maps"
	"testing"
)

func TestMapping(t *testing.T) {
	table := []struct {
		inp []string
		exp map[string]string
	}{
		{
			[]string{"a", "b", "c"},
			map[string]string{"a": "A", "b": "B", "c": "C"},
		},
		{
			[]string{"p", "s", "t"},
			map[string]string{"p": "P", "s": "S", "t": "T"},
		},
		{
			[]string{"a", "v", "y", "z"},
			map[string]string{"a": "A", "v": "V", "y": "Y", "z": "Z"},
		},
	}

	for _, r := range table {
		out := mapping(r.inp)
		if !maps.Equal(out, r.exp) {
			t.Errorf("mapping(%v) was incorrect, got: %v, expected: %v.", r.inp, out, r.exp)
		}
	}
}
