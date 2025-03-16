package main

import (
    "fmt"
    "strconv"
)

func sum(a, b string) (string, error) {
    num1, err1 := strconv.Atoi(a)
    num2, err2 := strconv.Atoi(b)
    
    if err1 != nil {
        return "", fmt.Errorf("string: %s cannot be converted", a)
    }
    if err2 != nil {
        return "", fmt.Errorf("string: %s cannot be converted", b)
    }

    return strconv.Itoa(num1 + num2), nil
}

func main() {
    testCases := []struct {
        a, b string
    }{
        {"1", "2"},
        {"10", "20"},
        {"a", "2"},
    }

    for _, tc := range testCases {
        result, err := sum(tc.a, tc.b)
        if err != nil {
            fmt.Println(err)
        } else {
            fmt.Println(result)
        }
    }
}
