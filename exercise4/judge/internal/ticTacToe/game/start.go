package game

import (
	"context"
	"fmt"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/match"
	"github.com/talgat-ruby/exercises-go/exercise4/judge/internal/ticTacToe/player"
)

func (g *Game) Start(ctx context.Context) error {
	/*
		 Проверка текущего состояния игры
		 Проверяется, находится ли игра в состоянии ожидания (StatePending).
		Если игра уже находится в другом состоянии (например, она уже запущена или завершена), функция возвращает ошибку с указанием текущего состояния игры.
	*/
	if g.State != StatePending {
		return fmt.Errorf("game state is %s", g.State)
	}

	/*
		Проверка количества игроков
		Проверяется, есть ли хотя бы 2 игрока в игре. Если игроков недостаточно, возвращается ошибка с указанием текущего количества игроков.
	*/
	if len(g.Players) < 2 {
		return fmt.Errorf("game has no enough players, %d player(s)", len(g.Players))
	}

	/*
		Подготовка к запуску игры
		Вызов метода g.startPrint() — вероятно, это просто вывод на консоль или логирование, что игра начинается.
		Затем состояние игры меняется на StateRunning, что означает, что игра теперь в процессе.
	*/
	g.startPrint()
	g.State = StateRunning

	/*
		Запуск матчей между игроками в горутине
		Игра запускает отдельную горутину для инициализации матчей между игроками.
		Внутри этой горутины происходит двойной цикл: каждый игрок играет матч с каждым другим игроком (i с j, где j > i). Это гарантирует, что каждый игрок встретится с каждым другим хотя бы один раз.
		Для каждой пары игроков создаётся новый матч с помощью match.New(pls), и матч добавляется в список матчей игры (g.Matches).
		Метод newMatch.Start(ctx, len(g.Matches)) запускает матч, где ctx — это контекст для управления матчем, а len(g.Matches) — номер матча.
		После завершения всех матчей вызывается g.endPrint(), который, вероятно, выводит информацию о завершении игры.
	*/
	go func(ctx context.Context, players []*player.Player) {
		for i := 0; i < len(players); i++ {
			for j := i + 1; j < len(players); j++ {
				pls := [2]*player.Player{players[i], players[j]}
				newMatch := match.New(pls)
				g.Matches = append(g.Matches, newMatch)
				newMatch.Start(ctx, len(g.Matches))
			}
		}
		g.endPrint()
	}(ctx, g.Players)

	return nil
}

func (g *Game) startPrint() {
	fmt.Printf(
		`
Starting Game!

Our players are:
%s
`,
		g.playersList(),
	)
	<-time.After(time.Second)
}

func (g *Game) endPrint() {
	fmt.Printf(
		`
Game Completed!

Leaderboard:
%s
`,
		g.leaderboard(),
	)
	<-time.After(time.Second)
}
