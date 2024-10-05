package problem4

func mapping(masLetters []string) map[string]string {
	mapLetters := make(map[string]string)
	for _, v := range masLetters {
		mapLetters[v] = string(v[0] - 32)
	}
	return mapLetters
}
