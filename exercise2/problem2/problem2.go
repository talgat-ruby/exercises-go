package problem2

func capitalize(names []string) []string {
	result := make([]string, len(names))
	for j, name := range names {
		var oneWord string
		for i, r := range name {
			if r >= 'A' && r <= 'Z' && i == 0 {
				oneWord += string(r)
			}
			if r >= 'a' && r <= 'z' && i == 0 {
				oneWord += string(r - 32)
			}
			if r >= 'A' && r <= 'Z' && i != 0 {
				oneWord += string(r + 32)
			}
			if r >= 'a' && r <= 'z' && i != 0 {
				oneWord += string(r)
			}
		}
		result[j] = oneWord
	}
	return result
}
