package main

func potatoes(s string) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		if i+len("potato") <= len(s) && s[i:i+len("potato")] == "potato" {
			counter++
		}
	}
	return counter
}
