package problem2

func capitalize(names []string) []string {
	var capitalizeNames []string
	for _, name := range names {
		var capitalizeName string
		for i, char := range name {
			if i == 0 && (char >= 'a' && char <= 'z') {
				char -= 'a' - 'A'
			} else if i != 0 && (char >= 'A' && char <= 'Z') {
				char += 'a' - 'A'
			}
			capitalizeName += string(char)
		}
		capitalizeNames = append(capitalizeNames, capitalizeName)
	}
	return capitalizeNames
}
