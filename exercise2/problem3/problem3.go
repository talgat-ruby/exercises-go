package main

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
		exp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			switch d {
			case "ul":
				exp[i][j] = i + j
			case "ll":
				exp[i][j] = n - i + j - 1
			case "ur":
				exp[i][j] = n + i - j - 1
			case "lr":
				exp[i][j] = (n - i - 1) + (n - j - 1)
			}
		}
	}
	return exp
}
