package move

import (
	"context"
	"fmt"
	"time"
)

func (m *Move) Start(ctx context.Context) (int, error) {
	ind, err := m.apply(ctx)
	m.Board = m.Board.Copy()
	if err != nil {
		return 0, err
	}

	m.print()

	return ind, nil
}

func (m *Move) apply(ctx context.Context) (int, error) {
	ind, err := m.Player.Move(ctx, m.Board)
	if err != nil {
		return 0, err
	}

	if err := m.Board.Set(ind, *m.Player.Token); err != nil {
		return 0, err
	}

	return ind, nil
}

func (m *Move) print() {
	fmt.Printf(
		`
Move by %s(%s)
%s
`,
		m.Player.Name,
		*m.Player.Token,
		m.Board.String(),
	)
	<-time.After(time.Second)
}
