package problem8

func simplify(list []string) map[string]int {
	indMap := make(map[string]int, len(list))
	for i, name := range list {
		indMap[name] = i
	}
	return indMap
}

func load(m *map[string]int, students *[]string) {
	for i, name := range *students {
		(*m)[name] = i
	}
}
