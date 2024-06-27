package problem7

func swap(a *int, b *int) (int, int) {
	*a, *b = *b, *a
	return *a, *b
}
