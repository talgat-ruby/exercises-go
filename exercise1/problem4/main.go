package main

func detectWord(s string) string {
	var str []rune
	for _, r := range s {
		if r >= 97 && r <= 122 {
			str = append(str, r)
		}
	}
	return string(str)
}
