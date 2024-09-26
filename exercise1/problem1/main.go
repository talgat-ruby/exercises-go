package main

func main() {
	//fmt.Println(5)
}

func addUp(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
