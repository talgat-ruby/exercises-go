package main

import (
	"testing"
)

func TestEmojify(t *testing.T) {
	table := []struct {
		text string
		exp  string
	}{
		{"Make me smile", "Make me 🙂"},
		{"Make me grin", "Make me 😀"},
		{"Make me sad", "Make me 😥"},
		{"Make me mad", "Make me 😠"},
		{"I feel sad and mad", "I feel 😥 and 😠"},
		{"smile, grin, sad, mad", "🙂, 😀, 😥, 😠"},
	}

	for _, r := range table {
		out := emojify(r.text)
		if out != r.exp {
			t.Errorf("emojify(%s) was incorrect, expected: %s, got: %s.", r.text, r.exp, out)
		}
	}
}
