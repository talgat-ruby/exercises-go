package main

import (
    "fmt"
    "strconv"
)

func binary(decimalStr string) (string, error) {
    decimal, err := strconv.Atoi(decimalStr)
    if err != nil {
        return "", err
    }
    binaryStr := strconv.FormatInt(int64(decimal), 2)
    return binaryStr, nil
}

func main() {
    results := []string{"1", "5", "10"}
    for _, r := range results {
        binaryStr, err := binary(r)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("binary(%s) = %s\n", r, binaryStr)
        }
    }
}