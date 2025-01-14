package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	matrix := make([][]int, n)
	switch d {
	case ul:
		count := 0
		getMatrix(count, matrix, n, true, true)
		break
	case ll:
		count := n - 1
		getMatrix(count, matrix, n, true, false)
		break
	case ur:
		count := n - 1
		getMatrix(count, matrix, n, false, true)
		break
	case lr:
		count := (n - 1) * 2
		getMatrix(count, matrix, n, false, false)
		break
	}
	return matrix
}

func getMatrix(count int, matrix [][]int, n int, o bool, countChecker bool) {
	for i := range n {
		matrix[i] = make([]int, n)
		for j := range n {
			if o {
				matrix[i][j] = count + j
			} else {
				matrix[i][j] = count - j
			}
		}
		if countChecker {
			count++
		} else {
			count--
		}
	}
}
