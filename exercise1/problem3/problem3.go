package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, direction dir) [][]int {
	result := make([][]int, n) // Allocate memory for n slices

	for i := range result {
		result[i] = make([]int, n) // Allocate memory for n integers in each slice
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch direction {
			case ul:
				result[i][j] = i + j
			case ur:
				result[i][j] = n - j - 1 + i
			case ll:
				result[i][j] = n + j - 1 - i
			case lr:
				result[i][j] = n*2 - j - i - 2
			}
		}
	}
	return result
}
