package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	m := make([][]int, n)

	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch d {
			case ul:
				m[i][j] = i + j
			case ur:
				m[i][j] = n - 1 - j + i
			case ll:
				m[i][j] = n - 1 - i + j
			case lr:
				m[i][j] = 2*n - 2 - i - j
			}
		}
	}

	return m
}
