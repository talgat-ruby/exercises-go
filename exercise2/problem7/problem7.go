package problem7

func swap(a *int, b *int) {
	c := *a
	*b = *a
	*a = c
}
