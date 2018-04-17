package pieces

import piece "chess/primitives"

type King struct {
	piece.Piece
}

func CreateKing(row, col int) King {
	return King{piece.Piece{"K", row, col}}
}

func (k King) Attacks() bool {
	return false
}
