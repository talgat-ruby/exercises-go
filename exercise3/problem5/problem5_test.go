package problem5

import (
	"testing"
)

func TestPerson(t *testing.T) {
	p1 := &Person{"Samuel", 24}
	p2 := &Person{"Joel", 36}
	p3 := &Person{"Lily", 24}

	table := []struct {
		out string
		exp string
	}{
		{
			p1.compareAge(p2),
			"Joel is older than me.",
		},
		{
			p1.compareAge(p3),
			"Lily is the same age as me.",
		},
		{
			p2.compareAge(p1),
			"Samuel is younger than me.",
		},
		{
			p2.compareAge(p3),
			"Lily is younger than me.",
		},
		{
			p3.compareAge(p1),
			"Samuel is the same age as me.",
		},
		{
			p3.compareAge(p2),
			"Joel is older than me.",
		},
	}

	for _, r := range table {
		if r.out != r.exp {
			t.Errorf("Compare() was incorrect, got len: %s, expected len: %s", r.out, r.exp)
		}
	}
}
