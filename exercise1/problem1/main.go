package main

import "fmt"

// Функция, которую вы уже написали
func addUp(n int) int {
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}

// Функция main, которая является точкой входа в программу
func main() {
    // Пример использования функции addUp
    result := addUp(10)
    fmt.Println("The sum is:", result)
}
