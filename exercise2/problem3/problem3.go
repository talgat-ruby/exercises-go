package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(size int, d dir) [][]int {
	out := make([][]int, size)
	for i := 0; i < size; i++ {
		out[i] = make([]int, size)
		for j := 0; j < size; j++ {
			if d == ul {
				out[i][j] = i + j
			} else if d == ur {
				out[i][j] = size + i - j - 1
			} else if d == ll {
				out[i][j] = size - i + j - 1
			} else {
				out[i][j] = 2*size - i - j - 2
			}
		}
	}
	return out
}
