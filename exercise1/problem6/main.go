package main

import (
	"strings"
)

func emojify(a string) string {
	emoji := []string{"smile", "grin", "sad", "mad"}
	emojires := []string{"🙂", "😀", "😥", "😠"}

	arr := strings.Split(a, " ")

	res := ""

	for _, i := range arr {
		em, tr := smileChek(i, emoji, emojires)
		if tr {
			res += em
		} else {
			res += i
		}
		res += " "
	}
	res = res[:len(res)-1]
	return res
}

func smileChek(a string, b []string, c []string) (string, bool) {
	char := ""
	if a[len(a)-1] < 'A' || a[len(a)-1] > 'z' {
		char = string(a[len(a)-1])
		a = a[:len(a)-1]
	}

	for j, i := range b {
		if i == a {
			if char != "" {
				return c[j]+char, true
			}
			return c[j], true
		}
	}
	return "", false
}
