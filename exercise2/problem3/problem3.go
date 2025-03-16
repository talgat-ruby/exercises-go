package main

type dir string

const (
	ul dir = "ul"
	ur dir = "ur"
	ll dir = "ll"
	lr dir = "lr"
)

func diagonalize(n int, d dir) [][]int {
	nums := []int{}
	res := [][]int{}
	if d == ul {
		count := 0
		for i := 0; i <n; i++ {
			for j := 0; j <n; j++ {
				nums = append(nums, count+j)
			}
			res = append(res, nums)
			nums = []int{}
			count++
		}
		return res
	}
	if d == ur {
		count := n-1
		for i := 0; i <n; i++ {
			for j := 0; j <n; j++ {
				nums = append(nums, count-j)
			}
			res = append(res, nums)
			nums = []int{}
			count++
		}
		return res
	}
	if d == ll {
		count := n-1
		for i := 0; i <n; i++ {
			for j := 0; j <n; j++ {
				nums = append(nums, count+j)
			}
			res = append(res, nums)
			nums = []int{}
			count--
		}
		return res
	}
	if d == lr {
		count := n + n - 2
		for i := 0; i <n; i++ {
			for j := 0; j <n; j++ {
				nums = append(nums, count-j)
			}
			res = append(res, nums)
			nums = []int{}
			count--
		}
		return res
	}
	return res
}
