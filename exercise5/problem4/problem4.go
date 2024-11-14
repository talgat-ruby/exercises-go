package problem4

import "fmt"

func iter(ch chan<- int, nums []int) {
	for _, n := range nums {
		ch <- n
	}
	close(ch) //
}

func sum(nums []int) int {
	ch := make(chan int)
	go iter(ch, nums)

	var sum int
	for n := range ch {
		sum += n
	}
	return sum
}

func main() {
	nums := []int{1}            //
	result := sum(nums)         //
	fmt.Println("Sum:", result) //
}
