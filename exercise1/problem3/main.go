package main

import "fmt"

func numberSquares(n int) int {
    total := 0
    for k := 1; k <= n; k++ {
        total += k * k
    }
    return total
}

func main() {
    fmt.Println(numberSquares(2))
    fmt.Println(numberSquares(3))
    fmt.Println(numberSquares(4))
    fmt.Println(numberSquares(5))
}