package main

func addUp(a int) int {
	res := 0

	for i:=1; i<=a; i++ {
		res +=i
	}
	return res
}
