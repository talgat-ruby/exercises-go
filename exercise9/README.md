# MyWorkout — Трекер тренировок

Это бэкенд-приложение на Go, позволяющее:
- Регистрацию пользователей (Sign Up)
- Аутентификацию через JWT (Login, Logout)
- CRUD по тренировкам (Workouts), включая 
  - Планирование (Scheduled date/time)
  - Упражнения внутри тренировки (sets, reps, weight)
- Генерацию отчётов о прошлых тренировках (по диапазону дат)
- Предзаполнение списка упражнений (Seeder)
- Защиту эндпоинтов авторизационным middleware

## Структура

- **cmd/main.go**: Точка входа (запуск сервера)
- **controllers/**: Обработчики (Handlers) для Auth, Workouts, Exercises
- **database/**: Подключение к SQLite и миграции, а также сидер
- **models/**: Определение структур (User, Exercise, Workout)
- **repositories/**: Прямая работа с БД (CRUD)
- **routes/**: Маршруты (Endpoints)
- **tests/**: Юнит-тесты (пример)
- **docker/Dockerfile**: Сборка Docker-образа
- **Makefile**: Упрощённые команды для сборки и запуска

## Запуск локально


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
  -d '{"email":"test@example.com","password":"secret"}' \
  http://localhost:8080/signup

curl -X POST -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"secret"}' \
  http://localhost:8080/login

curl -X POST -H "Authorization: Bearer <JWT_TOKEN>" \
     -H "Content-Type: application/json" \
     -d '{
       "title": "Leg Day",
       "scheduled": "2025-02-10T10:00:00Z",
       "comment": "Focus on squats",
       "exercises": [
         {"exerciseId": 2, "sets": 4, "reps": 12, "weight": 60},
         {"exerciseId": 1, "sets": 3, "reps": 10, "weight": 20}
       ]
     }' \
     http://localhost:8080/workouts

curl -H "Authorization: Bearer <JWT_TOKEN>" \
     http://localhost:8080/workouts?pending=true

curl -H "Authorization: Bearer <JWT_TOKEN>" \
     "http://localhost:8080/workouts/report?from=2025-02-01&to=2025-02-28"


curl -X POST -H "Content-Type: application/json" \
  -d '{
    "name": "Bench Press",
    "description": "Chest exercise with barbell",
    "category": "strength"
  }' \
  http://localhost:8080/exercises

curl http://localhost:8080/exercises

curl http://localhost:8080/exercises/1


curl -X PUT -H "Content-Type: application/json" \
  -d '{
    "name": "Barbell Bench Press",
    "description": "Bench press with barbell",
    "category": "strength"
  }' \
  http://localhost:8080/exercises/1

curl -X DELETE http://localhost:8080/exercises/1

```
