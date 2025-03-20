package controller

type ListPlayerReq interface {
}

type ListPlayerResp interface {
	GetList() []ItemPlayerResp
}
type ItemPlayerResp interface {
	GetID() string
	GetName() string
	GetScore() int
}
