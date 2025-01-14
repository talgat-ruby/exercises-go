package main

func potatoes(crowd string) int {
	var result int
	strL := len("potato")
	for i := 0; i <= len(crowd)-strL; i++ {
		if crowd[i:i+strL] == "potato" {
			result++
		}
	}
	return result
}
