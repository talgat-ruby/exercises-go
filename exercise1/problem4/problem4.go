package problem4

func mapping(input []string) map[string]string {
	output := make(map[string]string)
	for _, line := range input {
		ch := []byte(line)
		output[line] = string(toUpper(ch[0]))
	}
	return output
}

func toUpper(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		ch -= 32
	}

	return ch
}
