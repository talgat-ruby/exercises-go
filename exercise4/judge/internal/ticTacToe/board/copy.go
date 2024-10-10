package board

func (b *Board) Copy() *Board {
	var cpy Board
	cpy = *b
	return &cpy
}
