package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Структура для запроса на /join
type JoinRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Обработчик для Ping
func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "pong")
}

func main() {
	/*
		Создание контекста и отложенная отмена
		Создается контекст (ctx), который может быть отменен в любой момент с помощью функции cancel(). Контекст используется для контроля времени жизни операций, таких как работа HTTP-сервера.
		defer cancel() — контекст будет отменен автоматически, когда функция main завершится. Это важно для корректного завершения сервера.
	*/
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	/*
		Создание маршрутизатора и обработчика для пинга
		Создается новый HTTP маршрутизатор (ServeMux), который отвечает за обработку маршрутов (путей) запросов.
		Добавляется обработчик для пути /ping, который вызывает функцию pingHandler. Этот обработчик отвечает на пинг-запросы от сервера судейства, проверяющего, что бот активен и может принимать запросы.
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/move", moveHandler)

	/*
		Запуск сервера
		Вызов функции startServer(ctx, mux) запускает HTTP-сервер, используя переданный контекст и маршрутизатор.
		Канал ready сигнализирует, что сервер успешно запущен и готов принимать запросы. Программа приостанавливается здесь до тех пор, пока сервер не будет готов к работе.
	*/
	ready := startServer(ctx, mux)
	<-ready

	// инфо о запуске сервера
	fmt.Println("bot server started")

	// Получение имени бота из переменных окружения
	botName := os.Getenv("NAME")
	if botName == "" {
		botName = "DefaultBotName" // Можно задать имя по умолчанию
	}

	/*
		Присоединение к игре (join)
		Функция joinGame отправляет запрос на сервер судейства с именем бота и его URL, по которому сервер судейства сможет обращаться к боту.
		Бот слушает на порту, указанном в переменной окружения PORT. Если отправка запроса на присоединение не удалась (например, ошибка соединения или неправильный формат данных), программа выводит ошибку и завершает работу.
	*/
	err := joinGame(botName, fmt.Sprintf("http://localhost:%s", os.Getenv("PORT")))
	if err != nil {
		fmt.Println("Failed to join game:", err)
		return
	}

	/*
		Ожидание сигнала завершения
		Создается канал stop, который будет получать сигналы системы.
		С помощью функции signal.Notify программа подписывается на сигналы завершения, такие как SIGINT (нажатием Ctrl+C) или SIGTERM.
		Программа приостанавливается до тех пор, пока не получит один из этих сигналов.
	*/
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop // Wait for SIGINT or SIGTERM

	/*
		Завершение работы сервера
		Когда программа получает сигнал завершения, выводится сообщение, что сервер бота завершает работу.
		Контекст отменяется с помощью вызова cancel(), что инициирует завершение всех операций, включая остановку сервера.
	*/
	fmt.Println("Shutting down bot server...")
	cancel()
}

type RequestMove struct {
	Board Board `json:"board"`
	Token Token `json:"token"`
}

type Token string

const (
	TokenEmpty Token = " "
	TokenX     Token = "x"
	TokenO     Token = "o"
)

const (
	Cols = 3
	Rows = 3
)

type Board [Cols * Rows]Token

type ResponseMove struct {
	Index int `json:"index"`
}

// обработчик для хода
func moveHandler(w http.ResponseWriter, r *http.Request) {
	// чтение тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	fmt.Println(body)
	// Разбор JSON-данных из запроса
	var reqMove RequestMove
	if err := json.Unmarshal(body, &reqMove); err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusBadRequest)
		return
	}

	// Логирование состояния доски и токена
	fmt.Printf("Received move request for token %s on board:\n%s\n", reqMove.Token, reqMove.Board.String())

	// Выбор случайного индекса для хода (можно заменить на логику принятия решения)
	rand.Seed(time.Now().UnixNano())
	availableMoves := reqMove.Board.AvailableMoves() // Получаем все доступные ходы
	if len(availableMoves) == 0 {
		http.Error(w, "No available moves", http.StatusBadRequest)
		return
	}

	randomMove := availableMoves[rand.Intn(len(availableMoves))] // Выбираем случайный ход

	// Подготовка ответа
	respMove := ResponseMove{
		Index: randomMove,
	}

	// Возвращаем JSON-ответ
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(respMove); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Move selected: %d\n", respMove.Index)
}

// Метод для получения доступных ходов
func (b *Board) AvailableMoves() []int {
	var moves []int

	// Итерируем по каждому индексу доски
	for i := 0; i < len(b); i++ {
		// Проверяем, свободна ли клетка
		if b[i] == TokenEmpty {
			moves = append(moves, i)
		}
	}

	return moves // Возвращаем срез доступных ходов
}

// Метод для представления состояния доски в виде строки (для логирования)
func (b *Board) String() string {
	var result string
	for row := 0; row < Rows; row++ {
		result += "| "
		for col := 0; col < Cols; col++ {
			index := row*Cols + col
			result += string(b[index]) + " | "
		}
		result += "\n"
	}
	return result
}

func joinGame(name, url string) error {
	/*
		Создание структуры запроса
		Создается структура JoinRequest с полями Name и URL, которая будет отправлена на сервер судейства.
		Поле Name содержит имя бота, а URL — это URL, по которому сервер сможет отправлять запросы на этот бот.
	*/
	joinRequest := JoinRequest{
		Name: name,
		URL:  url,
	}

	/*
		Формирование JSON тела запроса
		Структура joinRequest сериализуется в JSON с помощью функции json.Marshal.
		Если произойдет ошибка во время сериализации, программа вернет её и завершит выполнение функции.
	*/
	data, err := json.Marshal(joinRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal join request: %w", err)
	}

	/*
		Отправка POST-запроса
		Функция http.Post отправляет POST-запрос на сервер судейства по адресу http://localhost:4444/join с телом запроса, содержащим JSON, и заголовком Content-Type: application/json.
		Если происходит ошибка при отправке запроса (например, сервер недоступен), она будет обработана и возвращена с описанием.
		После выполнения запроса ресурс тела ответа (resp.Body) будет закрыт с помощью defer resp.Body.Close().
	*/
	resp, err := http.Post("http://localhost:4444/join", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to join game: %w", err)
	}
	defer resp.Body.Close()

	/*
		Проверка статуса ответа
		Функция проверяет, успешен ли запрос. Статусом успеха для запроса на подключение является код 204 No Content, что означает, что сервер принял запрос и успешно добавил бота в игру.
		Если статус ответа отличается от 204, функция возвращает ошибку с сообщением о неудачной попытке подключения.
	*/
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to join game, status: %s", resp.Status)
	}

	fmt.Println("Successfully joined the game")
	return nil
}
