package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) NewPlayer(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	newPlayer := new(AddPlayer)

	err := json.NewDecoder(r.Body).Decode(newPlayer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.LeaderBoard.NewPlayer(newPlayer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	fmt.Println("NewPlayer")

}

type AddPlayer struct {
	Data *AddPlayerReq `json:"data"`
}

type AddPlayerReq struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
	Name  string `json:"name"`
}

func (h *AddPlayer) GetName() string {
	return h.Data.Name

}
func (h *AddPlayer) GetId() string {
	return h.Data.ID

}

func (h *AddPlayer) GetScore() int {
	return h.Data.Score

}
