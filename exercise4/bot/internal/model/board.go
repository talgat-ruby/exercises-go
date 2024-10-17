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
	for i, t := range b {
		if t == TokenEmpty {
			return i
		}
	}
	return -1
}
