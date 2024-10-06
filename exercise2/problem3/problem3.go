package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, direction dir) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	if direction == "ul" {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = i + j
			}
		}
	}
	if direction == "ur" {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - j) + i
			}
		}
	}
	if direction == "ll" {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - i) + j
			}
		}
	}
	if direction == "lr" {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - i) + (n - 1 - j)
			}
		}
	}

	return matrix
}
