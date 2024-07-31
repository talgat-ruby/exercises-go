package main

func detectWord(s string) string {
	word := ""

	for _, w := range s {
		if w >= 97 && w <= 122 {
			word += string(rune(w))
		}
	}
	return word

}
