package problem7

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
