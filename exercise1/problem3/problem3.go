package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	exp := make([][]int, n)
	for i := 0; i < n; i++ {
		arr := make([]int, n)
		exp[i] = arr
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			switch d {
			case ul:
				exp[i][j] = i + j
			case ur:
				exp[i][j] = i + (n - j - 1)
			case ll:
				exp[i][j] = (n - i - 1) + j
			case lr:
				exp[i][j] = (n - i - 1) + (n - j - 1)
			}
		}
	}
	return exp
}
