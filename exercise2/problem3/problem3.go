package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(size int, direction dir) [][]int {

	array := make([][]int, size)

	for i := range array {
		array[i] = make([]int, size)
	}

	switch direction {
	case ul:
		for i := range array {
			for j := range array[i] {
				array[i][j] = i + j
			}
		}
		return array
	case ur:
		for i := range array {
			for j := range array[i] {
				array[i][j] = (size - 1 - j) + i
			}
		}
		return array
	case ll:
		for i := range array {
			for j := range array[i] {
				array[i][j] = (size - 1 - i) + j
			}
		}
		return array
	case lr:
		for i := range array {
			for j := range array[i] {
				array[i][j] = (size - 1 - i) + (size - 1 - j)
			}
		}
		return array
	default:
		for i := range array {
			for j := range array[i] {
				array[i][j] = 0
			}
		}
		return array
	}
}
