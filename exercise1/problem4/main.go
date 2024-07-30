package main

import (
	"unicode"
)

func detectWord(s string) string {
	var ans string
 	for _, c := range s{
		if unicode.IsLower(c){
			ans += string(c)
		}
	}
	return ans
}
