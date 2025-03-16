package main

import (
	"strings"
)

func emojify(s string) string {
	var emotions = [4]string{"smile", "sad", "grin", "mad"}
	var emojis = [4]string{"🙂", "😥", "😀", "😠"}

	for i, emotion := range emotions {
		k := strings.Count(s, emotion)
		if k != 0 {
			s = strings.Replace(s, emotion, emojis[i], k)
		}
	}
	return s
}
