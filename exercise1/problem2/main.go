package main

func binary(n int) string {
	if n == 0 {
		return "0"
	}

	res := ""
	base := "01" // base of binary
	a := n

	for a > 0 {
		res = string(base[a%len(base)]) + res
		a /= len(base)
	}
	return res
}
