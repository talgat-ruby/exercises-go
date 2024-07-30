package main

import "strings"
func emojify(s string)  string{
	s = strings.ReplaceAll(s, "smile", "🙂")
	s = strings.ReplaceAll(s, "grin", "😀")
	s = strings.ReplaceAll(s, "sad", "😥")
	s = strings.ReplaceAll(s, "mad", "😠")
	return s
	
}
