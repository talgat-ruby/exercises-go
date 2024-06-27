package problem1

func isChangeEnough(changes [4]int, total float32) bool {
	totalAsCents := int(total * 100)
	totalAsCents -= min(totalAsCents/25, changes[0]) * 25
	totalAsCents -= min(totalAsCents/10, changes[1]) * 10
	totalAsCents -= min(totalAsCents/5, changes[2]) * 5
	totalAsCents -= min(totalAsCents, changes[3])

	if totalAsCents != 0 && changes[0] > 0 && total > 0.25 {
		totalAsCents = int(total * 100)
		totalAsCents -= min(totalAsCents/25-1, changes[0]-1) * 25
		totalAsCents -= min(totalAsCents/10, changes[1]) * 10
		totalAsCents -= min(totalAsCents/5, changes[2]) * 5
		totalAsCents -= min(totalAsCents, changes[3])
	}

	return totalAsCents == 0
}
