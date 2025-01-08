package problem4

func detectWord(crowd string) string {
    word := ""
    for _, char := range crowd {
        // Check if the character is lowercase
        if char >= 'a' && char <= 'z' {
            word += string(char)
        }
    }
    return word
}
