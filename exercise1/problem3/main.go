package main

func numberSquares(num int) int {
	positive := true
	if num<0 {
		positive = false
		num*=-1
	}
	res := 0
	for i:=1; i<=num;i++{
		res += i*i
	}
	if !positive{
		return -res
	}
	return res
}
