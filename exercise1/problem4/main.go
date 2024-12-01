package main

func detectWord(word string) string {
	newWord := ""
	for _, i := range word {
		if i >= 'a' && i <= 'z' {
			newWord += string(i)
		}
	}
	return newWord
}
