package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/talgat-ruby/exercises-go/exercise3/problem2/internal/request"
)

func Execute() {
	// Определение флагов
	headers := flag.String("H", "", "Заголовки для запроса")
	data := flag.String("d", "", "Данные для POST, PUT, PATCH запросов")
	help := flag.Bool("h", false, "Отобразить помощь")

	flag.Parse()

	// Отображение помощи, если передан флаг -h
	if *help {
		showHelp()
		return
	}

	// Проверка наличия аргументов метод и url
	if len(flag.Args()) < 2 {
		fmt.Println("Ошибка: Метод и URL обязательны")
		os.Exit(1)
	}

	method := flag.Arg(0)
	url := flag.Arg(1)

	// Создание нового HTTP запроса
	err := request.MakeRequest(method, url, *headers, *data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
}

func showHelp() {
	fmt.Println("Использование: go run main.go [метод] [url] [флаги]")
	fmt.Println("Методы: GET, POST, PUT, PATCH, DELETE")
	fmt.Println("Флаги:")
	fmt.Println("  -H string")
	fmt.Println("        Заголовки для запроса")
	fmt.Println("  -d string")
	fmt.Println("        Данные для POST, PUT, PATCH запросов")
	fmt.Println("  -h")
	fmt.Println("        Отобразить помощь")
}
