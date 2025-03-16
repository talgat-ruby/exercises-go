package main

import (
	"strings"
)

func emojify(str string) string {
	if strings.Contains(str, "smile") {
		str = strings.ReplaceAll(str, "smile", "🙂")
	}
	if strings.Contains(str, "grin") {
		str = strings.ReplaceAll(str, "grin", "😀")
	}
	if strings.Contains(str, "sad") {
		str = strings.ReplaceAll(str, "sad", "😥")
	}
	if strings.Contains(str, "mad") {
		str = strings.ReplaceAll(str, "mad", "😠")

	}
	return str
}
