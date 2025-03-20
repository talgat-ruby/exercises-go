package restboard

import (
	"github.com/my/repo/internal/types/controller"
)

type RestBoard struct {
	ctrl controller.Controller
}

func NewRestBoard(ctrl controller.Controller) *RestBoard {
	return &RestBoard{
		ctrl: ctrl,
	}
}
