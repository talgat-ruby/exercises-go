package main

func detectWord(str string) string {
	word := ""
	for _, ch := range str {
		if ch > 96 && ch < 123 {
			word += string(ch)
		}
	}
	return word
}
