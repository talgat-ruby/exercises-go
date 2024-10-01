package problem2

func capitalize(str []string) []string {
	res := []string{}
	for i := 0; i < len(str); i++ {
		name := ""
		if len(str[i]) > 0 {
			if str[i][0] > 96 && str[i][0] < 123 {
				name += string(str[i][0] - 32)
			} else {
				name += string(str[i][0])
			}
			for j := 1; j < len(str[i]); j++ {
				if str[i][j] > 64 && str[i][j] < 91 {
					name += string(str[i][j] + 32)
				} else {
					name += string(str[i][j])
				}

			}
		}
		res = append(res, name)
	}
	return res
}
