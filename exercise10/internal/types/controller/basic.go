package controller

type LeaderBoard interface {
	ListPlayers() (ListPlayerResp, error)
	NewPlayer(NewPlayerReq) (NewPlayerResp, error)
}

type Controller interface {
	LeaderBoard
}
