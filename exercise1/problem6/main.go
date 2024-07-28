package main

func emojify(input string) string {
    runes := []rune(input)
    result := make([]rune, 0, len(runes))

    for i := 0; i < len(runes); {
        if i+4 < len(runes) && string(runes[i:i+5]) == "smile" {
            result = append(result, '🙂')
            i += 5
        } else if i+2 < len(runes) && string(runes[i:i+3]) == "sad" {
            result = append(result, '😥')
            i += 3
        } else if i+3 < len(runes) && string(runes[i:i+4]) == "grin" {
            result = append(result, '😀')
            i += 4
        } else if i+2 < len(runes) && string(runes[i:i+3]) == "mad" {
            result = append(result, '😠')
            i += 3
        } else {
            result = append(result, runes[i])
            i++
        }
    }

    return string(result)
}