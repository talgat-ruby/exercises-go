package main

func countVowels(input string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		if input[i] == 'a' || input[i] == 'e' || input[i] == 'i' || input[i] == 'o' || input[i] == 'u' {
			count++
		}
	}
	return count
}
