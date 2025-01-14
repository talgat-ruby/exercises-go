package main

func countVowels(text string) int {
	var count int
	for _, v := range text {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			count++
		}
	}
	return count
}
