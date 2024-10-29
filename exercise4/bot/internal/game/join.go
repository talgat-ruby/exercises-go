package game

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	constants2 "github.com/talgat-ruby/exercises-go/exercise4/bot/internal/constants"
	"github.com/talgat-ruby/exercises-go/exercise4/bot/internal/dto"
	"log/slog"
	"net/http"
	"time"
)

func JoinGame() error {

	// timeout after 5 sec
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := dto.RequestJoin{Name: constants2.NAME, URL: constants2.FullUrl}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal join request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		constants2.Host+constants2.JoinPath,
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return fmt.Errorf("error creating join request %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making join request %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed making join request %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	slog.Info("successfully joined the game")

	return nil
}
