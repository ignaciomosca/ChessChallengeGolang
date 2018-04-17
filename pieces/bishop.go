package pieces

import piece "chess/primitives"

type Bishop struct {
	piece.Piece
}

func CreateBishop(row, col int) Bishop {
	return Bishop{piece.Piece{"B", row, col}}
}

func (bishop Bishop) Attacks() bool {
	return false
}
