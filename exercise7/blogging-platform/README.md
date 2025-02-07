# MyBlog

Простое RESTful API для личной блог-платформы.  
Реализован полный CRUD для постов, а также поиск по термину.

## Структура проекта

- **cmd/main.go**: Точка входа в приложение 
- **controllers/**: HTTP-хендлеры (CRUD-операции)
- **database/**: Инициализация и миграция БД (в примере — SQLite)
- **models/**: Модели данных (структуры GORM)
- **repositories/**: Логика работы с БД (CRUD)
- **routes/**: Определение маршрутов и привязка к контроллерам
- **docker/Dockerfile**: Для сборки Docker-образа
- **Makefile**: Упрощённая сборка/запуск
- **go.mod** / **go.sum**: Go-модули

## Запуск локально 

Убедитесь, что у вас установлен Go (>= 1.18). Затем:

```bash
go mod tidy

go run ./cmd/main.go
```
или через Make
```bash
make run
make docker-build
make docker-run
```

## Примеры запросов

```
curl -X POST -H "Content-Type: application/json" \
  -d '{
    "title": "My First Blog Post",
    "content": "Hello World!",
    "category": "General",
    "tags": ["Welcome","First"]
  }' \
  http://localhost:8080/posts

curl http://localhost:8080/posts

curl "http://localhost:8080/posts?term=welcome"

curl http://localhost:8080/posts/1

curl -X PUT -H "Content-Type: application/json" \
  -d '{
    "title": "My Updated Post",
    "content": "Updated content!",
    "category": "Updates",
    "tags": ["Update","News"]
  }' \
  http://localhost:8080/posts/1

curl -X DELETE http://localhost:8080/posts/1
```