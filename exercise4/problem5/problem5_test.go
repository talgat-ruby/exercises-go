package problem5

import (
	"testing"
)

func TestSend(t *testing.T) {
	table := []struct {
		list []string
		exp  string
	}{
		{
			list: []string{
				"Hello",
				"dear",
				"friend!",
				"Learn",
				"from",
				"yesterday.",
				"Save",
				"our",
				"soles.",
			},
			exp: "Hello dear friend! Learn from yesterday. Save our soles.",
		},
		{
			list: []string{
				"Frankly,",
				"my",
				"dear,",
				"I",
				"don’t",
				"give",
				"a",
				"damn.",
			},
			exp: "Frankly, my dear, I don’t give a damn.",
		},
		{
			list: []string{
				"Houston,",
				"we",
				"have",
				"a",
				"problem.",
			},
			exp: "Houston, we have a problem.",
		},
	}

	for _, r := range table {
		out := send(r.list)
		if out != r.exp {
			t.Errorf("sum(%v) was incorrect, got: %v, expected: %v.", r.list, out, r.exp)
		}
	}
}
