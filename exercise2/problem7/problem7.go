package problem7

func swap(a *int, b *int) {
	newa := *a
	newb := *b
	*a = newb
	*b = newa

}
