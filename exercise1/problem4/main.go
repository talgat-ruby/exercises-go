package main

import "fmt"

func main() {
	fmt.Println(detectWord("UcUNFYGaFYFYGtNUH"))
}

func detectWord(s string) string {
	letter := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			letter += string(c)
		}
	}
	return letter
}
