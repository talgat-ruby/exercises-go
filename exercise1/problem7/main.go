package main

func highestDigit(n int) (maxDigit int) {
	for n > 0 {
		if n%10 > maxDigit {
			maxDigit = n % 10
		}
		n /= 10
	}
	return
}
