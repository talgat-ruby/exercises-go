package main

func detectWord(word string) string {
	wanted := ""

	for _, char := range word {
		// if unicode.IsLower(char) {
		// 	wanted += string(char)
		// }
		code := int(char)
		if code >= 97 && code <= 122 {
			wanted += string(char)
		}
	}

	return wanted
}
