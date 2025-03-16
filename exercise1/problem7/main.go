package main

func highestDigit(n int) int {
	var arr []int
	for n > 10 {
		arr = append(arr, n%10)
		n = n / 10
	}
	arr = append(arr, n)
	max := -1
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
