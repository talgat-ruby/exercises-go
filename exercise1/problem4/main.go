package main

func detectWord(crowd string) string {
	word := ""
	for _, char := range crowd {
		if char >= 'a' && char <= 'z' {
			word += string(char)
		}
	}
	return word
}
