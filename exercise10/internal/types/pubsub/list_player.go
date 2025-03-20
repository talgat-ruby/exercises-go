package pubsub

type ListPlayerReq interface{}
type ListPlayerResp interface {
	GetList() []ItemPlayerResp
}

type ItemPlayerResp interface {
	GetPlayerID() string
	GetPlayerName() string
	GetPlayerScore() int
}
