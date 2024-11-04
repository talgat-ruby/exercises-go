package problem8

import (
	"time"
)

func withTimeout(ch <-chan string, ttl time.Duration) string {
	// Создаем таймер с заданным временем жизни
	timer := time.NewTimer(ttl)

	select {
	case msg := <-ch:
		// Сообщение получено до истечения времени
		return msg
	case <-timer.C:
		// Таймер сработал, время истекло
		return "timeout"
	}
}
