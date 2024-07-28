package main

func numberSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return
}
