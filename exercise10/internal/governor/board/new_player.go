package board

import "github.com/my/repo/internal/types/controller"

func (l *LeaderBoard) NewPlayer(req controller.NewPlayerReq) (controller.NewPlayerResp, error) {
	newPlayer := CreateNewPlayer(req.GetName(), req.GetId(), req.GetScore())
	resp, err := l.MyRedis.CreatePlayer(newPlayer)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type newPlayerReq struct {
	controller.NewPlayerReq
	id    string
	name  string
	score int
}

func CreateNewPlayer(name string, id string, score int) *newPlayerReq {
	return &newPlayerReq{
		id:    id,
		name:  name,
		score: score,
	}
}

// func (r newPlayerReq) GetName() string {
// 	return r.name
// }

// func (r newPlayerReq) GetId() string {
// 	return r.id
// }
// func (r newPlayerReq) GetScore() int {
// 	return r.score
// }

func (m *newPlayerReq) GetPlayerID() string {
	return m.id
}

func (m *newPlayerReq) GetPlayerName() string {
	return m.name
}

func (m *newPlayerReq) GetPlayerScore() int {
	return m.score
}
