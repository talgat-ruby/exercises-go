package handler

import (
	"log/slog"
	"net/http"
)

func (h *Handler) Start(w http.ResponseWriter, r *http.Request) {
	/*
		Получение контекста и логирование
		Получается контекст (ctx) из запроса r. Контекст может передавать данные, связанные с запросом, такие как таймауты, отмены и другие параметры.
		Затем создается объект для логирования с дополнительными параметрами — имя хендлера ("Start") и путь запроса (r.URL.Path).
	*/
	ctx := r.Context()
	log := slog.With(
		"handler", "Start",
		"path", r.URL.Path,
	)

	/*
		Запуск игры через h.game.Start
		Вызывается метод Start объекта игры (h.game.Start(ctx)), который запускает процесс игры.
		Если возникла ошибка при запуске игры (например, игра уже запущена или есть какая-то логическая ошибка), она будет залогирована и передана в ответ клиенту с кодом 400 Bad Request.
	*/
	if err := h.game.Start(ctx); err != nil {
		log.ErrorContext(
			ctx,
			"game was not started",
			"error", err,
		)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	/*
		Логирование успешного старта игры
		Если игра была успешно запущена, это событие логируется с сообщением "game started".
		В ответ клиенту отправляется код статуса 204 No Content, что означает, что запрос был успешно обработан, но тело ответа не содержит данных.
	*/
	log.InfoContext(
		ctx,
		"game started",
	)
	w.WriteHeader(http.StatusNoContent)
}
