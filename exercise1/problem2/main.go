package main

import (
	"fmt"
	"strconv"
)

func binary(de int) string {
	return strconv.FormatInt(int64(de), 2)
}

func main() {
	var b int
	fmt.Scan(&b)
	res := binary(b)
	fmt.Println(res)
}
