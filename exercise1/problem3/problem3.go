package problem3

import "math"

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	matrix := make([][]int, n)
	colAdder, rowAdder := 0, 0

	switch d {
	case ur:
		colAdder = n - 1
	case ll:
		rowAdder = n - 1
	case lr:
		colAdder = n - 1
		rowAdder = n - 1
	}

	for r := 0; r < n; r++ {
		matrix[r] = make([]int, n)
		for c := 0; c < n; c++ {
			matrix[r][c] = int(math.Abs(float64(colAdder-c))) +
				int(math.Abs(float64(rowAdder-r)))
		}
	}

	return matrix
}
