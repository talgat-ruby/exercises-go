package main

func addUp(num int) int {
	res := 0
	positive := true
	if num<0 {
		positive = false
		num*=-1
	}
	for i:=0;i<num;i++{
		res+=i+1
	}
	if !positive{
		return -res
	}
	return res
}
