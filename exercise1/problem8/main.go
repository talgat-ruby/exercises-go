package main

func countVowels(s string) int {
	k := 0
	for _, r := range s {
		if r == 'a' || r == 'e' || r == 'o' || r == 'i' || r == 'u' || r == 'y' {
			k++
		}
	}
	return k
}
