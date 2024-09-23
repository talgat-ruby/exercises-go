package main

import (
	"testing"
)

func TestAddUp(t *testing.T) {
	table := []struct {
		num int
		exp int
	}{
		{4, 10},
		{13, 91},
		{600, 180300},
		{392, 77028},
		{53, 1431},
		{897, 402753},
		{23, 276},
		{1000, 500500},
		{738, 272691},
		{100, 5050},
		{925, 428275},
		{1, 1},
		{999, 499500},
		{175, 15400},
		{111, 6216},
	}

	for _, r := range table {
		out := addUp(r.num)
		if out != r.exp {
			t.Errorf("addUp(%d) was incorrect, expected: %d, got: %d.", r.num, r.exp, out)
		}
	}
}
