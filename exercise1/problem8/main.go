package main

func countVowels(a string) int {
	res := 0

	for _, i := range a {
		if vowels(i) {
			res++
		}
	}

	return res 
}

func vowels(a rune) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' || a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'
}