package ticTacToe

type State string

const (
	StatePending  State = "PENDING"
	StateRunning  State = "RUNNING"
	StateFinished State = "FINISHED"
)

type Game struct {
	State   State     `json:"state"`
	Players []*Player `json:"players"`
	Matches []*Match  `json:"matches"`
}

type Match struct {
	Players [2]*Player `json:"players"`
	Rounds  []*Round   `json:"rounds"`
}

type Round struct {
	Players [2]*Player `json:"players"`
	Board   *Board     `json:"board"`
	Moves   []*Move    `json:"moves"`
	Winner  *Player    `json:"winner"`
}

type Move struct {
	Player *Player `json:"players"`
	Board  *Board  `json:"board"`
}

const (
	Cols = 3
	Rows = 3
)

type Board [Cols * Rows]Token

const (
	MinScore  = -1_000_000
	MaxScore  = 1_000_000
	WinScore  = 10
	LoseScore = -10
)

// Вычисляет наилучший ход для текущего игрока
func CalculateBestMove(board *Board, player Token) int {
	bestScore := MinScore
	bestMove := -1

	for i := 0; i < len(*board); i++ {
		if (*board)[i] == TokenEmpty {
			(*board)[i] = player
			score := minimax(board, 0, false, player, MinScore, MaxScore)
			(*board)[i] = TokenEmpty

			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}
	return bestMove
}

// Алгоритм Minimax с альфа-бета отсечением
func minimax(board *Board, depth int, isMaximizing bool, player Token, alpha, beta int) int {
	winner := checkWinner(board)
	if winner == player {
		return WinScore - depth
	} else if winner != TokenEmpty {
		return LoseScore + depth
	} else if isBoardFull(board) {
		return 0
	}

	opponent := switchPlayer(player)

	if isMaximizing {
		maxEval := MinScore
		for i := 0; i < len(*board); i++ {
			if (*board)[i] == TokenEmpty {
				(*board)[i] = player
				eval := minimax(board, depth+1, false, player, alpha, beta)
				(*board)[i] = TokenEmpty
				maxEval = max(maxEval, eval)
				alpha = max(alpha, eval)
				if beta <= alpha {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := MaxScore
		for i := 0; i < len(*board); i++ {
			if (*board)[i] == TokenEmpty {
				(*board)[i] = opponent
				eval := minimax(board, depth+1, true, player, alpha, beta)
				(*board)[i] = TokenEmpty
				minEval = min(minEval, eval)
				beta = min(beta, eval)
				if beta <= alpha {
					break
				}
			}
		}
		return minEval
	}
}

// Проверка победителя на доске
func checkWinner(board *Board) Token {
	winningCombinations := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Ряды
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Столбцы
		{0, 4, 8}, {2, 4, 6}, // Диагонали
	}

	for _, combo := range winningCombinations {
		if (*board)[combo[0]] != TokenEmpty &&
			(*board)[combo[0]] == (*board)[combo[1]] &&
			(*board)[combo[1]] == (*board)[combo[2]] {
			return (*board)[combo[0]]
		}
	}
	return TokenEmpty
}

// Проверка на заполненность доски
func isBoardFull(board *Board) bool {
	for _, cell := range *board {
		if cell == TokenEmpty {
			return false
		}
	}
	return true
}

// Переключение текущего игрока
func switchPlayer(player Token) Token {
	if player == TokenX {
		return TokenO
	}
	return TokenX
}

// Вспомогательные функции для нахождения максимума и минимума
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
