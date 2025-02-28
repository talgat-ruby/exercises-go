package main

import "fmt"

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	switch d {
	case ul:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = i + j
			}
		}
	case ur:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - j) + i
			}
		}
	case ll:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - i) + j
			}
		}
	case lr:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - i) + (n - 1 - j)
			}
		}
	}

	return matrix
}

func main() {

	n := 3
	result := diagonalize(n, ul)
	for _, row := range result {
		fmt.Println(row)
	}

	fmt.Println()

	result = diagonalize(4, ur)
	for _, row := range result {
		fmt.Println(row)
	}
}
