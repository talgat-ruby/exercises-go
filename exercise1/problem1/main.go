package main

func addUp(num int) int {
	res := 0
	for i := 1; i <= num; i++ {
		res = res + i
	}
	return res
}
