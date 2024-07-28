package main

func potatoes(s string) int {
	potato := "potato"
	sum := 0
	if len(s) < 6 {
		return 0
	}
	for i := 0; i <= len(s)-6; i++ {
		if string(s[i:i+6]) == potato {
			sum += 1
		}
	}
	return sum
}
