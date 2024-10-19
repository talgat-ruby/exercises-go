package main

func potatoes(word string) int {
	buffer := []rune("potato")
	count := 0
	bufferInd := 0
	for _, ch := range word {
		if ch == buffer[bufferInd] {
			bufferInd++
			if bufferInd == len(buffer) {
				count++
				bufferInd = 0
			}
		} else {
			bufferInd = 0
		}
	}
	return count
}
