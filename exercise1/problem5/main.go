package main

import (
	"fmt"
	"strings"
)

func potatoes(input string) int {
	return strings.Count(input, "potato")

}
func main() {
	fmt.Println(potatoes("potato"))
	fmt.Println(potatoes("potatopotato"))
	fmt.Println(potatoes("potatoapple"))
}
