package problem9

func factory(number int) func(...int) []int {
	return func(mas ...int) []int {
		for i := 0; i < len(mas); i++ {
			mas[i] = mas[i] * number
		}
		return mas
	}
}
