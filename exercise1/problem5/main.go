package main

import "fmt"

func main() {
	fmt.Println(potatoes("potatoapple"))
}

func potatoes(s string) int {
	target := "potato"
	count := 0
	targetLen := len(target)

	for i := 0; i <= len(s)-targetLen; {
		match := true
		for j := 0; j < targetLen; j++ {
			if s[i+j] != target[j] {
				match = false
				break
			}
		}
		if match {
			count++
			i += targetLen
		} else {
			i++
		}
	}
	return count
}
