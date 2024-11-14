package problem5

import (
	"strings"
)

// Producer отправляет каждое слово в канал
func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word
	}
	close(ch) // Закрываем канал после отправки всех слов
}

// Consumer получает слова из канала и объединяет их в одно сообщение
func consumer(ch <-chan string) string {
	var result []string
	for word := range ch {
		result = append(result, word)
	}
	return strings.Join(result, " ")
}

// Send запускает producer и consumer для объединения слов
func send(
	words []string,
	pr func([]string, chan<- string), // Ожидаемая сигнатура producer
	cons func(<-chan string) string, // Ожидаемая сигнатура consumer
) string {
	ch := make(chan string)

	go pr(words, ch) // Запускаем producer в отдельной горутине

	return cons(ch) // Вызываем consumer и возвращаем результат
}
