package problem2

func capitalize(names []string) []string {
	var capitalizedNames []string

	for _, name := range names {
		if len(name) == 0 {
			capitalizedNames = append(capitalizedNames, name)
			continue
		}

		nameBytes := []byte(name)
		nameBytes[0] = toUpper(nameBytes[0])

		for i := 1; i < len(nameBytes); i++ {
			nameBytes[i] = toLower(nameBytes[i])
		}

		capitalizedNames = append(capitalizedNames, string(nameBytes))
	}

	return capitalizedNames
}

func toUpper(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		ch -= 32
	}

	return ch
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		ch += 32
	}

	return ch
}
