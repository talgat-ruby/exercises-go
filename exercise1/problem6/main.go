package main

import "strings"

func emojify(s string) string {
	s = strings.Replace(s, "smile", "🙂", -1)
	s = strings.Replace(s, "grin", "😀", -1)
	s = strings.Replace(s, "sad", "😥", -1)
	s = strings.Replace(s, "mad", "😠", -1)
	return s
}
