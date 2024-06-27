package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			switch d {
			case ul:
				matrix[i][j] = i + j
			case ur:
				matrix[i][j] = n + i - 1 - j
			case ll:
				matrix[i][j] = n - 1 + j - i
			case lr:
				matrix[i][j] = (n-1)*2 - i - j
			}

		}
	}

	return matrix
}
