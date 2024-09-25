package problem4

func mapping(inp []string) map[string]string {
	out := make(map[string]string)
	for _, in := range inp {
		var oneWord string
		for _, v := range in {
			oneWord += string(v - 32)
		}
		out[in] = oneWord
	}
	return out
}
