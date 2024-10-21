package main

func emojify(sentence string) string {
	res := ""
	word := ""

	if len(sentence) == 0 {
		return res
	}

	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' || sentence[i] == ',' || sentence[i] == '.' || sentence[i] == '!' || sentence[i] == '?' {
			res += isEmoji(word) + string(sentence[i])
			word = ""
		} else {
			word += string(sentence[i])
		}
	}

	return res + isEmoji(word)
}

func isEmoji(word string) string {
	if word == "smile" {
		return "🙂"
	} else if word == "grin" {
		return "😀"
	} else if word == "sad" {
		return "😥"
	} else if word == "mad" {
		return "😠"
	}
	return word
}
