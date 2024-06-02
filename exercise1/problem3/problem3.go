package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(size int, direction dir) [][]int {
	matrix := createMatrix(size)
	for i := 0; i < size; i++ {
		switch direction {
		case ul:
			fillFields(matrix[i], i, true)
		case ur:
			fillFields(matrix[i], size+i-1, false)
		case ll:
			fillFields(matrix[i], size-i-1, true)
		case lr:
			fillFields(matrix[i], (size-1)*2-i, false)
		}
	}
	return matrix
}

func createMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	return matrix
}

func fillFields(slice []int, first int, increment bool) {
	for i := 0; i < len(slice); i++ {
		slice[i] = first

		if increment {
			first++
		} else {
			first--
		}
	}
}
