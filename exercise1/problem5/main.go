package main

func potatoes(potato string) int {
	var total int
	for i := 0; i < len(potato); i++ {
		if potato[i] == 112 && potato[i+1] == 111 && potato[i+2] == 116 &&
			potato[i+3] == 97 && potato[i+4] == 116 && potato[i+5] == 111 {
			total++
		}
	}
	return total
}
