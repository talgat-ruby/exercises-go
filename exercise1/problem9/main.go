package main

import (
    "fmt"
)

func toBinary(num int) string {
    return fmt.Sprintf("%08b", num)
}

func bitwiseAND(a, b int) string {
    return toBinary(a & b)
}

func bitwiseOR(a, b int) string {
    return toBinary(a | b)
}

func bitwiseXOR(a, b int) string {
    return toBinary(a ^ b)
}

func main() {
    a, b := 6, 23
    fmt.Println("bitwiseAND:", bitwiseAND(a, b))
    fmt.Println("bitwiseOR:", bitwiseOR(a, b))
    fmt.Println("bitwiseXOR:", bitwiseXOR(a, b))
}