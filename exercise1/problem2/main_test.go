package main

import (
	"testing"
)

func TestBinary(t *testing.T) {
	table := []struct {
		num int
		exp string
	}{
		{100, "1100100"},
		{1, "1"},
		{0, "0"},
		{69, "1000101"},
		{1023, "1111111111"},
		{511, "111111111"},
		{666, "1010011010"},
		{123, "1111011"},
	}

	for _, r := range table {
		out := Binary(r.num)
		if out != r.exp {
			t.Errorf("binary(%d) was incorrect, expected: %s, got: %s.", r.num, r.exp, out)
		}
	}
}
