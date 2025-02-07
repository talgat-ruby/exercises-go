package problem7

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
