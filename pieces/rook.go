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

func (k Rook) Row() int { return k.Piece.Row }

func (k Rook) Col() int { return k.Piece.Col }

func (k Rook) Name() string { return k.Piece.Name }
