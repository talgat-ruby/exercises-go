package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type joinRequest struct {
	name string `json:"name"`
	url  string `json:"url"`
}

func SendJoinRequest(ctx context.Context, name string, url string) {
	data := joinRequest{name, url}

	dataJson, err := json.Marshal(data)
	if err != nil {
		slog.Error("marshal error", data)
	}
	resp, err := http.NewRequest("POST", url, bytes.NewBuffer(dataJson))
	resp.Header.Set("Content-Type", "application/json")
	if err != nil {
		slog.Error("new request error", data)
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	fmt.Println(string(response))
}
