package main

import (
	"fmt"
)

// diagonalize creates an n x n matrix with numbers ordered diagonally based on the specified corner
func diagonalize(n int, corner string) [][]int {
	// Initialize an empty n x n matrix
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Fill the matrix based on the specified corner
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch corner {
			case "ul": // Upper-left corner
				matrix[i][j] = i + j
			case "ur": // Upper-right corner
				matrix[i][j] = (n - 1 - i) + j
			case "ll": // Lower-left corner
				matrix[i][j] = i + (n - 1 - j)
			case "lr": // Lower-right corner
				matrix[i][j] = (n - 1 - i) + (n - 1 - j)
			}
		}
	}

	return matrix
}

func main() {
	fmt.Println(diagonalize(3, "ul")) // [[0, 1, 2], [1, 2, 3], [2, 3, 4]]
	fmt.Println(diagonalize(4, "ur")) // [[3, 2, 1, 0], [4, 3, 2, 1], [5, 4, 3, 2], [6, 5, 4, 3]]
	fmt.Println(diagonalize(3, "ll")) // [[2, 3, 4], [1, 2, 3], [0, 1, 2]]
	fmt.Println(diagonalize(5, "lr")) // [[8, 7, 6, 5, 4], [7, 6, 5, 4, 3], [6, 5, 4, 3, 2], [5, 4, 3, 2, 1], [4, 3, 2, 1, 0]]
}
