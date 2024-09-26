package problem7

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
