package problem7

func swap(x *int, y *int) {
	*x, *y = *y, *x
}
