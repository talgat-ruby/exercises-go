package main

func addUp(n int) int {
	summa := 0
	for i := 1; i <= n; i++ {
		summa += i
	}
	return summa
}
