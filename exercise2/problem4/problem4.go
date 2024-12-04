package problem4

func mapping(letters []string) map[string]string {
	result := make(map[string]string)
	for _, letter := range letters {
		result[letter] = string(letter[0] - 32)
	}
	return result
}
