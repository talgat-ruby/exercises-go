package problem7

import "fmt"

func swap(x, y *int) {

	temp := *x
	*x = *y
	*y = temp
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
