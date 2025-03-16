package problem4

func mapping(char []string) map[string]string {
    res := make(map[string]string)
	for i := 0; i<len(char); i++ {
		res[char[i]] = string(rune(char[i][0]) - 32)
	}
    return res
}
