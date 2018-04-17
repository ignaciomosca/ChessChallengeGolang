package pieces

import piece "chess/primitives"

type Queen struct {
	piece.Piece
}

func CreateQueen(row, col int) Queen {
	return Queen{piece.Piece{"Q", row, col}}
}

func (queen Queen) Attacks() bool {
	return false
}
