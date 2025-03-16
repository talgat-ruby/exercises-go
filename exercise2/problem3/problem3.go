package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(d int, direction dir) [][]int {
	matrix := make([][]int, d)
	for i := 0; i < d; i++ {
		matrix[i] = make([]int, d)
		for j := 0; j < d; j++ {
			switch {
			case direction == ul:
				matrix[i][j] = j + i
			case direction == lr:
				matrix[i][j] = (d-1)*2 - i - j
			case direction == ur:
				matrix[i][j] = (d - 1) + i - j
			case direction == ll:
				matrix[i][j] = (d - 1) + j - i
			}
		}
	}
	return matrix
}
