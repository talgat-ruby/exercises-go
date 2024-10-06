package main

func countVowels(str string) int {
	var count int
	for _, ch := range str {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
		}
	}

	return count
}
