package main

import (
	"fmt"
)

func detectWord(crowd, word string) string {
	result := ""
	wordIndex := 0
	for _, char := range crowd {
		if char == rune(word[wordIndex]) {
			result += string(char)
			wordIndex++
			if wordIndex == len(word) {
				break
			}
		}
	}
	return result
}
func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
	fmt.Println(detectWord("bEEFGBuFBRrHgUHlNFYaYr"))
	fmt.Println(detectWord("YFemHUFBbezFBYzFBYLleGBYEFGBMENTment"))
}

//##не совсем понял где ошибка и в отладке пытался понять,смотреть и вообще нормально ли то что я все гуглю где только возможно, используя все возможности интернета? т.к интенсив и вроде как вводный курс для новичков(p.s Меняю проф деятельтность),но голова ломается при решений, понятно что не ничего понятно. Но все это очень интересно,хоть и не совсем понятно.
