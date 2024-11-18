package main

import "fmt"

func addUp(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

func main() {
    fmt.Println(addUp(4))   // 10
    fmt.Println(addUp(13))  // 91
    fmt.Println(addUp(600)) // 180300
}

