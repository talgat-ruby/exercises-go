package main

import "strings"

func emojify(str string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(str, "smile", "🙂", -1),
				"grin",
				"😀",
				-1,
			),
			"sad",
			"😥",
			-1,
		),
		"mad",
		"😠",
		-1,
	)
}
