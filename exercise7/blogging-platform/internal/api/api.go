package api

import "context"

type Api struct {
}

func New() *Api {
	return &Api{}
}

func (a *Api) Start(ctx context.Context) error {
	return nil
}
