package main

func detectWord(str string) string {
	var result string
	for _, letter := range str {
		if letter <= 122 && letter >= 97 {
			result = result + string(letter)
		}
	}
	return result
}
