package problem7

import (
	"fmt"
)

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
