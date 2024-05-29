package problem3

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, direction dir) [][]int {
	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	if direction == ul {
		implementUl(&matrix)
	} else if direction == ur {
		implementUr(&matrix)
	} else if direction == ll {
		implementLl(&matrix)
	} else if direction == lr {
		implementLr(&matrix)
	}
	return matrix
}

func implementUl(matrix *[][]int) {
	topLeft := 0
	for i := range *matrix {
		for j := range (*matrix)[i] {
			(*matrix)[i][j] = topLeft + j
		}
		topLeft++
	}
}

func implementUr(matrix *[][]int) {
	topLeft := len(*matrix) - 1
	for i := range *matrix {
		for j := range (*matrix)[i] {
			(*matrix)[i][j] = topLeft - j
		}
		topLeft++
	}
}

func implementLl(matrix *[][]int) {
	topLeft := len(*matrix) - 1
	for i := range *matrix {
		for j := range (*matrix)[i] {
			(*matrix)[i][j] = topLeft + j
		}
		topLeft--
	}
}

func implementLr(matrix *[][]int) {
	topLeft := (len(*matrix) - 1) * 2
	for i := range *matrix {
		for j := range (*matrix)[i] {
			(*matrix)[i][j] = topLeft - j
		}
		topLeft--
	}
}
