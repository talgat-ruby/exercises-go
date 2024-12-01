package main

func highestDigit(n int) int {
	max:=0
	for n>0 {
		if n%10>max {
			max=n%10
		}
		n/=10
	}
	return max
}
