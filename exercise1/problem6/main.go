package main

import (
	"strings"
)

func emojify(str string) string {
	str = strings.Replace(str, "smile", "🙂", 1)

	str = strings.Replace(str, "grin", "😀", 1)

	str = strings.Replace(str, "sad", "😥", 1)

	str = strings.Replace(str, "mad", "😠", 1)

	return str
}
