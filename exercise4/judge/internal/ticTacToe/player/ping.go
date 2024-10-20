package player

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (p *Player) Ping(ctx context.Context) error {
	// timeout after 5 sec
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/ping", p.URL),
		nil,
	)
	if err != nil {
		return fmt.Errorf("error creating request %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making ping request %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed making ping request %d - %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return nil
}
