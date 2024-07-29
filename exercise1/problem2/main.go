package main

import (
	"fmt"
	"strconv"
)

func binary(d int) string {
	if d == 0 {
		return "0"
	}

	var bin string
	var rem int
	for d > 0 {
		rem = d & 2
		fmt.Print(rem)
		bin = strconv.Itoa(rem) + bin
		fmt.Print(bin)
		d = d / 2
		fmt.Print(d)
	}
	return bin

}

func main() {
	var b int
	fmt.Scan(&b)
	res := binary(b)
	fmt.Println(res)
}
