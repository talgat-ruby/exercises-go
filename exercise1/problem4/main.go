package main

func detectWord(a string) string {
	res := ""

	for _, i := range a {
		if i >='a' && i <= 'z' {
			res += string(i)
		}
	}
	return res
}
