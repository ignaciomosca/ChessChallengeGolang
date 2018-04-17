package pieces

import piece "chess/primitives"

type Knight struct {
	piece.Piece
}

func CreateKnight(row, col int) Knight {
	return Knight{piece.Piece{"N", row, col}}
}

func (k Knight) Attacks() bool {
	return false
}
