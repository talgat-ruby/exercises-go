package main

func potatoes(s string) int {

	n := 0
	word := "potato"
	wordLen := len(word)
	for i := 0; i <= len(s)-wordLen; i++ {
		if s[i:i+wordLen] == word {
			n++
		}
	}
	return n

}
