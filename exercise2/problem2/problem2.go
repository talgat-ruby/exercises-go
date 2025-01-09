package problem2

func capitalize(arr []string) []string {
	if len(arr) < 1 {
		return arr
	}
	for i, el := range arr {
		if len(el) == 0 {
			continue
		}
		arr[i] = toCapital(el)
	}

	return arr
}

func toCapital(s string) string {
	runes := []rune(s)
	for i, char := range runes {
		if i == 0 {
			if char >= 'a' && char <= 'z' {
				runes[i] = char - 32
				continue
			}
			runes[i] = char
		} else {
			if char >= 'A' && char <= 'Z' {
				runes[i] = char + 32
				continue
			}
			runes[i] = char
		}
	}

	return string(runes)
}
