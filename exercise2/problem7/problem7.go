package problem7

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
