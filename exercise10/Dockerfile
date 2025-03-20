# Используем официальный образ Golang
FROM golang:1.21

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем приложение
RUN go build -o main .

# Запускаем приложение
CMD ["/app/main"]
