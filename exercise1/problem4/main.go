package main

func detectWord(s string) string {
	result:=""
	for i:=0;i<len(s);i++ {
		if s[i]>='a' && s[i]<='z' {
			result+=string(s[i])
		}
	}
	return result
}
