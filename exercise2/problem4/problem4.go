package problem4

func mapping(a []string) map[string]string {
	res := make(map[string]string)
	letter := ""
	for _, v := range a {
		letter = string(v[0] - 32)
		res[v] = letter
		letter = ""
	}
	return res
}
