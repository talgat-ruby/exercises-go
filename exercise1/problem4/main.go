package main

func detectWord(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 123 {
			res += string(s[i])
		}
	}
	return res
}
