package board

import (
	"fmt"
)

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
