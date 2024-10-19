package player

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/board"
)

type RequestMove struct {
	Board *board.Board   `json:"board"`
	Token internal.Token `json:"token"`
}

type ResponseMove struct {
	Index int `json:"index"`
}

func (p *Player) Move(ctx context.Context, b *board.Board) (int, error) {
	// timeout after 5 sec
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	move := RequestMove{
		Board: b,
		Token: *p.Token,
	}
	jsonData, err := json.Marshal(move)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal move request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/move", p.URL),
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("request failed with status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var respBody ResponseMove
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return respBody.Index, nil
}
