package controller

type NewPlayerReq interface {
	GetName() string
	GetId() string
	GetScore() int
}
type NewPlayerResp interface {
}
