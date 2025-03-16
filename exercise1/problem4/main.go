package main

func detectWord(crowd string) string {
	word := ""
	for _, ch := range crowd {
		if ch >= 'a' && ch <= 'z' {
			word += string(ch)
		}
	}
	return word
}
