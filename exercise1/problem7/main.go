package main

import (
    "fmt"
    "strconv"
)

func highestDigit(num int) int {
    strNum := strconv.Itoa(num)
    maxDigit := 0
    for _, char := range strNum {
        digit, _ := strconv.Atoi(string(char))
        if digit > maxDigit {
            maxDigit = digit
        }
    }
    return maxDigit
}

func main() {
    fmt.Println(highestDigit(379))
    fmt.Println(highestDigit(2))
    fmt.Println(highestDigit(377401))
}
