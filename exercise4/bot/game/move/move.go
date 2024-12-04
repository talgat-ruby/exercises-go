package move

import "math"

func GetBestMove(board []string, token string) int {
	bestScore := math.Inf(-1)
	bestMove := -1

	for i := 0; i < 9; i++ {
		if board[i] == " " {
			board[i] = token
			score := minimax(board, 0, false, token)
			board[i] = " "
			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}
	return bestMove
}

func minimax(board []string, depth int, isMaximizing bool, token string) float64 {
	opponentToken := "o"
	if token == "o" {
		opponentToken = "x"
	}

	if isWinning(board, token) {
		return 1
	}
	if isWinning(board, opponentToken) {
		return -1
	}
	if isBoardFull(board) {
		return 0
	}

	if isMaximizing {
		bestScore := math.Inf(-1)
		for i := 0; i < 9; i++ {
			if board[i] == " " {
				board[i] = token
				score := minimax(board, depth+1, false, token)
				board[i] = " "
				bestScore = math.Max(bestScore, score)
			}
		}
		return bestScore
	} else {
		bestScore := math.Inf(1)
		for i := 0; i < 9; i++ {
			if board[i] == " " {
				board[i] = opponentToken
				score := minimax(board, depth+1, true, token)
				board[i] = " "
				bestScore = math.Min(bestScore, score)
			}
		}
		return bestScore
	}
}

func isBoardFull(board []string) bool {
	for _, cell := range board {
		if cell == " " {
			return false
		}
	}
	return true
}
