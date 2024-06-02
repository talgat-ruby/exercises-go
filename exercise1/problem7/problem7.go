package problem7

func swap(a *int, b *int) {

	t := *a
	*a = *b
	*b = t

}
