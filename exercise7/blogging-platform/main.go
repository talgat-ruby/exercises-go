package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"blogging-platform/internal/api"
	"blogging-platform/internal/db"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Подключение к БД
	_, err := db.InitDB() // Теперь вызываем `InitDB()`
	if err != nil {
		log.Fatalf("❌ Ошибка подключения к БД: %v", err)
	}

	// Запуск API
	a := api.New()
	if err := a.Start(ctx); err != nil {
		log.Fatalf("❌ Ошибка запуска API: %v", err)
	}

	// Завершаем работу при нажатии Ctrl+C
	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		sig := <-shutdown
		fmt.Println("⚠️ Получен сигнал завершения:", sig)
		cancel()
	}()
}
