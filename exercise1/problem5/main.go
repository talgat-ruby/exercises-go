package main

func potatoes(s string) int {
	count:=0
	sub:="potato"
	j:=0
	for i:=0;i<len(s);i++{
		if s[i]==sub[j] {
			j++
		} else {
			j=0
		}
		if j==5 {
			count++
			j=0
		}
	}
	return count
}
