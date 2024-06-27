package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, corner dir) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	switch corner {
	case "ul":
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = i + j
			}
		}
	case "ur":
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - j) + i
			}
		}
	case "ll":
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - i) + j
			}
		}
	case "lr":
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = (n - 1 - i) + (n - 1 - j)
			}
		}
	}

	return matrix
}
