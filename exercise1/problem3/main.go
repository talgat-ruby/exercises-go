package main

func numberSquares(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total = total + (n-i+1)*(n-i+1)
	}
	return total
}
func main() {
}
