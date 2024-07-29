package main

import (
	"strings"
)

func potatoes(crowd string) int {
	count := 0
	substr := "potato"
	for {
		index := strings.Index(crowd, substr)
		if index == -1 {
			break
		}
		count++
		crowd = crowd[index+len(substr):] // Продвигаем строку после найденного вхождения
	}
	return count
}
