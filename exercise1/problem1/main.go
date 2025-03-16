package main

func addUp(n int) int {
	if n == 1 {
		return 1
	} else {
		return addUp(n-1) + n
	}

}
