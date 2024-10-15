package round

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/move"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

func (r *Round) Start(ctx context.Context, num int) {
	r.startPrint(num)

	for i := 0; ; i++ {
		newMove := move.New(r.Players[i%len(r.Players)], r.Board)
		r.Moves = append(r.Moves, newMove)

		ind, err := newMove.Start(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "invalid move", "error", err)
			r.Winner = r.Players[(i+1)%len(r.Players)]
			break
		}

		if winnerToken := r.Board.Evaluate(ind, *newMove.Player.Token); winnerToken != nil {
			r.Winner = r.findWinnerByToken(*winnerToken)
			break
		}
	}

	r.endPrint(num)
}

func (r *Round) findWinnerByToken(t internal.Token) *player.Player {
	for _, p := range r.Players {
		if *p.Token == t {
			return p
		}
	}

	return nil
}

func (r *Round) startPrint(num int) {
	fmt.Printf(
		`
Round #%d: %s(%s) vs %s(%s)
`,
		num,
		r.Players[0].Name,
		*r.Players[0].Token,
		r.Players[1].Name,
		*r.Players[1].Token,
	)
	<-time.After(time.Second)
}

func (r *Round) endPrint(num int) {
	if r.Winner == nil {
		fmt.Printf(
			`
Round #%d: is draw
`,
			num,
		)
		<-time.After(time.Second)
		return
	}

	fmt.Printf(
		`
Round #%d: %s(%s) won!
`,
		num,
		r.Winner.Name,
		*r.Winner.Token,
	)
	<-time.After(time.Second)
}
