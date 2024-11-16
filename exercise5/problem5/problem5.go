package problem5

import "strings"

func producer(words []string, ch chan<- string) {
	// Отправляем каждое слово из words в канал
	for _, word := range words {
		ch <- word
	}
	// Закрываем канал после отправки всех слов
	close(ch)
}

func consumer(ch <-chan string) string {
	// Читаем слова из канала и объединяем их в строку
	var result []string
	for word := range ch {
		result = append(result, word)
	}
	// Возвращаем объединенную строку
	return strings.Join(result, " ")
}

func send(words []string, pr func([]string, chan<- string), cons func(<-chan string) string) string {
	ch := make(chan string)

	go pr(words, ch)

	return cons(ch)
}
