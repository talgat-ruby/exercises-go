package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(num int, corner dir) [][]int {
	row := corner[0]
	col := corner[1]
	arr := make([][]int, num)

	if row == 'u' {
		for i := 0; i < num; i++ {
			arr[i] = make([]int, num)
			if col == 'l' {
				// ul
				arr[i] = iteration(i, num, true)
				continue
			}
			// ur
			arr[i] = iteration(num+i-1, num, false)
		}
	} else {
		for i := num - 1; i >= 0; i-- {
			arr[i] = make([]int, num)
			if col == 'l' {
				// ll
				arr[i] = iteration(num-1-i, num, true)
				continue
			}
			// lr
			arr[i] = iteration(num+num-2-i, num, false)
		}
	}
	return arr
}

func iteration(start int, size int, isAsc bool) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		if isAsc {
			arr[i] = start + i
		} else {
			arr[i] = start - i
		}
	}
	return arr
}
