package problem1

func isChangeEnough(nums [4]int, num float32) bool {

	var res float32
	n := [4]float32{0.25, 0.10, 0.05, 0.01}
	for i := 0; i < len(nums); i++ {
		count := float32(nums[i]) * n[i]
		res += count
	}

	return res >= num
}
