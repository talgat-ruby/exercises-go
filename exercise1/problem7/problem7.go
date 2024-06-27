package problem7

func swap(a *int, b *int) {
	*b, *a = *a, *b
}
