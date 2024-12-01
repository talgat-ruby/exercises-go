package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countVowels(s string) int {
	s = strings.ToLower(s)

	vowels := "aeiou"

	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) { //чекает если в текущей букве из стринга есть руна (char) из vowels
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	res := countVowels(input)

	fmt.Println(res)
}
