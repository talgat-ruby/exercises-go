package main

func binary(num int) int {
	a := 0
	b := 0
	for num != 0{
		a = num %2
		b  = num /2
		num = b


	}
	return a
}
