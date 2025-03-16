package main

func numberSquares(a int) int {
	res := 0

	for i:=1; i<=a; i++ {
		b := a - i + 1
		b = b * b 
		res+= b
	}

	return res
}
