package main

func potatoes(a string) int {
	res := 0
	str := ""
	for i := 0; i<len(a); i++ {
		str += string(a[i])
		if str[0] != 'p' {
			str = ""
		}
		if str == "potato" {
			res++
			str = ""
		}
		
	}
	return res 
}
