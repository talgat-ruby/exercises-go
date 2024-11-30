package main

func addUp(n int) int {
	if n <= 0 {
		return -1
	}
	// использовал сумму ариф. прогрес.
	return n * (n + 1) / 2
}
