package rest

import (
	"context"
	"github.com/my/repo/internal/config"
	"github.com/my/repo/internal/rest/handler"
	"github.com/my/repo/internal/rest/router"
	"github.com/my/repo/internal/types/controller"
	"net/http"
)

type Rest struct {
	conf   *config.Config
	router *router.Router
}

func NewRest(conf *config.Config, ctrl controller.Controller) *Rest {

	h := handler.NewHandler(ctrl)
	r := router.NewRouter(h)
	return &Rest{
		conf:   conf,
		router: r,
	}
}

func (a *Rest) Start(ctx context.Context) error {
	//mux := http.NewServeMux()
	mux := a.router.Start(ctx)
	srv := &http.Server{
		Handler: mux,
		Addr:    a.conf.API.Rest.Host + ":" + a.conf.API.Rest.Port,
	}
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
