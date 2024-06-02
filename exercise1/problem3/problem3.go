package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(size int, direction dir) [][]int {
	matrix := make([][]int, size)
	if size == 0 {
		return matrix
	}
	dirRunes := []rune(direction)
	isToUp := 0
	isToLeft := 0
	if dirRunes[0] == 'l' {
		isToUp = 1
	}
	if dirRunes[1] == 'r' {
		isToLeft = 1
	}
	for i := 0; i < size; i++ {
		correctedY := isToUp*(size-1) + (-2*isToUp+1)*i
		matrix[correctedY] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[correctedY][j] = isToLeft*(size+2*i-1) + (-2*isToLeft+1)*(i+j)
		}
	}
	return matrix
}
