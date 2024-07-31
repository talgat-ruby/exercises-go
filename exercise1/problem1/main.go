package main

func addUp(num int) int {
	if num == 1 {
		return num
	}
	return (addUp((num - 1)) + num)
}
