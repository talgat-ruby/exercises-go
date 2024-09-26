package main

func diagonalize(n int, corner string) [][]int {

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	switch corner {
	case "ul":
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				matrix[r][c] = r + c
			}
		}
	case "ur":
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				matrix[r][c] = (n - 1 + r) - c
			}
		}
	case "ll":
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				matrix[r][c] = (n - 1 - r) + c
			}
		}
	case "lr":
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				matrix[r][c] = (n - 1 - c) + (n - 1 - r)
			}
		}
	}

	return matrix
}
