package problem7

func swap(a *int, b *int) (*int, *int) {
	t := *a
	*a = *b
	*b = t

	return a, b
}
