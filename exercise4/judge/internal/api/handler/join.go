package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/pkg/httputils/request"
)

type RequestJoin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ResponseJoin struct {
	Message string `json:"name"`
}

func (h *Handler) Join(w http.ResponseWriter, r *http.Request) {
	/*
		Получение контекста и логирование:
		ctx := r.Context() — извлекает контекст из запроса. Контекст используется для передачи информации о запросе, включая отмену или тайм-ауты.
		Лог создается с метками handler (указывающей на функцию Join) и path (с путём, по которому поступил запрос).
	*/
	ctx := r.Context()
	log := slog.With(
		"handler", "Join",
		"path", r.URL.Path,
	)

	/*
		Чтение и парсинг тела запроса:
		Здесь создается структура RequestJoin, которая содержит имя игрока (Name) и URL для взаимодействия с игроком (URL).
		Далее происходит парсинг тела запроса. В этом случае ожидается, что запрос будет содержать JSON с полями name и url. Функция request.JSON считывает JSON из тела запроса и пытается поместить данные в структуру reqBody.
		Если произошла ошибка при парсинге (например, JSON невалидный), она логируется, и возвращается ошибка 500 (Internal Server Error), завершив обработку.
	*/
	var reqBody RequestJoin
	if err := request.JSON(w, r, &reqBody); err != nil {
		log.ErrorContext(
			ctx,
			"failed to parse json body",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		Создание нового игрока:
		Если парсинг успешен, создается новый объект Player с именем и URL, переданными в теле запроса.
	*/
	p := player.New(reqBody.Name, reqBody.URL)

	/*
		Проверка доступности игрока (Ping):
		Выполняется пинг игрока с использованием его URL. Метод Ping отправляет HTTP-запрос по указанному URL, чтобы убедиться, что игрок доступен и может принимать запросы.
		Если игрок недоступен (например, возвращает 404 или другой код ошибки), то в логах регистрируется ошибка, и сервер возвращает ответ с кодом 400 (Bad Request) и сообщением, что пинг не удался.
	*/
	if err := p.Ping(ctx); err != nil {
		log.ErrorContext(
			ctx,
			"failed to ping player",
			"player.name", p.Name,
			"player.remote", p.URL,
			"error", err,
		)
		http.Error(
			w,
			fmt.Errorf("ping to url failed %w, check if player is running", err).Error(),
			http.StatusBadRequest,
		)
		return
	}

	/*
		Добавление игрока в игру:
		Если пинг успешен, сервер пытается добавить игрока в игру с помощью метода h.game.AddPlayer.
		Если добавление игрока не удалось (например, он уже в игре или возникла другая ошибка), она регистрируется в логах, и сервер возвращает ответ с кодом 400 (Bad Request).
	*/
	if err := h.game.AddPlayer(ctx, p); err != nil {
		log.ErrorContext(
			ctx,
			"player was not added",
			"player.name", p.Name,
			"player.url", p.URL,
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*
		Логирование успешного присоединения и завершение ответа:
		Если игрок успешно добавлен в игру, в логах сохраняется информация о новом игроке, и сервер завершает запрос, отправляя HTTP-код 204 (No Content), что означает успешное выполнение операции без тела ответа.
	*/
	log.InfoContext(
		ctx,
		"new player joined the game",
		"player.name", p.Name,
		"player.url", p.URL,
	)
	w.WriteHeader(http.StatusNoContent)
}
