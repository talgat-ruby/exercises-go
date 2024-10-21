package problem7

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
