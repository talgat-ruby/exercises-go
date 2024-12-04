package main

func potatoes(input string) int {

    str := "potato"
    strLength := len(str)
    count := 0

    for i := 0; i <= len(input)-strLength; i++ {
        if input[i:i+strLength] == str {
            count++
        }
    }

    return count
}