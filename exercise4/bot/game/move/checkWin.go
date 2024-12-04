package move

func isWinning(board []string, token string) bool {
	winningPositions := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // columns
		{0, 4, 8}, {2, 4, 6}, // diagonals
	}
	for _, positions := range winningPositions {
		if board[positions[0]] == token &&
			board[positions[1]] == token &&
			board[positions[2]] == token {
			return true
		}
	}
	return false
}
