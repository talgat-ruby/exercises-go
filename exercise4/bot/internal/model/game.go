package model

type Game struct {
	State string `json:"state"`
}

func NewGame() *Game {
	return &Game{
		State: "start",
	}
}
