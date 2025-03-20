package pubsub

type LeaderBoard interface {
	CreatePlayer(CreatePlayerReq) (CreatePlayerResp, error)
	ListPlayers() (ListPlayerResp, error)
}

type MyRedis interface {
	LeaderBoard
}
