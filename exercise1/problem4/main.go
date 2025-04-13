package main

func detectWord(str string) string {
    res := ""

    for i := 0; i < len(str); i++ {
        if str[i] >= 'a' && str[i] <= 'z' {
            res = res + string(str[i])
        }
    }

    return res
}
