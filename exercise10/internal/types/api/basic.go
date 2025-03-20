package api

import "context"

type Api interface {
	Start(ctx context.Context) error
}
