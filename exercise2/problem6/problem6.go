package problem6

import "fmt"

func sumOfTwo(a, b []int, v int) bool {

	complements := make(map[int]bool)

	for _, numA := range a {
		complements[v-numA] = true
	}

	for _, numB := range b {

		if complements[numB] {
			return true
		}
	}

	return false
}

func main() {

	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 5)) // ➞ true
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 8)) // ➞ true
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 3)) // ➞ false
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 9)) // ➞ false
}
