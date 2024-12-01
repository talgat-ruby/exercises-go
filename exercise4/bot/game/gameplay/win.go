package gameplay

func isWinning(board []string, token string) bool {
	// Задаем индексы выигрышных позиций
	winningPositions := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Горизонтальные
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Вертикальные
		{0, 4, 8}, {2, 4, 6}, // Диагонали
	}

	// Проверяем каждую выигрышную позицию
	for _, positions := range winningPositions {
		// Если все три клетки в позиции заняты нужным токеном
		if board[positions[0]] == token && board[positions[1]] == token && board[positions[2]] == token {
			return true
		}
	}

	// Если ни одна позиция не подошла
	return false
}
