package game

import "github.com/talgat-ruby/exercises-go/exercise4/bot/internal/types"

const (
	PlayerX = "x"
	PlayerO = "o"
	Empty   = " "
)

func BestMove(request *types.MoveRequest) int {

	bestScore := -1000
	bestMove := -1

	for i := 0; i < len(request.Board); i++ {
		if request.Board[i] == Empty {
			request.Board[i] = request.Token
			score := minimax(request.Board, 0, false, request.Token)
			request.Board[i] = Empty

			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}

	return bestMove
}

func minimax(board []string, depth int, isMaximizing bool, token string) int {
	opponent := PlayerO
	if token == PlayerO {
		opponent = PlayerX
	}

	score := evaluate(board, token, opponent)
	if score != 0 {
		return score
	}
	if isBoardFull(board) {
		return 0
	}

	if isMaximizing {
		bestScore := -1000
		for i := 0; i < len(board); i++ {
			if board[i] == Empty {
				board[i] = token
				score := minimax(board, depth+1, false, token)
				board[i] = Empty
				bestScore = max(score, bestScore)
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i := 0; i < len(board); i++ {
			if board[i] == Empty {
				board[i] = opponent
				score := minimax(board, depth+1, true, token)
				board[i] = Empty
				bestScore = min(score, bestScore)
			}
		}
		return bestScore
	}
}

// Оценка игрового поля
func evaluate(board []string, player string, opponent string) int {
	winPatterns := [][]int{
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // вертикальные линии для победы
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // горизонтальные линии для победы
		{0, 4, 8}, {2, 4, 6}, // диагонали для победы
	}

	for _, pattern := range winPatterns {
		if board[pattern[0]] == player && board[pattern[1]] == player && board[pattern[2]] == player {
			return 10
		}
		if board[pattern[0]] == opponent && board[pattern[1]] == opponent && board[pattern[2]] == opponent {
			return -10
		}
	}

	return 0
}

// Проверка, заполнено ли игровое поле
func isBoardFull(board []string) bool {
	for _, cell := range board {
		if cell == Empty {
			return false
		}
	}
	return true
}
