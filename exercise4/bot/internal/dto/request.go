package dto

import (
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/model"
)

type RequestJoin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RequestMove struct {
	Board *model.Board `json:"board"`
	Token model.Token  `json:"token"`
}
