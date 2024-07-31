package main

import "fmt"

func main() {
	var s string = "helldfsgdfgEFEo"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}
