package util

func ToUpper(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if 'a' <= runes[i] && runes[i] <= 'z' {
			runes[i] -= ' '
		}
	}

	return string(runes)
}

func CapitalizeString(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)
	if 'a' <= runes[0] && runes[0] <= 'z' {
		runes[0] -= ' '
	}

	for i := 1; i < len(runes); i++ {
		if 'A' <= runes[0] && runes[i] <= 'Z' {
			runes[i] += ' '
		}
	}

	return string(runes)
}
