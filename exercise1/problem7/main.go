package main

func highestDigit(in int) int {
	highest := 0
	for i := in; i > 0; i /= 10 {
		rem := i % 10
		if rem > highest {
			highest = rem
		}
	}
	return highest
}
