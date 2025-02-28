package problem7

func swap(x, y *int) {
	*x, *y = *y, *x
}
