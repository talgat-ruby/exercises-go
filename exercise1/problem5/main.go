package main

func potatoes(s string) int {

	k := 0
	for i, r := range s {
		if r == 'p' {
			if string(s[i:i+6]) == "potato" {
				k++

			}
		}
	}
	return k
}
