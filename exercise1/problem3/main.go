package main

func numberSquares(n int) int {
	if n <= 0 {
		return -1
	}

	res := 0
	for k := 1; k <= n; k++ {
		res += k * k
	}
	return res
}
