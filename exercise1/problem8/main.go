package main

func countVowels(word string) (count int) {
	for _, r := range word {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return
}
