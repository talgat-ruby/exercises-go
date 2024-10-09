package main

import (
	"testing"
)

func TestPotatoes(t *testing.T) {
	table := []struct {
		crowd string
		exp   int
	}{
		{"potato", 1},
		{"potatopotatocherry", 2},
		{"potatopotatopotatoorange", 3},
		{"potatopotatobananapotatopotato", 4},
		{"potatopotatomangopotatopotatopotato", 5},
		{"potatocucumberpotatopotatopotatopotatopotato", 6},
	}

	for _, r := range table {
		out := potatoes(r.crowd)
		if out != r.exp {
			t.Errorf("potatoes(%s) was incorrect, expected: %d, got: %d.", r.crowd, r.exp, out)
		}
	}
}
