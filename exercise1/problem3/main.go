package main

func numberSquares(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += (n - i + 1) * (n - i + 1)
	}

	return res
}
