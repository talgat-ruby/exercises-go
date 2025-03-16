package problem4

func mapping(arr []string) map[string]string {
	letterMap := make(map[string]string)

	for _, key := range arr {
		letterMap[key] = toUpper(key)
	}
	return letterMap
}

func toUpper(str string) string {
	return string(rune(str[0]) - 32)
}
