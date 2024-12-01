package main

import (
	"fmt"
	"strconv"
)

func highestDigit(n int) int {
	strnum := strconv.Itoa(n) // переводим число в стринг чтобы итерировать по нему

	highest := 0
	for _, asciidigit := range strnum {
		digitInt := int(asciidigit - '0') //Это переводит значение цифры в АСКИ в саму цифру, АСКИ 0 (имеет значение 48) - 48 (='0') = 0

		if digitInt > highest {
			highest = digitInt
		}
	}

	return highest
}

func main() {
	var b int
	fmt.Scan(&b)
	res := highestDigit(b)
	fmt.Println(res)
}
