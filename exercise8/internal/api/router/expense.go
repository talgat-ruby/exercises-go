package router

import (
	"context"
	"net/http"
)

func (r *Router) expense(ctx context.Context) {
	r.router.Handle("GET /expenses", r.mid.Authenticator(http.HandlerFunc(r.handler.FindExpenses)))
	r.router.Handle("GET /expense/{id}", r.mid.Authenticator(http.HandlerFunc(r.handler.FindExpense)))
	r.router.Handle("POST /expense", r.mid.Authenticator(http.HandlerFunc(r.handler.CreateExpense)))
	r.router.Handle("PUT /expense/{id}", r.mid.Authenticator(http.HandlerFunc(r.handler.UpdateExpense)))
	r.router.Handle("DELETE /expense/{id}", r.mid.Authenticator(http.HandlerFunc(r.handler.DeleteExpense)))

}
