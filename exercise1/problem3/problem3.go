package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	mat := make([][]int, n)
	if d == ul {
		for i := 0; i < n; i++ {
			mat[i] = make([]int, n)
			for j := 0; j < n; j++ {
				mat[i][j] = i + j
			}
		}
	} else if d == ur {
		for i := 0; i < n; i++ {
			mat[i] = make([]int, n)
			curVal := 0
			for j := n - 1; j >= 0; j-- {
				mat[i][j] = curVal + i
				curVal++
			}
		}
	} else if d == ll {
		curVal := 0
		for i := n - 1; i >= 0; i-- {
			mat[i] = make([]int, n)
			for j := 0; j < n; j++ {
				mat[i][j] = curVal + j
			}
			curVal++
		}
	} else if d == lr {
		for i := n - 1; i >= 0; i-- {
			mat[i] = make([]int, n)
			curVal := n - i - 1
			for j := n - 1; j >= 0; j-- {
				mat[i][j] = curVal
				curVal++
			}
		}
	}

	return mat
}
