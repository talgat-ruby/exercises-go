package main

func detectWord(word string) string {
	letters := []rune(word)
	out := make([]rune, 0, len(letters))
	for _, r := range letters {
		if r >= 'a' && r <= 'z' {
			out = append(out, r)
		}
	}
	return string(out)
}
