package main

func potatoes(s string) int {

	count := 0
	target := "potato"
	targetLength := len(target)

	for i := 0; i <= len(s)-targetLength; i++ {
		if s[i:i+targetLength] == target {
			count++
			i += targetLength - 1 // Move the index to the end of the current match
		}
	}

	return count

}
