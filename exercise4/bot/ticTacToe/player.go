package ticTacToe

type Token string

const (
	TokenEmpty Token = " "
	TokenX     Token = "x"
	TokenO     Token = "o"
)

type Player struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Token Token  `json:"token"`
}

func New(name string, url string) *Player {
	return &Player{
		Name:  name,
		URL:   url,
		Token: TokenEmpty,
	}
}
