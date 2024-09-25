package main

import (
	"slices"
	"testing"

	"github.com/talgat-ruby/exercises-go/pkg/util"
)

func TestDiagonalize(t *testing.T) {
	util.SkipTestOptional(t)

	table := []struct {
		n   int
		d   dir
		exp [][]int
	}{
		{
			3,
			ul,
			[][]int{
				{0, 1, 2},
				{1, 2, 3},
				{2, 3, 4},
			},
		},
		{
			4,
			ur,
			[][]int{
				{3, 2, 1, 0},
				{4, 3, 2, 1},
				{5, 4, 3, 2},
				{6, 5, 4, 3},
			},
		},
		{
			3,
			ll,
			[][]int{
				{2, 3, 4},
				{1, 2, 3},
				{0, 1, 2},
			},
		},
		{
			5,
			lr,
			[][]int{
				{8, 7, 6, 5, 4},
				{7, 6, 5, 4, 3},
				{6, 5, 4, 3, 2},
				{5, 4, 3, 2, 1},
				{4, 3, 2, 1, 0},
			},
		},
	}

	for _, r := range table {
		matrix := diagonalize(r.n, r.d)
		if len(matrix) != len(r.exp) {
			t.Errorf("diagonalize(%d, %s) was incorrect, got: %v, expected: %v.", r.n, r.d, matrix, r.exp)
			return
		}
		for i, row := range matrix {
			if !slices.Equal(r.exp[i], row) {
				t.Errorf("diagonalize(%d, %s) was incorrect, got: %v, expected: %v.", r.n, r.d, matrix, r.exp)
				return
			}
		}
	}
}
