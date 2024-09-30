package main

func countVowels(str string) int {
	vowels := "aeiou"
	counter := 0
	for _, x := range str {
		for _, y := range vowels {
			if string(x) == string(y) {
				counter++
			}
		}
	}
	return counter
}
