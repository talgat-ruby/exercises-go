package main

func potatoes(str string) int {
	count := 0
	target := "potato"
	targetLen := len(target)
	strLen := len(str)
	for i := 0; i <= strLen-targetLen; i++ {
		if str[i:i+targetLen] == target {
			count++
			i += targetLen - 1
		}
	}
	return count
}
