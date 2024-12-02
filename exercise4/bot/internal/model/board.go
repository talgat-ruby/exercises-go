package model

import "fmt"

type Token string

const (
	TokenEmpty Token = " "
	TokenX     Token = "x"
	TokenO     Token = "o"
)

const (
	Cols = 3
	Rows = 3
)

var motion int

type Board [Cols * Rows]Token

func (b *Board) CalculateNewPosition(token Token) int {
	ind := b.calculateNewPosition(token)
	b.logMove(ind, token)
	return ind
}

func (b *Board) logMove(pos int, token Token) {
	b[pos] = token
	fmt.Println(b.String())
}

func (b *Board) String() string {
	str := fmt.Sprintf(
		`
	%s|%s|%s
	-----
	%s|%s|%s
	-----
	%s|%s|%s
`,
		b[0],
		b[1],
		b[2],
		b[3],
		b[4],
		b[5],
		b[6],
		b[7],
		b[8],
	)

	return str
}

func (b *Board) calculateNewPosition(token Token) int {
	// Увеличиваем счётчик ходов
	motion++
	
	// Проверяем количество заполненных ячеек на доске
	count := 0
	for _, p := range b {
		if p != TokenEmpty {
			count++
		}
	}
	
	// Если на доске только один ход и центральная ячейка пуста, ставим токен в центр
	if count == 1 && b[4] == TokenEmpty {
		return 4
	}

	// Получаем все возможные варианты выигрышных комбинаций
	alloptions := b.allWin()             
	// Находим оптимальные ходы для текущего игрока
	optmoves := optimalMoves(alloptions) 

	// Проверяем, есть ли выигрышные или защитные ходы
	if len(optmoves) > 0 {
		// Если есть возможность выиграть, выбираем этот ход. 
		// Если нет, выбираем ход для защиты.
		emptyToken := findEmptySlotForToken(optmoves, token)
		return emptyToken
	}

	// Если нет выигрышных или защитных ходов, ищем пустые углы
	emtyCorner := b.findEmptyCorners()
	return emtyCorner
}

func (b *Board) allWin() []map[int]Token {
	// Определяем возможные выигрышные комбинации
	winCombinations := [][]int{
		{0, 1, 2}, // Горизонтальные линии
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6}, // Вертикальные линии
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8}, // Диагонали
		{2, 4, 6},
	}

	// Инициализируем массив для хранения всех выигрышных комбинаций
	alloptions := []map[int]Token{}

	// Заполняем массив выигрышных комбинаций текущими значениями с доски
	for _, combination := range winCombinations {
		option := make(map[int]Token)
		for _, index := range combination {
			option[index] = b[index]
		}
		alloptions = append(alloptions, option)
	}

	return alloptions
}

func optimalMoves(alloptions []map[int]Token) []map[int]Token {
	// Инициализируем массив для хранения оптимальных ходов
	optmoves := []map[int]Token{}

	// Перебираем все выигрышные комбинации
	for _, option := range alloptions {
		a := []Token{}
		for _, value := range option {
			a = append(a, value)
		}
		
		// Проверяем условия для победы или защиты
		if (a[0] == a[1] && a[0] != TokenEmpty && a[2] == TokenEmpty) ||
			(a[0] == a[2] && a[0] != TokenEmpty && a[1] == TokenEmpty) ||
			(a[1] == a[2] && a[1] != TokenEmpty && a[0] == TokenEmpty) {
			// Добавляем текущую комбинацию в оптимальные ходы
			optmoves = append(optmoves, option)
		}
	}
	return optmoves
}

func findEmptySlotForToken(variants []map[int]Token, token Token) int {
	var emptyIndex int
	
	// Перебираем все варианты
	for _, variant := range variants {
		hasToken := false

		// Проверяем, есть ли токен и запоминаем индекс пустой ячейки
		for key, value := range variant {
			if value == token {
				hasToken = true
			} else if value == TokenEmpty {
				emptyIndex = key
			}
		}

		// Если найден токен, возвращаем индекс пустой ячейки
		if hasToken {
			return emptyIndex
		}
	}

	return emptyIndex // Возвращаем индекс пустой ячейки 
}

func (b *Board) findEmptyCorners() int {
	// Определяем индексы угловых ячеек
	corners := []int{0, 2, 6, 8} 

	// Перебираем угловые ячейки
	for _, index := range corners {
		if b[index] == TokenEmpty {
			// Пропускаем индекс 2, если Motion == 2
			if index == 2 && motion == 2 {
				continue
			}
			return index // Возвращаем индекс пустой угловой ячейки
		}
	}

	return -1 // Возвращаем -1, если все угловые ячейки заняты
}