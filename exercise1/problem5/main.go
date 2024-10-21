package main

func potatoes(word string) int {
	lenWord := len(word)
	lenPotato := len("potato")
	res := 0

	if lenWord < lenPotato {
		return res
	}

	for i := 0; i <= lenWord-lenPotato; i++ {
		if word[i:i+lenPotato] == "potato" {
			res += 1
			i += lenPotato - 1
		}
	}

	return res
}
