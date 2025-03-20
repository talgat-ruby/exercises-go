package pubsub

type CreatePlayerReq interface {
	GetPlayerID() string
	GetPlayerName() string
	GetPlayerScore() int
}
type CreatePlayerResp interface {
}
