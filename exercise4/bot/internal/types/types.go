package types

type MoveRequest struct {
	Board []string `json:"board"`
	Token string   `json:"token"`
}

type ResponseMove struct {
	Index int `json:"index"`
}
