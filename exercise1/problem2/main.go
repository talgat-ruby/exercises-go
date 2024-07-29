package main

func binary(n int) string{
	if n == 0 {
		return "0"
	}
	binaryStr := ""
	for n > 0 {
		bit := n % 2
		if bit==0{
			binaryStr = "0"+binaryStr
		} else {
			binaryStr = "1"+binaryStr
		}
		n = n / 2
	}
	return binaryStr
}
