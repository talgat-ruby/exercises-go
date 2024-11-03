package player

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type RequestJoin struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

var (
	host = "http://127.0.0.1"
	name = os.Getenv("NAME")
	port = os.Getenv("PORT")
	URL  = fmt.Sprintf("%s:%s", host, port)
)

func JoinToGame(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	join := RequestJoin{
		Name: name,
		URL:  URL,
	}
	jsonData, err := json.Marshal(join)
	if err != nil {
		return fmt.Errorf("failed to marshal move request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s:4444/join", host),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed making join request %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return nil
}
