package player

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/board"
)

type RequestMove struct {
	Board *board.Board   `json:"board"`
	Token internal.Token `json:"token"`
}

type ResponseMove struct {
	Index int `json:"index"`
}

func (p *Player) Move(ctx context.Context, b *board.Board) (int, error) {
	/*
		Создание контекста с тайм-аутом
		Контекст создаётся с тайм-аутом в 5 секунд с помощью context.WithTimeout. Это ограничивает время ожидания ответа от удалённого сервиса, где клиент совершает ход. Если сервис не успеет ответить за это время, операция будет отменена.
		defer cancel() гарантирует, что после завершения метода контекст будет завершён и ресурсы освобождены.
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	/*
		Формирование запроса
		Здесь создаётся структура RequestMove, которая включает текущую доску (Board) и токен игрока (Token), например, X или O.
		С помощью json.Marshal эта структура сериализуется в JSON-формат для отправки в теле HTTP-запроса.
		В случае ошибки при сериализации возвращается ошибка.
	*/
	move := RequestMove{
		Board: b,
		Token: *p.Token,
	}
	jsonData, err := json.Marshal(move)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal move request: %w", err)
	}

	/*
		Создание HTTP-запроса
		Создаётся HTTP POST-запрос на URL игрока с добавлением /move к адресу (p.URL).
		Тело запроса (jsonData) передаётся как JSON в виде буфера.
		В заголовок запроса добавляется тип содержимого — application/json.
	*/
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/move", p.URL),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	/*
		Отправка запроса и обработка ответа
		HTTP-запрос отправляется через http.Client.
		Если запрос не удался (например, если не удаётся подключиться к удалённому серверу), возвращается ошибка.
		После завершения запроса соединение закрывается с помощью defer resp.Body.Close().
	*/
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	/*
		Чтение и проверка ответа
		Ответ сервера читается в body.
		Если при чтении произошла ошибка, она возвращается.
		Также проверяется статус ответа — если он не в пределах успешных кодов (200–299), возвращается ошибка с кодом и телом ответа.
	*/
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("request failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var respBody ResponseMove
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return respBody.Index, nil
}
