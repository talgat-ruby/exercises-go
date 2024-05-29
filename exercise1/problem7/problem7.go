package problem7

func swap(a *int, b *int) (int, int) {
	temp := *a
	*a = *b
	*b = temp

	return *a, *b
}
