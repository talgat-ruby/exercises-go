package governor

import "github.com/my/repo/internal/governor/board"

type Governor struct {
	*board.LeaderBoard
}

func NewGovernor() *Governor {
	return &Governor{
		LeaderBoard: new(board.LeaderBoard),
	}
}
