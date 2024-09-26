package problem6

import (
	"fmt"
)

func sumOfTwo(a []int, b []int, v int) bool {
	numsSet := make(map[int]struct{})

	for _, num := range a {
		numsSet[num] = struct{}{}
	}

	for _, num := range b {
		if _, exists := numsSet[v-num]; exists {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 5))
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 8))
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 3))
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 9))
}
