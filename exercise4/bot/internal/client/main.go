package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

var (
	lcl  = "http://127.0.0.1"
	port = os.Getenv("PORT")
	url  = fmt.Sprintf("%s:%s", lcl, port)
	name = os.Getenv("NAME")
)

func SendJoinRequest(ctx context.Context) {
	data := map[string]string{
		"name": name,
		"url":  url,
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		slog.Error("marshal error", data)
	}

	serverUrl := fmt.Sprintf("%s:4444/join", lcl)
	req, err := http.NewRequestWithContext(ctx, "POST", serverUrl, bytes.NewBuffer(dataJson))
	if err != nil {
		slog.Error("join request error", "error", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("join request error do", "error", err)
		return
	}
	fmt.Println("ok")
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	fmt.Println(string(response))
}
