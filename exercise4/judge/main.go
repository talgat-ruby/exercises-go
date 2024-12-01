package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/api"
)

func main() {
	// создается контекст с возможностью отмены (cancel), это позволяет завершить любые операции, которые используют этот контекст, когда будет вызвана функция cancel.
	// Контекст используется для управления временем жизни сервера и его завершения.
	ctx, cancel := context.WithCancel(context.Background())

	// получение переменной из переменной окружения ) из переменной окружения считывается порт на котором будет запущен HTTP-сервер
	port := os.Getenv("PORT")

	// создание экземпляра структуры API, которая управляет запуском и остановкой http-сервера
	a := api.New()

	// здесь вызывается метод START API, который запускает сервер на указанном порту. Если при запуске возникнет ошибка, она будет залогирована и программа завершится с кодом выхода 1.
	if err := a.Start(ctx, port); err != nil {
		slog.ErrorContext(ctx, "api start error", "error", err)
		os.Exit(1)
	}

	/* обработка системных сигналов для остановки сервера
	   в этой горутине программа ожидает получения системного сигнала (sigint или нажатие ctrl+c). когда сигнал поступает,
	   он записывается в канал shutdown, после чего: логируется полученный сигнал, вызывается функция cancel, которая завершает контекст ctx что приводит к остановке всех операций завязанных на этот контекст.
	*/
	go func() {
		shutdown := make(chan os.Signal, 1)   // Create channel to signify s signal being sent
		signal.Notify(shutdown, os.Interrupt) // When an interrupt is sent, notify the channel

		sig := <-shutdown
		slog.WarnContext(ctx, "signal received - shutting down...", "signal", sig)

		cancel()
	}()

	//Программа будет заблокирована до тех пор, пока контекст ctx не будет завершен (например, при получении системного сигнала). Когда контекст завершится, программа продолжит выполнение.
	<-ctx.Done()

	// После завершения контекста происходит выполнение завершающих операций, которые могут включать в себя освобождение ресурсов, закрытие соединений и т.д.
	slog.InfoContext(ctx, "cleaning up ...")

	// Здесь вызывается метод Stop API, который останавливает сервер. Если при остановке возникнет ошибка, она будет залогирована.
	if err := a.Stop(ctx); err != nil {
		slog.ErrorContext(ctx, "service stop error", "error", err)
	}

	// После успешной остановки сервера логируется сообщение о том, что сервер был корректно завершен.
	slog.InfoContext(ctx, "server was successfully shutdown.")
}
