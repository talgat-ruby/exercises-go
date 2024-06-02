package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(size int, direction dir) [][]int {
	matrix := make([][]int, 0)

	for i := 0; i < size; i++ {
		row := make([]int, 0)
		for j := 0; j < size; j++ {
			row = append(row, calcCellValue(size, i, j, direction))
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func calcCellValue(size int, i int, j int, direction dir) int {
	switch direction {
	case ul:
		return i + j
	case ur:
		return size + i - j - 1
	case ll:
		return size + j - i - 1
	case lr:
		return 2*size - i - j - 2
	default:
		panic("unrecognized value")
	}
}
