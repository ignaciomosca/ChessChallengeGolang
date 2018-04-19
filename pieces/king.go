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

func (k King) Row() int { return k.Piece.Row }

func (k King) Col() int { return k.Piece.Col }

func (k King) Name() string { return k.Piece.Name }
