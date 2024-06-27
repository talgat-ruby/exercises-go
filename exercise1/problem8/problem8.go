package problem8

func simplify(list []string) map[string]int {
	newMap := make(map[string]int, len(list))

	load(newMap, list)

	return newMap
}

func load(newMap map[string]int, list []string) {
	for i, name := range list {
		newMap[name] = i
	}
}