package main

func detectWord(word string) string {
	out := make([]rune, 0, len(word))
	for _, l := range word {
		if l >= 'a' && l <= 'z' {
			out = append(out, l)
		}
	}
	return string(out)
}
