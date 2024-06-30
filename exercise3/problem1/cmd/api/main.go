package main

import (
	"github.com/talgat-ruby/exercises-go/exercise3/problem1/internal/api"
)

func main() {
	// Создаем конфиг
	// appConfig := config.NewConfig("./.env")

	// fmt.Printf("User: %s\n", appConfig.Database.User)
	// fmt.Printf("Host: %s\n", appConfig.Database.Host)
	// fmt.Printf("Password: %s\n", appConfig.Database.Password)
	// fmt.Printf("DbName: %s\n", appConfig.Database.DbName)
	// fmt.Printf("Port: %d", appConfig.Database.Port)

	api.Run()
}
