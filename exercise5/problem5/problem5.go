package problem5

import (
	"strings"
)

func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word // Отправляем каждое слово в канал
	}
	close(ch) // Закрываем канал после отправки всех слов
}

func consumer(ch <-chan string) string {
	var result []string
	for word := range ch {
		result = append(result, word) // Добавляем каждое слово в срез
	}
	return strings.Join(result, " ") // Объединяем слова в одну строку с пробелами
}

func send(
	words []string,
	pr func([]string, chan<- string),
	cons func(<-chan string) string,
) string {
	ch := make(chan string)

	go pr(words, ch)

	return cons(ch)
}
