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

	for j := 0; j < len(b); j++ {
		if b[j] == TokenEmpty {
			if j == 0 {
				if (b[j+1] == b[j+2] && b[j+2] != TokenEmpty) || (b[j+3] == b[j+6] && b[j+6] != TokenEmpty) || (b[j+4] == b[j+8] && b[j+8] != TokenEmpty){
					return j
				}
			}
			if j == 1 {
				if (b[j-1] == b[j+1] && b[j+1] != TokenEmpty) || (b[j+3] == b[j+6] && b[j+6] != TokenEmpty) {
					return j
				}
			}
			if j == 2 {
				if (b[j-1] == b[j-2] && b[j-2] != TokenEmpty) || (b[j+3] == b[j+6] && b[j+6] != TokenEmpty) || (b[j+2] == b[j+4] && b[j+4] != TokenEmpty) {
					return j
				}
			}
			if j == 3 {
				if (b[j+1] == b[j+2] && b[j+2] != TokenEmpty) || (b[j+3] == b[j-3] && b[j-3] != TokenEmpty) {
					return j
				}
			}
			if j == 4 {
				if (b[j-1] == b[j+1] && b[j+1] != TokenEmpty) || (b[j+3] == b[j-3] && b[j-3] != TokenEmpty) || (b[j-2] == b[j+2] && b[j+2] != TokenEmpty) || (b[j-4] == b[j+4] && b[j+4] != TokenEmpty) {
					return j
				}

			}
			if j == 5 {
				if (b[j-1] == b[j-2] && b[j-2] != TokenEmpty) || (b[j+3] == b[j-3] && b[j-3] != TokenEmpty) {
					return j
				}
			}
			if j == 6 {
				if (b[j+1] == b[j+2] && b[j+2] != TokenEmpty) || (b[j-3] == b[j-6] && b[j-6] != TokenEmpty) || (b[j-2] == b[j-4] && b[j-4] != TokenEmpty) {
					return j
				}
			}
			if j == 7 {
				if (b[j-1] == b[j+1] && b[j+1] != TokenEmpty) || (b[j-3] == b[j-6] && b[j-6] != TokenEmpty) {
					return j
				}
			}
			if j == 8 {
				if (b[j-1] == b[j-2] && b[j-2]!= TokenEmpty) || (b[j-3] == b[j-6] && b[j-6] != TokenEmpty) || (b[j-4] == b[j-8] && b[j-8] != TokenEmpty) {
					return j
				}
			}
		}
	}
	if b[0] == TokenEmpty {
		return 0
	}
	if b[2] == TokenEmpty {
		return 2
	}
	if b[6] == TokenEmpty {
		return 6
	}
	if b[8] == TokenEmpty {
		return 8
	}
	for i, t := range b {
		if t == TokenEmpty {
			return i
		}
	}
	return -1
}
