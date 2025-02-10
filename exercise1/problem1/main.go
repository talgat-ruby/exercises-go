package main

func addUp(i int) int {
	var sum int = 0
	for j := 1; j <= i; j++ {
		sum += int(j)
	}

	return sum
}
