package router

import (
	"context"
)

func (r Router) player(_ context.Context) {
	r.router.HandleFunc("GET /players", r.handler.ListPlayers)
	r.router.HandleFunc("POST /player", r.handler.NewPlayer)

}
