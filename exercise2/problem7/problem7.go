package problem7

func swap(a, b *int) {
	*a, *b = *b, *a
}
