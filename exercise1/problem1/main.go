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
    result := addUp(10)
    fmt.Println("The sum is:", result)
}
