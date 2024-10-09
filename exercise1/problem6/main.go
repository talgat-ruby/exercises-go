package main

import "fmt"

func main() {
	fmt.Println(emojify("smile, grin, sad, mad"))
}

func emojify(s string) string {
	words := splitIntoWords(s)
	fmt.Println(words)
	text := ""

	for _, word := range words {
		text += getEmoji(word)
	}
	return text

}
func getEmoji(word string) string {
	switch word {
	case "smile":
		return "🙂"
	case "grin":
		return "😀"
	case "sad":
		return "😥"
	case "mad":
		return "😠"
	default:
		return word
	}
}
func isSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r' || char == '(' || char == ')' || char == ','
}

func splitIntoWords(text string) []string {
	var words []string
	tempText := ""

	for i := 0; i < len(text); i++ {
		if i == len(text)-1 && !isSpace(text[i]) {
			tempText += string(text[i])
			words = append(words, tempText)
		} else if !isSpace(text[i]) {
			tempText += string(text[i])
		} else {
			words = append(words, tempText)
			words = append(words, string(text[i]))
			tempText = ""
		}
	}

	return words
}
