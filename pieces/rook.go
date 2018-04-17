package pieces

import piece "chess/primitives"

type Rook struct {
	piece.Piece
}

func CreateRook(row, col int) Rook {
	return Rook{piece.Piece{"R", row, col}}
}

func (rook Rook) Attacks() bool {
	return false
}
