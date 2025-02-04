package main

type wordStruct struct {
	word  []rune
	emoji string
}

func emojify(word string) string {
	buffer := []wordStruct{
		{word: []rune("smile"), emoji: "🙂"},
		{word: []rune("grin"), emoji: "😀"},
		{word: []rune("sad"), emoji: "😥"},
		{word: []rune("mad"), emoji: "😠"},
	}
	for _, ws := range buffer {
		word = replaceWord(word, &ws)
	}
	return word
}

func replaceWord(str string, ws *wordStruct) string {
	buffer := []rune(str)
	wordInd := 0
	start := make([]int, 0)
	tmp := -1
	for i, ch := range buffer {
		if ch == ws.word[wordInd] {
			if wordInd == 0 {
				tmp = i
			}
			wordInd++
			if wordInd == len(ws.word) {
				wordInd = 0
				start = append(start, tmp)
			}
		} else {
			wordInd = 0
		}
	}

	if len(start) == 0 {
		return str
	}

	out := make([]rune, 0)
	for _, i := range start {
		out = append(out, buffer[:i]...)
		out = append(out, []rune(ws.emoji)...)
		out = append(out, buffer[i+len(ws.word):]...)
	}
	return string(out)
}
