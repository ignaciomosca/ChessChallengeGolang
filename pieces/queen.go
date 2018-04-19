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

func (k Queen) Row() int { return k.Piece.Row }

func (k Queen) Col() int { return k.Piece.Col }

func (k Queen) Name() string { return k.Piece.Name }
