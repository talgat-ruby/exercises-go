package board

import (
	"errors"
	"fmt"
	"github.com/my/repo/internal/types/controller"
)

func (l *LeaderBoard) ListPlayers() (controller.ListPlayerResp, error) {
	fmt.Println("method", "ListPlayers:governor")
	//request := newListPlayerReq()
	response, err := l.MyRedis.ListPlayers()
	if response == nil {
		return nil, errors.New("response is nil")
	}
	if err != nil {
		return nil, err
	}

	ps := make([]*Player, 0, len(response.GetList()))

	for _, it := range response.GetList() {
		ps = append(ps, newPlayer(it.GetPlayerID(), it.GetPlayerName(), it.GetPlayerScore()))
	}
	fmt.Println("method", "ListPlayers:governor", "end")
	return newListPlayer(ps), nil

}

type ListPlayerReq struct{}
type ListPlayerResp struct {
	Players []*Player
}

func newListPlayerReq() *ListPlayerReq {
	return &ListPlayerReq{}
}

func newListPlayer(ps []*Player) *ListPlayerResp {
	return &ListPlayerResp{ps}
}

func (i *ListPlayerResp) GetList() []controller.ItemPlayerResp {
	list := make([]controller.ItemPlayerResp, 0, len(i.Players))

	for _, p := range i.Players {
		list = append(list, p)
	}

	return list
}

type Player struct {
	id    string
	name  string
	score int
}

func newPlayer(id, name string, score int) *Player {
	return &Player{id, name, score}
}

func (p *Player) GetID() string {
	return p.id
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetScore() int {
	return p.score
}
