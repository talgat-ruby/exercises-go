package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) ListPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListPlayers", "handler")
	//ctx := r.Context()
	resp, err := h.LeaderBoard.ListPlayers()
	if resp == nil {
		fmt.Println("resp is nil")
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
