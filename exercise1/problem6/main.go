package main

import "strings"

func emojify(s string) string {
	smile, grin, sad, mad := "🙂", "😀", "😥", "😠"
	s = strings.Replace(s, "smile", smile, -1)
	s = strings.Replace(s, "grin", grin, -1)
	s = strings.Replace(s, "sad", sad, -1)
	s = strings.Replace(s, "mad", mad, -1)
	return s
}
