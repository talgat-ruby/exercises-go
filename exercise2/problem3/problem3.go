package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, direction dir) [][]int {
	if n <= 0 {
		return [][]int{}
	}
	if !(direction == ul || direction == ur || direction == ll || direction == lr) {
		return [][]int{}
	}
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	if direction == ul {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = i + j
			}
		}
	}
	if direction == ur {
		for i := 0; i < n; i++ {
			for j := n - 1; j >= 0; j-- {
				matrix[i][j] = i + n - 1 - j
			}
		}
	}
	if direction == ll {
		for i := n - 1; i >= 0; i-- {
			for j := 0; j < n; j++ {
				matrix[i][j] = n - 1 - i + j
			}
		}
	}
	if direction == lr {
		for i := n - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				matrix[i][j] = 2*(n-1) - (i + j)
			}
		}
	}
	return matrix
}
