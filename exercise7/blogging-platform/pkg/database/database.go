package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("postgres", "user=postgres password=1234 dbname=blog sslmode=disable")
    if err != nil {
        log.Fatal("Ошибка подключения к базе данных:", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("База данных недоступна:", err)
    }
    log.Println("Успешное подключение к базе данных")
}

