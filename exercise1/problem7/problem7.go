package problem7

func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
