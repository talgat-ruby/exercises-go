package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		arr := make([]int, n)
		res[i] = arr
	}
	switch d {
	case ul:
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				res[i][j] = i + j
			}
		}
	case ur:
		for i := 0; i < n; i++ {
			for j := n - 1; j >= 0; j-- {
				res[i][j] = n - 1 - j + i
			}
		}
	case ll:
		for i := n - 1; i >= 0; i-- {
			for j := 0; j < n; j++ {
				res[i][j] = n - 1 - i + j
			}
		}
	case lr:
		for i := n - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- {
				res[i][j] = 2*(n-1) - j - i
			}
		}
	}
	return res

}
