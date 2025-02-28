package main

import "strings"

func potatoes(input string) int {

	searchPotatoes := "potato"
	count := strings.Count(input, searchPotatoes)

	return count
}
