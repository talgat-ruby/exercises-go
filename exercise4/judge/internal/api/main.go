package api

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/api/router"
)

type Api struct {
	srv *http.Server
}

func New() *Api {
	return &Api{}
}

func (api *Api) Start(ctx context.Context, port string) error {
	// Функция создает новый экземпляр роутера, который, вероятно, управляет маршрутизацией HTTP-запросов (подобно http.ServeMux). Этот роутер будет использоваться в качестве обработчика запросов для сервера.
	r := router.New()

	// Инициализация HTTP-сервера
	api.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	/*
		Логирование старта сервиса
		Перед запуском сервера выводится информационное сообщение в лог, которое говорит о том, что сервис запускается на указанном порту.
	*/
	slog.InfoContext(
		ctx,
		"starting service",
		"port", port,
	)

	/*
		Запуск сервера
		api.srv.ListenAndServe() запускает сервер и начинает прослушивание запросов на указанном порту.
		Если метод ListenAndServe возвращает ошибку, которая не является http.ErrServerClosed, это означает, что произошла серьезная ошибка при запуске сервера. В этом случае ошибка логируется, и функция возвращает эту ошибку.
		Если сервер был корректно остановлен (например, при вызове http.Server.Close), он возвращает ошибку http.ErrServerClosed, которая игнорируется в этом коде (так как это ожидаемое поведение).
	*/
	if err := api.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.ErrorContext(ctx, "service error", "error", err)
		return err
	}

	/*
		Возвращение успешного результата
		Если сервер был запущен успешно и не возникло ошибок при старте, функция возвращает nil, указывая, что сервер был запущен корректно.
	*/
	return nil
}

func (api *Api) Stop(ctx context.Context) error {
	if err := api.srv.Shutdown(ctx); err != nil {
		slog.ErrorContext(ctx, "server shutdown error", "error", err)
		return err
	}

	return nil
}
