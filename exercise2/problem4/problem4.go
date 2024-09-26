package problem4

import "strings"

func mapping(input []string) map[string]string {

	result := make(map[string]string)

	for _, letter := range input {
		// Добавляем в карту пару (строчная буква : заглавная буква)
		result[letter] = strings.ToUpper(letter)
	}

	return result

}
