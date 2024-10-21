package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(number int, diagonal dir) [][]int {
	matrix := make([][]int, number)
	for i := 0; i < number; i++ {
		matrix[i] = make([]int, number)
		for j := 0; j < number; j++ {
			if diagonal == ul {
				matrix[i][j] = i + j
			} else if diagonal == ur {
				matrix[i][j] = number + i - j - 1
			} else if diagonal == ll {
				matrix[i][j] = number - i + j - 1
			} else {
				matrix[i][j] = 2*number - i - j - 2
			}
		}
	}
	return matrix
}
