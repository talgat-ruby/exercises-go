package main

func addUp(number int) int {
	sum := 0
	for i := 0; i <= number; i++ {
		sum += i
	}

	return sum
}

func main() {
	print(addUp(13))
}
