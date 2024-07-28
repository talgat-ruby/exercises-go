package main

func addUp(num int) int {
	var res int

	for i := 0; i <= num; i++ {
		res += i
	}

	return res
}
